// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

library BytesHelper {
    function getFirstToken(bytes calldata path_) internal pure returns (address) {
        require(path_.length > 42, "BytesHelper: invalid path length");

        return toAddress(path_, 0);
    }

    function getLastToken(bytes calldata path_) internal pure returns (address) {
        require(path_.length > 42, "BytesHelper: invalid path length");

        return toAddress(path_, path_.length - 20);
    }

    function toAddress(
        bytes memory path_,
        uint256 start_
    ) internal pure returns (address tokenAddress_) {
        assembly {
            tokenAddress_ := shr(96, mload(add(add(path_, 0x20), start_)))
        }
    }
}
