// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";
import "@uniswap/v3-periphery/contracts/interfaces/IPeripheryPayments.sol";

import "../../libs/Constants.sol";
import "../../libs/Approver.sol";
import "../../libs/BytesHelper.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/UniswapV3RouterStorage.sol";

contract UniswapV3Router is OwnableDiamondStorage, MasterRouterStorage, UniswapV3RouterStorage {
    using BytesHelper for bytes;
    using Approver for *;
    using SafeERC20 for IERC20;

    function setUniswapV3RouterAddress(address swapV3Router_) external onlyOwner {
        getUniswapV3RouterStorage().swapV3Router = swapV3Router_;
    }

    function exactInput(
        bool callerPayer_,
        bool isNative_,
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
                recipient: address(this),
                deadline: block.timestamp,
                amountIn: amountIn_,
                amountOutMinimum: amountOutMinimum_
            })
        );
    }

    function exactOutput(
        bool callerPayer_,
        bool isNative_,
        address changeReceiver_,
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

        uint256 spentFundsAmount = ISwapRouter(swapV3Router_).exactOutput{
            value: isNative_ ? amountInMaximum_ : 0
        }(
            ISwapRouter.ExactOutputParams({
                path: path_,
                recipient: address(this),
                deadline: block.timestamp,
                amountOut: amountOut_,
                amountInMaximum: amountInMaximum_
            })
        );

        if (isNative_) {
            IPeripheryPayments(swapV3Router_).refundETH();

            _payNativeChange(changeReceiver_, amountInMaximum_ - spentFundsAmount);
        } else {
            _payERC20Change(tokenIn_, changeReceiver_, amountInMaximum_ - spentFundsAmount);
        }
    }

    function _payERC20Change(address token_, address changeReceiver_, uint256 amount_) internal {
        if (changeReceiver_ == Constants.CALLER_ADDRESS) {
            IERC20(token_).safeTransfer(getCallerAddress(), amount_);
        } else if (changeReceiver_ != Constants.THIS_ADDRESS) {
            IERC20(token_).safeTransfer(changeReceiver_, amount_);
        }
    }

    function _payNativeChange(address changeReceiver_, uint256 amount_) internal {
        if (changeReceiver_ == Constants.THIS_ADDRESS) {
            return;
        }

        bool ok_;
        if (changeReceiver_ == Constants.CALLER_ADDRESS) {
            (ok_, ) = getCallerAddress().call{value: amount_}("");
        } else {
            (ok_, ) = changeReceiver_.call{value: amount_}("");
        }

        require(ok_, "UniswapV3Router: failed to refund native");
    }
}
