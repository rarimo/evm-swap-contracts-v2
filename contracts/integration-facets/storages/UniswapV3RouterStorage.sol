// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract UniswapV3RouterStorage {
    bytes32 public constant UNISWAP_V3_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.uniswapv3router.storage");

    struct U3Storage {
        address swapV3Router;
    }

    function getUniswapV3RouterStorage() internal pure returns (U3Storage storage _ds) {
        bytes32 slot_ = UNISWAP_V3_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }

    function getSwapV3Router() public view returns (address swapV3Router_) {
        return getUniswapV3RouterStorage().swapV3Router;
    }
}
