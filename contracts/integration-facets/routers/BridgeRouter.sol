// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@rarimo/evm-bridge/interfaces/bridge/IBridge.sol";
import "@rarimo/evm-bridge/interfaces/bundle/IBundler.sol";

import "../../libs/Approver.sol";
import "../../libs/Constants.sol";
import "../../libs/Resolver.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/BridgeRouterStorage.sol";

contract BridgeRouter is OwnableDiamondStorage, MasterRouterStorage, BridgeRouterStorage {
    using Approver for *;
    using Resolver for uint256;
    using SafeERC20 for IERC20;

    function setBridgeAddress(address bridge_) external onlyOwner {
        _getBridgeRouterStorage().bridge = bridge_;
    }

    function bridgeERC20(
        address token_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
        bool isWrapped_
    ) external payable {
        address bridge_ = getBridgeAddress();

        IERC20(token_).approveMax(bridge_);
        IBridge(bridge_).depositERC20(
            token_,
            amount_.resolve(IERC20(token_)),
            bundle_,
            network_,
            receiver_,
            isWrapped_
        );
    }

    function bridgeERC721(
        address token_,
        uint256 tokenId_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
        bool isWrapped_
    ) external payable {
        address bridge_ = getBridgeAddress();

        IERC721(token_).approveMax(bridge_);
        IBridge(bridge_).depositERC721(token_, tokenId_, bundle_, network_, receiver_, isWrapped_);
    }

    function bridgeERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
        bool isWrapped_
    ) external payable {
        address bridge_ = getBridgeAddress();

        IERC1155(token_).approveMax(bridge_);
        IBridge(bridge_).depositERC1155(
            token_,
            tokenId_,
            amount_.resolve(IERC1155(token_), tokenId_),
            bundle_,
            network_,
            receiver_,
            isWrapped_
        );
    }

    function bridgeNative(
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_
    ) external payable {
        IBridge(getBridgeAddress()).depositNative{value: amount_.resolve()}(
            bundle_,
            network_,
            receiver_
        );
    }
}
