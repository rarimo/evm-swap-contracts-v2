// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract BridgeRouterStorage {
    bytes32 public constant BRIDGE_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.bridgerouter.storage");

    struct BRStorage {
        address bridge;
    }

    function getBridgeRouterStorage() internal pure returns (BRStorage storage _ds) {
        bytes32 slot_ = BRIDGE_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }

    function getBridgeAddress() public view returns (address bridge_) {
        return getBridgeRouterStorage().bridge;
    }
}
