//SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@dlsl/dev-modules/diamond/DiamondStorage.sol";

import "../libs/Commands.sol";

import "../integration-facets/BridgeRouter.sol";
import "./MasterRouterStorage.sol";

contract MasterRouter is DiamondStorage, MasterRouterStorage {
    struct Payload {
        uint256 command;
        bool skipRevert;
        bool callerPayer;
        bytes data;
    }

    function make(Payload[] calldata payload_) external payable onlyCaller {
        for (uint256 i = 0; i < payload_.length; ++i) {
            uint256 command_ = payload_[i].command;

            if (command_ >= 1 && command_ <= 9) {
                _makeBridge(payload_[i]);
            }
        }
    }

    function _makeBridge(Payload calldata payload_) internal {
        uint256 command_ = payload_.command;

        if (command_ == Commands.BRIDGE_ERC20) {
            bytes4 funcSelector_ = BridgeRouter.bridgeERC20.selector;
            address bridgeFacet_ = getFacetBySelector(funcSelector_);

            (bool ok_, ) = bridgeFacet_.delegatecall(
                abi.encodePacked(funcSelector_, payload_.data)
            );

            require(ok_ || payload_.skipRevert, "MasterRouter: command reverted");
        } else if (command_ == Commands.BRIDGE_ERC721) {} else if (
            command_ == Commands.BRIDGE_ERC1155
        ) {} else if (command_ == Commands.BRIDGE_NATIVE) {}
    }
}
