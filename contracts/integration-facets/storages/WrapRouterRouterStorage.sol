// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract WrapRouterStorage {
    bytes32 public constant WRAP_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.wraprouter.storage");

    struct WRStorage {
        address wrappedNative;
    }

    function getWrappedNativeAddress() public view returns (address wrappedNative_) {
        return _getWrapRouterStorage().wrappedNative;
    }

    function _getWrapRouterStorage() internal pure returns (WRStorage storage _ds) {
        bytes32 slot_ = WRAP_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }
}
