// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";
import "@uniswap/v3-periphery/contracts/interfaces/IPeripheryPayments.sol";

import "../../libs/BytesHelper.sol";
import "../../libs/Approver.sol";
import "../../libs/Resolver.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/UniswapV3RouterStorage.sol";
import "./TransferRouter.sol";

contract UniswapV3Router is
    OwnableDiamondStorage,
    MasterRouterStorage,
    UniswapV3RouterStorage,
    TransferRouter
{
    using SafeERC20 for IERC20;
    using BytesHelper for bytes;
    using Approver for *;
    using Resolver for address;

    function setUniswapV3RouterAddress(address swapV3Router_) external onlyOwner {
        _getUniswapV3RouterStorage().swapV3Router = swapV3Router_;
    }

    function exactInput(
        bool callerPayer_,
        bool isNative_,
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMinimum_,
        bytes calldata path_
    ) external {
        address tokenIn_ = path_.getFirstToken();

        if (callerPayer_ && !isNative_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address swapV3Router_ = getSwapV3Router();

        if (!isNative_) {
            IERC20(tokenIn_).approveMax(swapV3Router_);
        }

        ISwapRouter(swapV3Router_).exactInput{value: isNative_ ? amountIn_ : 0}(
            ISwapRouter.ExactInputParams({
                path: path_,
                recipient: receiver_.resolve(),
                deadline: block.timestamp,
                amountIn: amountIn_,
                amountOutMinimum: amountOutMinimum_
            })
        );
    }

    function exactOutput(
        bool callerPayer_,
        bool isNative_,
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMaximum_,
        bytes calldata path_
    ) external {
        address tokenIn_ = path_.getLastToken();

        if (callerPayer_ && !isNative_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountInMaximum_);
        }

        address swapV3Router_ = getSwapV3Router();

        if (!isNative_) {
            IERC20(tokenIn_).approveMax(swapV3Router_);
        }

        uint256 spentFundsAmount_ = ISwapRouter(swapV3Router_).exactOutput{
            value: isNative_ ? amountInMaximum_ : 0
        }(
            ISwapRouter.ExactOutputParams({
                path: path_,
                recipient: receiver_.resolve(),
                deadline: block.timestamp,
                amountOut: amountOut_,
                amountInMaximum: amountInMaximum_
            })
        );

        if (isNative_) {
            IPeripheryPayments(swapV3Router_).refundETH();

            transferNative(receiver_, amountInMaximum_ - spentFundsAmount_);
        } else {
            transferERC20(tokenIn_, receiver_, amountInMaximum_ - spentFundsAmount_);
        }
    }
}
