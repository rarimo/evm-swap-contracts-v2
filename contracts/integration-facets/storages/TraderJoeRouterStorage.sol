// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract TraderJoeRouterStorage {
    bytes32 public constant TRADER_JOE_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.traderjoerouter.storage");

    struct TJStorage {
        address traderJoeRouter;
    }

    function getTraderJoeRouterStorage() internal pure returns (TJStorage storage _ds) {
        bytes32 slot_ = TRADER_JOE_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }

    function getTraderJoeRouter() public view returns (address traderJoeRouter_) {
        return getTraderJoeRouterStorage().traderJoeRouter;
    }
}
