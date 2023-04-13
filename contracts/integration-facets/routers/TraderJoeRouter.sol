// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@traderjoe-xyz/core/contracts/traderjoe/interfaces/IJoeRouter01.sol";

import "../../libs/Constants.sol";
import "../../libs/Approver.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/TraderJoeRouterStorage.sol";

contract TraderJoeRouter is OwnableDiamondStorage, MasterRouterStorage, TraderJoeRouterStorage {
    using Approver for *;
    using SafeERC20 for IERC20;

    function setTraderJoeRouterAddress(address traderJoeRouter_) external onlyOwner {
        getTraderJoeRouterStorage().traderJoeRouter = traderJoeRouter_;
    }

    function swapExactTokensForTokens(
        bool callerPayer_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapExactTokensForTokens(
            amountIn_,
            amountOutMin_,
            path_,
            address(this),
            block.timestamp
        );
    }

    function swapTokensForExactTokens(
        bool callerPayer_,
        address changeReceiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn).safeTransferFrom(getCallerAddress(), address(this), amountInMax_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn).approveMax(traderJoeRouter_);
        uint256 spentFundsAmount_ = IJoeRouter01(traderJoeRouter_).swapTokensForExactTokens(
            amountOut_,
            amountInMax_,
            path_,
            address(this),
            block.timestamp
        )[0];

        if (amountInMax_ > spentFundsAmount_) {
            _payERC20Change(tokenIn, changeReceiver_, amountInMax_ - spentFundsAmount_);
        }
    }

    function swapExactAVAXForTokens(
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        IJoeRouter01(getTraderJoeRouter()).swapExactAVAXForTokens{value: amountIn_}(
            amountOutMin_,
            path_,
            address(this),
            block.timestamp
        );
    }

    function swapTokensForExactAVAX(
        bool callerPayer_,
        address changeReceiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn).safeTransferFrom(getCallerAddress(), address(this), amountInMax_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn).approveMax(traderJoeRouter_);
        uint256 spentFundsAmount_ = IJoeRouter01(traderJoeRouter_).swapTokensForExactAVAX(
            amountOut_,
            amountInMax_,
            path_,
            address(this),
            block.timestamp
        )[0];

        if (amountInMax_ > spentFundsAmount_) {
            _payERC20Change(tokenIn, changeReceiver_, amountInMax_ - spentFundsAmount_);
        }
    }

    function swapExactTokensForAVAX(
        bool callerPayer_,
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_
    ) external {
        require(path_.length >= 2, "TraderJoeRouter: invalid path");

        address tokenIn = path_[0];

        if (callerPayer_) {
            IERC20(tokenIn).safeTransferFrom(getCallerAddress(), address(this), amountIn_);
        }

        address traderJoeRouter_ = getTraderJoeRouter();

        IERC20(tokenIn).approveMax(traderJoeRouter_);
        IJoeRouter01(traderJoeRouter_).swapExactTokensForAVAX(
            amountIn_,
            amountOutMin_,
            path_,
            address(this),
            block.timestamp
        );
    }

    function swapAVAXForExactTokens(
        address changeReceiver_,
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_
    ) external {
        uint256 spentFundsAmount_ = IJoeRouter01(getTraderJoeRouter()).swapAVAXForExactTokens{
            value: amountInMax_
        }(amountOut_, path_, address(this), block.timestamp)[0];

        if (amountInMax_ > spentFundsAmount_) {
            _payNativeChange(changeReceiver_, amountInMax_ - spentFundsAmount_);
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

        require(ok_, "TraderJoeRouter: failed to refund native");
    }
}
