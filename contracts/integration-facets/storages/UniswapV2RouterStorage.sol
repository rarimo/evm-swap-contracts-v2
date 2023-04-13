// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract UniswapV2RouterStorage {
    bytes32 public constant UNISWAP_V2_ROUTER_STORAGE_SLOT =
        keccak256("diamond.standard.uniswapv2router.storage");

    struct U2Storage {
        address swapV2Router;
    }

    function getUniswapV2RouterStorage() internal pure returns (U2Storage storage _ds) {
        bytes32 slot_ = UNISWAP_V2_ROUTER_STORAGE_SLOT;

        assembly {
            _ds.slot := slot_
        }
    }

    function getSwapV2Router() public view returns (address swapV2Router_) {
        return getUniswapV2RouterStorage().swapV2Router;
    }
}
