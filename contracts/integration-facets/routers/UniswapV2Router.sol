// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol";

import "../../libs/Approver.sol";
import "../../libs/Resolver.sol";
import "../../libs/Constants.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/UniswapV2RouterStorage.sol";
import "./TransferRouter.sol";

contract UniswapV2Router is
    OwnableDiamondStorage,
    MasterRouterStorage,
    UniswapV2RouterStorage,
    TransferRouter
{
    using SafeERC20 for IERC20;
    using Approver for *;
    using Resolver for address;

    function setUniswapV2RouterAddress(address swapV2Router_) external onlyOwner {
        _getUniswapV2RouterStorage().swapV2Router = swapV2Router_;
    }

    function swapExactTokensForTokensV2(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address swapV2router_ = getSwapV2Router();

        IERC20(path_[0]).approveMax(swapV2router_);
        IUniswapV2Router01(swapV2router_).swapExactTokensForTokens(
            amountIn_,
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactTokensV2(
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address tokenIn_ = path_[0];
        address swapV2router_ = getSwapV2Router();

        IERC20(tokenIn_).approveMax(swapV2router_);
        uint256 spentFundsAmount_ = IUniswapV2Router01(swapV2router_).swapTokensForExactTokens(
            amountOut_,
            amountInMax_,
            path_,
            receiver_.resolve(),
            block.timestamp
        )[0];

        if (amountInMax_ > spentFundsAmount_) {
            transferERC20(tokenIn_, Constants.CALLER_ADDRESS, amountInMax_ - spentFundsAmount_);
        }
    }

    function swapExactETHForTokens(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        IUniswapV2Router01(getSwapV2Router()).swapExactETHForTokens{value: amountIn_}(
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactETH(
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address tokenIn_ = path_[0];
        address swapV2router_ = getSwapV2Router();

        IERC20(tokenIn_).approveMax(swapV2router_);
        uint256 spentFundsAmount_ = IUniswapV2Router01(swapV2router_).swapTokensForExactETH(
            amountOut_,
            amountInMax_,
            path_,
            receiver_.resolve(),
            block.timestamp
        )[0];

        if (amountInMax_ > spentFundsAmount_) {
            transferERC20(tokenIn_, Constants.CALLER_ADDRESS, amountInMax_ - spentFundsAmount_);
        }
    }

    function swapExactTokensForETH(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address swapV2router_ = getSwapV2Router();

        IERC20(path_[0]).approveMax(swapV2router_);
        IUniswapV2Router01(swapV2router_).swapExactTokensForETH(
            amountIn_,
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapETHForExactTokens(
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        uint256 spentFundsAmount_ = IUniswapV2Router01(getSwapV2Router()).swapETHForExactTokens{
            value: amountInMax_
        }(amountOut_, path_, receiver_.resolve(), block.timestamp)[0];

        if (amountInMax_ > spentFundsAmount_) {
            transferNative(Constants.CALLER_ADDRESS, amountInMax_ - spentFundsAmount_);
        }
    }

    function _validatePath(address[] calldata path_) internal pure {
        require(path_.length >= 2, "UniswapV2Router: invalid path");
    }
}
