// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {OwnableDiamondStorage} from "@solarity/solidity-lib/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import {IJoeRouter01} from "@traderjoe-xyz/core/contracts/traderjoe/interfaces/IJoeRouter01.sol";

import {Approver} from "../../libs/Approver.sol";
import {Resolver} from "../../libs/Resolver.sol";
import {TraderJoeRouterStorage} from "../storages/TraderJoeRouterStorage.sol";

contract TraderJoeRouter is OwnableDiamondStorage, TraderJoeRouterStorage {
    using SafeERC20 for IERC20;
    using Approver for *;
    using Resolver for address;

    function setTraderJoeRouterAddress(address traderJoeRouter_) external onlyOwner {
        _getTraderJoeRouterStorage().traderJoeRouter = traderJoeRouter_;
    }

    function swapExactTokensForTokensTJ(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(path_[0]).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapExactTokensForTokens(
            amountIn_,
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactTokensTJ(
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(path_[0]).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapTokensForExactTokens(
            amountOut_,
            amountInMax_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapExactAVAXForTokens(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        IJoeRouter01(getTraderJoeRouter()).swapExactAVAXForTokens{value: amountIn_}(
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactAVAX(
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(path_[0]).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapTokensForExactAVAX(
            amountOut_,
            amountInMax_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapExactTokensForAVAX(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(path_[0]).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapExactTokensForAVAX(
            amountIn_,
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapAVAXForExactTokens(
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external payable {
        _validatePath(path_);

        IJoeRouter01(getTraderJoeRouter()).swapAVAXForExactTokens{value: amountInMax_}(
            amountOut_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function _validatePath(address[] calldata path_) internal pure {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");
    }
}
