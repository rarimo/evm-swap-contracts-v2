// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../../interfaces/tokens/IWrappedNative.sol";

abstract contract AbstractSwapRouterMock {
    using SafeERC20 for IERC20;

    address public immutable WRAPPED_NATIVE;

    mapping(address => uint256) internal _reserves;
    mapping(address => mapping(address => address)) internal _pairs;

    constructor(address wrappedNative_) {
        WRAPPED_NATIVE = wrappedNative_;
    }

    receive() external payable {}

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

        for (uint256 i = 0; i < path_.length - 1; ++i) {
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

    function _exactIn(
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] memory path_,
        address receiver_,
        bool toNative_
    ) internal returns (uint256[] memory amounts_) {
        amounts_ = _getAmountsOut(amountIn_, path_);

        require(
            amounts_[amounts_.length - 1] >= amountOutMin_,
            "AbstractSwapRouterMock: insufficient amount out"
        );

        _swap(amounts_, path_, receiver_, toNative_);
    }

    function _exactOut(
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] memory path_,
        address receiver_,
        bool toNative_
    ) internal returns (uint256[] memory amounts_) {
        amounts_ = _getAmountsIn(amountOut_, path_);

        require(amounts_[0] <= amountInMax_, "AbstractSwapRouterMock: excessive input amount");

        _swap(amounts_, path_, receiver_, toNative_);
    }

    function _swap(
        uint256[] memory amounts_,
        address[] memory path_,
        address receiver_,
        bool toNative_
    ) internal {
        (address tokenIn_, address tokenOut_) = (path_[0], path_[path_.length - 1]);
        (uint256 amountIn_, uint256 amountOut_) = (amounts_[0], amounts_[amounts_.length - 1]);

        if (tokenIn_ != WRAPPED_NATIVE || msg.value == 0) {
            IERC20(tokenIn_).safeTransferFrom(msg.sender, address(this), amountIn_);
        } else {
            IWrappedNative(WRAPPED_NATIVE).deposit{value: amountIn_}();
        }

        if (tokenOut_ == WRAPPED_NATIVE && toNative_) {
            IWrappedNative(WRAPPED_NATIVE).withdraw(amountOut_);

            (bool ok_, ) = receiver_.call{value: amountOut_}("");
            require(ok_, "AbstractSwapRouterMock: failed to transfer native");
        } else {
            IERC20(tokenOut_).safeTransfer(receiver_, amountOut_);
        }

        _reserves[tokenIn_] = IERC20(tokenIn_).balanceOf(address(this));
        _reserves[tokenOut_] = IERC20(tokenOut_).balanceOf(address(this));
    }
}
