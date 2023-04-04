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
                require(
                    _makeBridge(payload_[i]) || payload_[i].skipRevert,
                    "MasterRouter: command reverted"
                );
            }
        }
    }

    function _makeBridge(Payload calldata payload_) internal returns (bool ok_) {
        uint256 command_ = payload_.command;
        bytes4 funcSelector_;
        address bridgeFacet_;

        if (command_ == Commands.BRIDGE_ERC20) {
            funcSelector_ = BridgeRouter.bridgeERC20.selector;
            bridgeFacet_ = getFacetBySelector(funcSelector_);
        } else if (command_ == Commands.BRIDGE_ERC721) {
            funcSelector_ = BridgeRouter.bridgeERC721.selector;
            bridgeFacet_ = getFacetBySelector(funcSelector_);
        } else if (command_ == Commands.BRIDGE_ERC1155) {} else if (
            command_ == Commands.BRIDGE_NATIVE
        ) {}

        (ok_, ) = bridgeFacet_.delegatecall(abi.encodePacked(funcSelector_, payload_.data));
    }
}
