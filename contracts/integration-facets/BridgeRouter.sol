//SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "@rarimo/evm-bridge/interfaces/bridge/IBridge.sol";
import "@rarimo/evm-bridge/interfaces/bundle/IBundler.sol";

import "../libs/Approver.sol";

import "../master-facet/MasterRouterStorage.sol";
import "./BridgeRouterStorage.sol";

contract BridgeRouter is OwnableDiamondStorage, MasterRouterStorage, BridgeRouterStorage {
    using Approver for *;
    using SafeERC20 for IERC20;

    function setBridgeAddress(address bridge_) external onlyOwner {
        BRStorage storage _ds = getBridgeRouterStorage();

        _ds.bridge = bridge_;
    }

    function bridgeERC20(
        bool callerPayer_,
        address token_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
        bool isWrapped_
    ) external {
        address bridge_ = getBridgeAddress();

        if (callerPayer_) {
            IERC20(token_).safeTransferFrom(getCallerAddress(), address(this), amount_);
        }

        IERC20(token_).approveMax(bridge_);
        IBridge(bridge_).depositERC20(token_, amount_, bundle_, network_, receiver_, isWrapped_);
    }
}
