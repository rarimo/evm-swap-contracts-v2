// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract WrapRouterStorage {
    bytes32 public constant WRAP_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.wraprouter.storage");

    struct WRStorage {
        address weth9;
    }

    function getWETH9Address() public view returns (address weth9_) {
        return _getWrapRouterStorage().weth9;
    }

    function _getWrapRouterStorage() internal pure returns (WRStorage storage _ds) {
        bytes32 slot_ = WRAP_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }
}
