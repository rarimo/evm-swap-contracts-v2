// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@traderjoe-xyz/core/contracts/traderjoe/interfaces/IJoeRouter01.sol";

import "../../libs/Approver.sol";
import "../../libs/Payer.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/TraderJoeRouterStorage.sol";

contract TraderJoeRouter is OwnableDiamondStorage, MasterRouterStorage, TraderJoeRouterStorage {
    using SafeERC20 for IERC20;
    using Approver for *;
    using Payer for *;
    using Resolver for address;

    function setTraderJoeRouterAddress(address traderJoeRouter_) external onlyOwner {
        getTraderJoeRouterStorage().traderJoeRouter = traderJoeRouter_;
    }

    function swapExactTokensForTokens(
        bool callerPayer_,
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn_).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapExactTokensForTokens(
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
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountInMax_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn_).approveMax(traderJoeRouter_);
        uint256 spentFundsAmount_ = IJoeRouter01(traderJoeRouter_).swapTokensForExactTokens(
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

    function swapExactAVAXForTokens(
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        IJoeRouter01(getTraderJoeRouter()).swapExactAVAXForTokens{value: amountIn_}(
            amountOutMin_,
            path_,
            receiver_.resolve(),
            block.timestamp
        );
    }

    function swapTokensForExactAVAX(
        bool callerPayer_,
        address receiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountInMax_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn_).approveMax(traderJoeRouter_);
        uint256 spentFundsAmount_ = IJoeRouter01(traderJoeRouter_).swapTokensForExactAVAX(
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

    function swapExactTokensForAVAX(
        bool callerPayer_,
        address receiver_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn_ = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn_).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn_).approveMax(traderJoeRouter_);
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
    ) external {
        uint256 spentFundsAmount_ = IJoeRouter01(getTraderJoeRouter()).swapAVAXForExactTokens{
            value: amountInMax_
        }(amountOut_, path_, receiver_.resolve(), block.timestamp)[0];

        if (amountInMax_ > spentFundsAmount_) {
            receiver_.pay(amountInMax_ - spentFundsAmount_);
        }
    }
}
