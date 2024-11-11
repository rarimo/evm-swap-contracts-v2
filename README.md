# EVM Swap Contracts V2

**ERC-2535 diamond-based batcher and router smart contract**

- UniswapV3 swapping integration.
- UniswapV2 ERC20 and native swapping support.
- TraderJoe (avalanche) swapping support.
- Multicall router.
- Native wrapping/unwrapping functionality.
- ERC20/ERC721/ERC1155/native transferring batcher.
- Full Rarimo Bridge integration.

## Usage

The router solution may be deployed with the following config:

```json
{
  "accounts": {
    "owner": "0x0000000000000000000000000000000000000001"
  },
  "integrations": {
    "bridgeFacade": "0x0000000000000000000000000000000000000002",
    "wrappedNative": "0x0000000000000000000000000000000000000003",
    "uniswapV2Router": "0x0000000000000000000000000000000000000004",
    "uniswapV3Router": "0x0000000000000000000000000000000000000005",
    "traderJoeRouter": "0x0000000000000000000000000000000000000006"
  },
  "facets": {
    "bridge": true,
    "multicall": true,
    "transfer": true,
    "wrap": true,
    "uniswapV2": true,
    "uniswapV3": true,
    "traderJoe": true
  }
}
```

The supported facets may be "batch-executed" via the `MasterRouter` facet `make()` function.

Here is the list of supported commands:

```solidity
/// @dev bridge facet commands: 1 <= command <= 9
uint256 constant BRIDGE_ERC20 = 1;
uint256 constant BRIDGE_ERC721 = 2;
uint256 constant BRIDGE_ERC1155 = 3;
uint256 constant BRIDGE_NATIVE = 4;

/// @dev transfer facet commands: 10 <= command <= 19
uint256 constant TRANSFER_ERC20 = 10;
uint256 constant TRANSFER_ERC721 = 11;
uint256 constant TRANSFER_ERC1155 = 12;
uint256 constant TRANSFER_NATIVE = 13;
uint256 constant TRANSFER_FROM_ERC20 = 14;
uint256 constant TRANSFER_FROM_ERC721 = 15;
uint256 constant TRANSFER_FROM_ERC1155 = 16;

/// @dev wrap facet commands: 20 <= command <= 24
uint256 constant WRAP_NATIVE = 20;
uint256 constant UNWRAP_NATIVE = 21;

/// @dev multicall facet commands: 25 <= command <= 29
uint256 constant MULTICALL = 25;

/// @dev UniswapV2 facet commands: 50 <= command <= 59
uint256 constant SWAP_EXACT_TOKENS_FOR_TOKENS_V2 = 50;
uint256 constant SWAP_TOKENS_FOR_EXACT_TOKENS_V2 = 51;
uint256 constant SWAP_EXACT_ETH_FOR_TOKENS = 52;
uint256 constant SWAP_TOKENS_FOR_EXACT_ETH = 53;
uint256 constant SWAP_EXACT_TOKENS_FOR_ETH = 54;
uint256 constant SWAP_ETH_FOR_EXACT_TOKENS = 55;

/// @dev UniswapV3 facet commands: 60 <= command <= 69
uint256 constant EXACT_INPUT = 60;
uint256 constant EXACT_OUTPUT = 61;

/// @dev TraderJoe facet commands: 70 <= command <= 79
uint256 constant SWAP_EXACT_TOKENS_FOR_TOKENS_TJ = 70;
uint256 constant SWAP_TOKENS_FOR_EXACT_TOKENS_TJ = 71;
uint256 constant SWAP_EXACT_AVAX_FOR_TOKENS = 72;
uint256 constant SWAP_TOKENS_FOR_EXACT_AVAX = 73;
uint256 constant SWAP_EXACT_TOKENS_FOR_AVAX = 74;
uint256 constant SWAP_AVAX_FOR_EXACT_TOKENS = 75;
```

## Disclaimer

GLHF!
