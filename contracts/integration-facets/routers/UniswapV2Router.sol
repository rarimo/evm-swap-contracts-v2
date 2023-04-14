// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol";

import "../../libs/Approver.sol";
import "../../libs/Payer.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/UniswapV2RouterStorage.sol";

contract UniswapV2Router is OwnableDiamondStorage, MasterRouterStorage, UniswapV2RouterStorage {
    using SafeERC20 for IERC20;
    using Approver for *;
    using Payer for *;
    using Resolver for address;

    function setUniswapV2RouterAddress(address swapV2Router_) external onlyOwner {
        getUniswapV2RouterStorage().swapV2Router = swapV2Router_;
    }

    function swapExactTokensForTokens(
        bool callerPayer_,
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "UniswapV2Router: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address swapV2router_ = getSwapV2Router();

        IERC20(tokenIn_).approveMax(swapV2router_);
        IUniswapV2Router01(swapV2router_).swapExactTokensForTokens(
            amountIn_,
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactTokens(
        bool callerPayer_,
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "UniswapV2Router: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountInMax_);
        }

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
            IERC20(tokenIn_).pay(receiver_, amountInMax_ - spentFundsAmount_);
        }
    }

    function swapExactETHForTokens(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        IUniswapV2Router01(getSwapV2Router()).swapExactETHForTokens{value: amountIn_}(
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactETH(
        bool callerPayer_,
        address receiver_,
        address changeReceiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "UniswapV2Router: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountInMax_);
        }

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
            IERC20(tokenIn_).pay(changeReceiver_, amountInMax_ - spentFundsAmount_);
        }
    }

    function swapExactTokensForETH(
        bool callerPayer_,
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "UniswapV2Router: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address swapV2router_ = getSwapV2Router();

        IERC20(tokenIn_).approveMax(swapV2router_);
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
    ) external {
        uint256 spentFundsAmount_ = IUniswapV2Router01(getSwapV2Router()).swapETHForExactTokens{
            value: amountInMax_
        }(amountOut_, path_, receiver_.resolve(), block.timestamp)[0];

        if (amountInMax_ > spentFundsAmount_) {
            receiver_.pay(amountInMax_ - spentFundsAmount_);
        }
    }
}
