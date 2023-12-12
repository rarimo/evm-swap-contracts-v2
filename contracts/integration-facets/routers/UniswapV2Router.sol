// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {OwnableDiamondStorage} from "@solarity/solidity-lib/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import {IUniswapV2Router01} from "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol";

import {Approver} from "../../libs/Approver.sol";
import {Resolver} from "../../libs/Resolver.sol";
import {UniswapV2RouterStorage} from "../storages/UniswapV2RouterStorage.sol";

contract UniswapV2Router is OwnableDiamondStorage, UniswapV2RouterStorage {
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

        address swapV2router_ = getSwapV2Router();

        IERC20(path_[0]).approveMax(swapV2router_);
        IUniswapV2Router01(swapV2router_).swapTokensForExactTokens(
            amountOut_,
            amountInMax_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
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

        address swapV2router_ = getSwapV2Router();

        IERC20(path_[0]).approveMax(swapV2router_);
        IUniswapV2Router01(swapV2router_).swapTokensForExactETH(
            amountOut_,
            amountInMax_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
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

        IUniswapV2Router01(getSwapV2Router()).swapETHForExactTokens{value: amountInMax_}(
            amountOut_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function _validatePath(address[] calldata path_) internal pure {
        require(path_.length >= 2, "UniswapV2Router: invalid path");
    }
}
