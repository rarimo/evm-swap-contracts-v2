// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

abstract contract AbstractSwapRouterMock {
    using SafeERC20 for IERC20;

    address public immutable WRAPPED_NATIVE;

    mapping(address => uint256) internal _reserves;
    mapping(address => mapping(address => address)) internal _pairs;

    constructor(address wrappedNative_) {
        WRAPPED_NATIVE = wrappedNative_;
    }

    function setReserve(address token_, uint256 amount_) external {
        uint256 balance_ = IERC20(token_).balanceOf(address(this));

        _reserves[token_] = amount_;

        if (amount_ > balance_) {
            IERC20(token_).safeTransferFrom(msg.sender, address(this), amount_ - balance_);
        } else if (amount_ < balance_) {
            IERC20(token_).safeTransfer(msg.sender, balance_ - amount_);
        }
    }

    function enablePair(address tokenA_, address tokenB_) external {
        _pairs[tokenA_][tokenB_] = address(1);
        _pairs[tokenB_][tokenA_] = address(1);
    }

    function _getReserves(
        address tokenA_,
        address tokenB_
    ) internal view returns (uint256 reserveA_, uint256 reserveB_) {
        require(tokenA_ != tokenB_, "UniswapRouterV2Mock: identical addresses");
        require(
            _pairs[tokenA_][tokenB_] != address(0) || _pairs[tokenB_][tokenA_] != address(0),
            "UniswapRouterV2Mock: pair does not exist"
        );

        return (_reserves[tokenA_], _reserves[tokenB_]);
    }

    function _getAmountIn(
        uint256 amountOut_,
        uint256 reserveIn_,
        uint256 reserveOut_
    ) internal pure returns (uint256 amountIn_) {
        require(amountOut_ > 0, "UniswapRouterV2Mock: insufficient output amount");
        require(reserveIn_ > 0 && reserveOut_ > 0, "UniswapRouterV2Mock: insufficient liquidity");

        amountIn_ = (amountOut_ * reserveIn_) / reserveOut_;
    }

    function _getAmountOut(
        uint256 amountIn_,
        uint256 reserveIn_,
        uint256 reserveOut_
    ) internal pure returns (uint256 amountOut_) {
        require(amountIn_ > 0, "UniswapRouterV2Mock: insufficient input amount");
        require(reserveIn_ > 0 && reserveOut_ > 0, "UniswapRouterV2Mock: insufficient liquidity");

        amountOut_ = (amountIn_ * reserveOut_) / reserveIn_;
    }

    function _getAmountsOut(
        uint256 amountIn_,
        address[] memory path_
    ) public view returns (uint256[] memory amounts_) {
        require(path_.length >= 2, "UniswapRouterV2Mock: invalid path");

        amounts_ = new uint256[](path_.length);
        amounts_[0] = amountIn_;

        for (uint256 i; i < path_.length - 1; ++i) {
            (uint256 reserveIn_, uint256 reserveOut_) = _getReserves(path_[i], path_[i + 1]);
            amounts_[i + 1] = _getAmountOut(amounts_[i], reserveIn_, reserveOut_);
        }
    }

    function _getAmountsIn(
        uint256 amountOut_,
        address[] memory path_
    ) public view returns (uint256[] memory amounts_) {
        require(path_.length >= 2, "UniswapRouterV2Mock: invalid path");

        amounts_ = new uint256[](path_.length);
        amounts_[amounts_.length - 1] = amountOut_;

        for (uint256 i = path_.length - 1; i > 0; --i) {
            (uint256 reserveIn_, uint256 reserveOut_) = _getReserves(path_[i - 1], path_[i]);
            amounts_[i - 1] = _getAmountIn(amounts_[i], reserveIn_, reserveOut_);
        }
    }

    function _swap(uint256[] memory amounts_, address[] memory path_, address receiver_) internal {
        (address tokenIn_, address tokenOut_) = (path_[0], path_[path_.length - 1]);

        if (tokenIn_ == WRAPPED_NATIVE) {
            (bool ok_, ) = WRAPPED_NATIVE.call{value: msg.value}(
                abi.encodeWithSignature("deposit()")
            );
            require(ok_, "AbstractSwapRouterMock: failed to deposit");
        } else {
            IERC20(tokenIn_).safeTransferFrom(msg.sender, address(this), amounts_[0]);
        }

        uint256 amountOut_ = amounts_[amounts_.length - 1];

        if (tokenOut_ == WRAPPED_NATIVE) {
            (bool ok_, ) = WRAPPED_NATIVE.call(
                abi.encodeWithSignature("withdraw(uint256)", amountOut_)
            );
            require(ok_, "AbstractSwapRouterMock: failed to withdraw");

            (ok_, ) = receiver_.call{value: amountOut_}("");
            require(ok_, "AbstractSwapRouterMock: failed to transfer");
        } else {
            IERC20(tokenOut_).safeTransfer(receiver_, amountOut_);
        }

        for (uint256 i = 0; i < path_.length; ++i) {
            _reserves[path_[i]] = IERC20(path_[i]).balanceOf(address(this));
        }
    }
}
