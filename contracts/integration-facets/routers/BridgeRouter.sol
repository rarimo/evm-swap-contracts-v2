// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@rarimo/evm-bridge/interfaces/facade/IBridgeFacade.sol";
import "@rarimo/evm-bridge/interfaces/bundle/IBundler.sol";
import {Constants as BridgeConstants} from "@rarimo/evm-bridge/libs/Constants.sol";

import "../../libs/Approver.sol";
import "../../libs/Constants.sol";
import "../../libs/Resolver.sol";
import "../storages/BridgeRouterStorage.sol";

contract BridgeRouter is OwnableDiamondStorage, BridgeRouterStorage {
    using Approver for *;
    using Resolver for uint256;
    using SafeERC20 for IERC20;

    function setBridgeAddress(address bridge_) external onlyOwner {
        _getBridgeRouterStorage().bridge = bridge_;
    }

    function bridgeERC20(
        IBridgeFacade.DepositFeeERC20Parameters calldata feeParams_,
        IBridge.DepositERC20Parameters memory depositParams_
    ) external payable {
        address bridge_ = getBridgeAddress();
        (uint256 nativeFee_, uint256 erc20Fee_) = _approveFee(bridge_, feeParams_.feeToken);

        depositParams_.amount = depositParams_.amount.resolve(
            IERC20(depositParams_.token),
            feeParams_.feeToken == depositParams_.token ? erc20Fee_ : 0
        );

        IERC20(depositParams_.token).approveMax(bridge_);

        IBridgeFacade(bridge_).depositERC20{value: nativeFee_}(feeParams_, depositParams_);
    }

    function bridgeERC721(
        IBridgeFacade.DepositFeeERC721Parameters calldata feeParams_,
        IBridge.DepositERC721Parameters calldata depositParams_
    ) external payable {
        address bridge_ = getBridgeAddress();
        (uint256 nativeFee_, ) = _approveFee(bridge_, feeParams_.feeToken);

        IERC721(depositParams_.token).approveMax(bridge_);

        IBridgeFacade(bridge_).depositERC721{value: nativeFee_}(feeParams_, depositParams_);
    }

    function bridgeERC1155(
        IBridgeFacade.DepositFeeERC1155Parameters calldata feeParams_,
        IBridge.DepositERC1155Parameters memory depositParams_
    ) external payable {
        address bridge_ = getBridgeAddress();
        (uint256 nativeFee_, ) = _approveFee(bridge_, feeParams_.feeToken);

        depositParams_.amount = depositParams_.amount.resolve(
            IERC1155(depositParams_.token),
            depositParams_.tokenId
        );

        IERC1155(depositParams_.token).approveMax(bridge_);

        IBridgeFacade(bridge_).depositERC1155{value: nativeFee_}(feeParams_, depositParams_);
    }

    function bridgeNative(
        IBridgeFacade.DepositFeeNativeParameters calldata feeParams_,
        IBridge.DepositNativeParameters memory depositParams_
    ) external payable {
        address bridge_ = getBridgeAddress();

        (uint256 nativeFee_, ) = _approveFee(bridge_, feeParams_.feeToken);

        depositParams_.amount = depositParams_.amount.resolve(nativeFee_);

        IBridgeFacade(bridge_).depositNative{value: depositParams_.amount + nativeFee_}(
            feeParams_,
            depositParams_
        );
    }

    function _approveFee(
        address bridge_,
        address feeToken_
    ) private returns (uint256 nativeFee_, uint256 erc20Fee_) {
        if (feeToken_ == BridgeConstants.ETHEREUM_ADDRESS) {
            nativeFee_ = IFeeManager(bridge_).getCommission(feeToken_);
        } else if (feeToken_ != BridgeConstants.COMMISSION_ADDRESS) {
            erc20Fee_ = IFeeManager(bridge_).getCommission(feeToken_);

            IERC20(feeToken_).approveMax(bridge_);
        }
    }
}
