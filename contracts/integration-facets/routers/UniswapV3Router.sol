// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {OwnableDiamondStorage} from "@solarity/solidity-lib/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import {ISwapRouter} from "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";
import {IPeripheryPayments} from "@uniswap/v3-periphery/contracts/interfaces/IPeripheryPayments.sol";

import {BytesHelper} from "../../libs/BytesHelper.sol";
import {Approver} from "../../libs/Approver.sol";
import {Resolver} from "../../libs/Resolver.sol";
import {UniswapV3RouterStorage} from "../storages/UniswapV3RouterStorage.sol";

contract UniswapV3Router is OwnableDiamondStorage, UniswapV3RouterStorage {
    using SafeERC20 for IERC20;
    using BytesHelper for bytes;
    using Approver for *;
    using Resolver for address;

    function setUniswapV3RouterAddress(address swapV3Router_) external onlyOwner {
        _getUniswapV3RouterStorage().swapV3Router = swapV3Router_;
    }

    function exactInput(
        bool isNative_,
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMinimum_,
        bytes calldata path_
    ) external payable {
        address swapV3Router_ = getSwapV3Router();

        if (!isNative_) {
            IERC20(path_.getFirstToken()).approveMax(swapV3Router_);
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
        bool isNative_,
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMaximum_,
        bytes calldata path_
    ) external payable {
        address swapV3Router_ = getSwapV3Router();

        if (!isNative_) {
            IERC20(path_.getLastToken()).approveMax(swapV3Router_);
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

        if (isNative_ && amountInMaximum_ > spentFundsAmount_) {
            IPeripheryPayments(swapV3Router_).refundETH();
        }
    }
}
