// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract MasterRouterStorage {
    bytes32 public constant MASTER_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.masterrouter.storage");

    struct MRStorage {
        address caller;
    }

    modifier onlyCaller() {
        MRStorage storage _ds = getMasterRouterStorage();

        require(_ds.caller == address(0), "MasterRouterStorage: new caller");

        _ds.caller = msg.sender;
        _;
        _ds.caller = address(0);
    }

    function getMasterRouterStorage() internal pure returns (MRStorage storage _ds) {
        bytes32 slot_ = MASTER_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }

    function getCallerAddress() public view returns (address caller_) {
        return getMasterRouterStorage().caller;
    }
}
