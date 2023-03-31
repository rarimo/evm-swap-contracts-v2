//SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

library Commands {
    /// @dev bridge commands 1 <= command <= 9
    uint256 internal constant BRIDGE_ERC20 = 1;
    uint256 internal constant BRIDGE_ERC721 = 2;
    uint256 internal constant BRIDGE_ERC1155 = 3;
    uint256 internal constant BRIDGE_NATIVE = 4;
}
