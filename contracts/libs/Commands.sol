// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

library Commands {
    /// @dev bridge facet commands: 1 <= command <= 9
    uint256 internal constant BRIDGE_ERC20 = 1;
    uint256 internal constant BRIDGE_ERC721 = 2;
    uint256 internal constant BRIDGE_ERC1155 = 3;
    uint256 internal constant BRIDGE_NATIVE = 4;

    /// @dev transfer facet commands: 10 <= command <= 19
    uint256 internal constant TRANSFER_ERC20 = 10;
    uint256 internal constant TRANSFER_ERC721 = 11;
    uint256 internal constant TRANSFER_ERC1155 = 12;
    uint256 internal constant TRANSFER_NATIVE = 13;

    /// @dev wrap facet commands: 20 <= command <= 24
    uint256 internal constant WRAP_NATIVE = 20;
    uint256 internal constant UNWRAP_NATIVE = 21;

    /// @dev multicall facet commands: 25 <= command <= 29
    uint256 internal constant MULTICALL = 25;

    /// @dev UniswapV2 facet commands: 50 <= command <= 59
    uint256 internal constant UV2_SWAP_EXACT_TOKENS_FOR_TOKENS = 50;
    uint256 internal constant UV2_SWAP_TOKENS_FOR_EXACT_TOKENS = 51;
    uint256 internal constant UV2_SWAP_EXACT_ETH_FOR_TOKENS = 52;
    uint256 internal constant UV2_SWAP_TOKENS_FOR_EXACT_ETH = 53;
    uint256 internal constant UV2_SWAP_EXACT_TOKENS_FOR_ETH = 54;
    uint256 internal constant UV2_SWAP_ETH_FOR_EXACT_TOKENS = 55;

    /// @dev UniswapV3 facet commands: 60 <= command <= 69
    uint256 internal constant EXACT_INPUT = 60;
    uint256 internal constant EXACT_OUTPUT = 61;

    /// @dev TraderJoe facet commands: 70 <= command <= 79
    uint256 internal constant TJ_SWAP_EXACT_TOKENS_FOR_TOKENS = 70;
    uint256 internal constant TJ_SWAP_TOKENS_FOR_EXACT_TOKENS = 71;
    uint256 internal constant TJ_SWAP_EXACT_AVAX_FOR_TOKENS = 72;
    uint256 internal constant TJ_SWAP_TOKENS_FOR_EXACT_AVAX = 73;
    uint256 internal constant TJ_SWAP_EXACT_TOKENS_FOR_AVAX = 74;
    uint256 internal constant TJ_SWAP_AVAX_FOR_EXACT_TOKENS = 75;
}
