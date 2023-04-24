const { Commands } = require("./contants");

const parse = (abi, params, command) => {
  const selector = web3.eth.abi.encodeFunctionSignature(abi);

  let functionCallData = "";
  let argumentsData = "";

  if (params.length === 0 || params[0] !== undefined) {
    functionCallData = web3.eth.abi.encodeFunctionCall(abi, params);
    argumentsData = "0x" + functionCallData.slice(10);
  }

  return {
    asSelector: () => selector,
    asArgumentsData: () => argumentsData,
    asFunctionCall: () => functionCallData,
    asPayload: (skipRevert, callerPayer) => ({
      command: command,
      skipRevert: skipRevert,
      callerPayer: callerPayer,
      data: argumentsData,
    }),
  };
};

const getBridgeERC20Data = (callerPayer, token, amount, bundle, network, receiver, isWrapped) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "token_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
        {
          components: [
            {
              internalType: "bytes32",
              name: "salt",
              type: "bytes32",
            },
            {
              internalType: "bytes",
              name: "bundle",
              type: "bytes",
            },
          ],
          internalType: "struct IBundler.Bundle",
          name: "bundle_",
          type: "tuple",
        },
        {
          internalType: "string",
          name: "network_",
          type: "string",
        },
        {
          internalType: "string",
          name: "receiver_",
          type: "string",
        },
        {
          internalType: "bool",
          name: "isWrapped_",
          type: "bool",
        },
      ],
      name: "bridgeERC20",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, token, amount, bundle, network, receiver, isWrapped],
    Commands.BRIDGE_ERC20
  );
};

const getBridgeERC721Data = (callerPayer, token, tokenId, bundle, network, receiver, isWrapped) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "token_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "tokenId_",
          type: "uint256",
        },
        {
          components: [
            {
              internalType: "bytes32",
              name: "salt",
              type: "bytes32",
            },
            {
              internalType: "bytes",
              name: "bundle",
              type: "bytes",
            },
          ],
          internalType: "struct IBundler.Bundle",
          name: "bundle_",
          type: "tuple",
        },
        {
          internalType: "string",
          name: "network_",
          type: "string",
        },
        {
          internalType: "string",
          name: "receiver_",
          type: "string",
        },
        {
          internalType: "bool",
          name: "isWrapped_",
          type: "bool",
        },
      ],
      name: "bridgeERC721",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, token, tokenId, bundle, network, receiver, isWrapped],
    Commands.BRIDGE_ERC721
  );
};

const getBridgeERC1155Data = (callerPayer, token, tokenId, amount, bundle, network, receiver, isWrapped) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "token_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "tokenId_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
        {
          components: [
            {
              internalType: "bytes32",
              name: "salt",
              type: "bytes32",
            },
            {
              internalType: "bytes",
              name: "bundle",
              type: "bytes",
            },
          ],
          internalType: "struct IBundler.Bundle",
          name: "bundle_",
          type: "tuple",
        },
        {
          internalType: "string",
          name: "network_",
          type: "string",
        },
        {
          internalType: "string",
          name: "receiver_",
          type: "string",
        },
        {
          internalType: "bool",
          name: "isWrapped_",
          type: "bool",
        },
      ],
      name: "bridgeERC1155",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, token, tokenId, amount, bundle, network, receiver, isWrapped],
    Commands.BRIDGE_ERC1155
  );
};

const getBridgeNativeData = (amount, bundle, network, receiver) => {
  return parse(
    {
      inputs: [
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
        {
          components: [
            {
              internalType: "bytes32",
              name: "salt",
              type: "bytes32",
            },
            {
              internalType: "bytes",
              name: "bundle",
              type: "bytes",
            },
          ],
          internalType: "struct IBundler.Bundle",
          name: "bundle_",
          type: "tuple",
        },
        {
          internalType: "string",
          name: "network_",
          type: "string",
        },
        {
          internalType: "string",
          name: "receiver_",
          type: "string",
        },
      ],
      name: "bridgeNative",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [amount, bundle, network, receiver],
    Commands.BRIDGE_NATIVE
  );
};

const getTransferERC20Data = (token, receiver, amount) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "token_",
          type: "address",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
      ],
      name: "transferERC20",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [token, receiver, amount],
    Commands.TRANSFER_ERC20
  );
};

const getTransferERC721Data = (token, receiver, nftIds) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "token_",
          type: "address",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256[]",
          name: "nftIds_",
          type: "uint256[]",
        },
      ],
      name: "transferERC721",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [token, receiver, nftIds],
    Commands.TRANSFER_ERC721
  );
};

const getTransferERC1155Data = (token, receiver, tokenIds, amounts) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "token_",
          type: "address",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256[]",
          name: "tokenIds_",
          type: "uint256[]",
        },
        {
          internalType: "uint256[]",
          name: "amounts_",
          type: "uint256[]",
        },
      ],
      name: "transferERC1155",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [token, receiver, tokenIds, amounts],
    Commands.TRANSFER_ERC1155
  );
};

const getTransferNativeData = (receiver, amount) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
      ],
      name: "transferNative",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amount],
    Commands.TRANSFER_NATIVE
  );
};

const getWrapNativeData = (receiver, amount) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
      ],
      name: "wrap",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amount],
    Commands.WRAP_NATIVE
  );
};

const getUnwrapNativeData = (receiver, amount) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amount_",
          type: "uint256",
        },
      ],
      name: "unwrap",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amount],
    Commands.UNWRAP_NATIVE
  );
};

const getMulticallData = (targets, data, values) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address[]",
          name: "targets_",
          type: "address[]",
        },
        {
          internalType: "bytes[]",
          name: "data_",
          type: "bytes[]",
        },
        {
          internalType: "uint256[]",
          name: "values_",
          type: "uint256[]",
        },
      ],
      name: "multicall",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [targets, data, values],
    Commands.MULTICALL
  );
};

const getUniswapV2SwapExactTokensForTokensData = (callerPayer, receiver, amountIn, amountOutMin, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMin_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapExactTokensForTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountIn, amountOutMin, path],
    Commands.UV2_SWAP_EXACT_TOKENS_FOR_TOKENS
  );
};

const getUniswapV2SwapTokensForExactTokensData = (callerPayer, receiver, amountOut, amountInMax, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMax_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapTokensForExactTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountOut, amountInMax, path],
    Commands.UV2_SWAP_TOKENS_FOR_EXACT_TOKENS
  );
};

const getUniswapV2SwapExactETHForTokensData = (receiver, amountIn, amountOutMin, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMin_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapExactETHForTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amountIn, amountOutMin, path],
    Commands.UV2_SWAP_EXACT_ETH_FOR_TOKENS
  );
};

const getUniswapV2SwapTokensForExactETHData = (callerPayer, receiver, amountOut, amountInMax, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMax_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapTokensForExactETH",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountOut, amountInMax, path],
    Commands.UV2_SWAP_TOKENS_FOR_EXACT_ETH
  );
};

const getUniswapV2SwapExactTokensForETHData = (callerPayer, receiver, amountIn, amountOutMin, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMin_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapExactTokensForETH",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountIn, amountOutMin, path],
    Commands.UV2_SWAP_EXACT_TOKENS_FOR_ETH
  );
};

const getUniswapV2SwapETHForExactTokensData = (receiver, amountOut, amountInMax, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMax_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapETHForExactTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amountOut, amountInMax, path],
    Commands.UV2_SWAP_ETH_FOR_EXACT_TOKENS
  );
};

const getExactInputData = (callerPayer, isNative, receiver, amountIn, amountOutMinimum, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "bool",
          name: "isNative_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMinimum_",
          type: "uint256",
        },
        {
          internalType: "bytes",
          name: "path_",
          type: "bytes",
        },
      ],
      name: "exactInput",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, isNative, receiver, amountIn, amountOutMinimum, path],
    Commands.EXACT_INPUT
  );
};

const getExactOutputData = (callerPayer, isNative, receiver, amountOut, amountInMaximum, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "bool",
          name: "isNative_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMaximum_",
          type: "uint256",
        },
        {
          internalType: "bytes",
          name: "path_",
          type: "bytes",
        },
      ],
      name: "exactOutput",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, isNative, receiver, amountOut, amountInMaximum, path],
    Commands.EXACT_OUTPUT
  );
};

const getTradeJoeSwapExactTokensForTokensData = (callerPayer, receiver, amountIn, amountOutMin, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMin_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapExactTokensForTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountIn, amountOutMin, path],
    Commands.TJ_SWAP_EXACT_TOKENS_FOR_TOKENS
  );
};

const getTradeJoeSwapTokensForExactTokensData = (callerPayer, receiver, amountOut, amountInMax, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMax_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapTokensForExactTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountOut, amountInMax, path],
    Commands.TJ_SWAP_TOKENS_FOR_EXACT_TOKENS
  );
};

const getTradeJoeSwapExactAVAXForTokensData = (receiver, amountIn, amountOutMin, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMin_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapExactAVAXForTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amountIn, amountOutMin, path],
    Commands.TJ_SWAP_EXACT_AVAX_FOR_TOKENS
  );
};

const getTradeJoeSwapTokensForExactAVAXData = (callerPayer, receiver, amountOut, amountInMax, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMax_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapTokensForExactAVAX",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountOut, amountInMax, path],
    Commands.TJ_SWAP_TOKENS_FOR_EXACT_AVAX
  );
};

const getTradeJoeSwapExactTokensForAVAXData = (callerPayer, receiver, amountIn, amountOutMin, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bool",
          name: "callerPayer_",
          type: "bool",
        },
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountIn_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountOutMin_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapExactTokensForAVAX",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [callerPayer, receiver, amountIn, amountOutMin, path],
    Commands.TJ_SWAP_EXACT_TOKENS_FOR_AVAX
  );
};

const getTradeJoeSwapAVAXForExactTokensData = (receiver, amountOut, amountInMax, path) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "receiver_",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "amountOut_",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "amountInMax_",
          type: "uint256",
        },
        {
          internalType: "address[]",
          name: "path_",
          type: "address[]",
        },
      ],
      name: "swapAVAXForExactTokens",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [receiver, amountOut, amountInMax, path],
    Commands.TJ_SWAP_AVAX_FOR_EXACT_TOKENS
  );
};

const getMakeData = (payloads) => {
  return parse(
    {
      inputs: [
        {
          components: [
            {
              internalType: "uint256",
              name: "command",
              type: "uint256",
            },
            {
              internalType: "bool",
              name: "skipRevert",
              type: "bool",
            },
            {
              internalType: "bool",
              name: "callerPayer",
              type: "bool",
            },
            {
              internalType: "bytes",
              name: "data",
              type: "bytes",
            },
          ],
          internalType: "struct MasterRouter.Payload[]",
          name: "payloads_",
          type: "tuple[]",
        },
      ],
      name: "make",
      outputs: [],
      stateMutability: "payable",
      type: "function",
    },
    [payloads],
    Commands.NONE
  );
};

const getOnERC1155BatchReceivedData = (operator, from, ids, amounts, data) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
        {
          internalType: "address",
          name: "",
          type: "address",
        },
        {
          internalType: "uint256[]",
          name: "",
          type: "uint256[]",
        },
        {
          internalType: "uint256[]",
          name: "",
          type: "uint256[]",
        },
        {
          internalType: "bytes",
          name: "",
          type: "bytes",
        },
      ],
      name: "onERC1155BatchReceived",
      outputs: [
        {
          internalType: "bytes4",
          name: "",
          type: "bytes4",
        },
      ],
      stateMutability: "nonpayable",
      type: "function",
    },
    [operator, from, ids, amounts, data],
    Commands.NONE
  );
};

const getOnERC1155ReceivedData = (operator, from, id, amount, data) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
        {
          internalType: "address",
          name: "",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "",
          type: "uint256",
        },
        {
          internalType: "bytes",
          name: "",
          type: "bytes",
        },
      ],
      name: "onERC1155Received",
      outputs: [
        {
          internalType: "bytes4",
          name: "",
          type: "bytes4",
        },
      ],
      stateMutability: "nonpayable",
      type: "function",
    },
    [operator, from, id, amount, data],
    Commands.NONE
  );
};

const getOnERC721ReceivedData = (from, to, tokenId, data) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
        {
          internalType: "address",
          name: "",
          type: "address",
        },
        {
          internalType: "uint256",
          name: "",
          type: "uint256",
        },
        {
          internalType: "bytes",
          name: "",
          type: "bytes",
        },
      ],
      name: "onERC721Received",
      outputs: [
        {
          internalType: "bytes4",
          name: "",
          type: "bytes4",
        },
      ],
      stateMutability: "nonpayable",
      type: "function",
    },
    [from, to, tokenId, data],
    Commands.NONE
  );
};

const getSupportsInterfaceData = (interfaceId) => {
  return parse(
    {
      inputs: [
        {
          internalType: "bytes4",
          name: "interfaceId",
          type: "bytes4",
        },
      ],
      name: "supportsInterface",
      outputs: [
        {
          internalType: "bool",
          name: "",
          type: "bool",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    [interfaceId],
    Commands.NONE
  );
};

const getSetBridgeAddressData = (bridge) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "bridge_",
          type: "address",
        },
      ],
      name: "setBridgeAddress",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [bridge],
    Commands.NONE
  );
};

const getSetWETH9AddressData = (weth9) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "weth9_",
          type: "address",
        },
      ],
      name: "setWETH9Address",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [weth9],
    Commands.NONE
  );
};

const getSetUniswapV3RouterAddressData = (swapV3Router) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "swapV3Router_",
          type: "address",
        },
      ],
      name: "setUniswapV3RouterAddress",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [swapV3Router],
    Commands.NONE
  );
};

const getSetUniswapV2RouterAddressData = (swapV2Router) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "swapV2Router_",
          type: "address",
        },
      ],
      name: "setUniswapV2RouterAddress",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [swapV2Router],
    Commands.NONE
  );
};

const getSetTraderJoeRouterAddressData = (traderJoeRouter) => {
  return parse(
    {
      inputs: [
        {
          internalType: "address",
          name: "traderJoeRouter_",
          type: "address",
        },
      ],
      name: "setTraderJoeRouterAddress",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    [traderJoeRouter],
    Commands.NONE
  );
};

module.exports = {
  getBridgeERC20Data,
  getBridgeERC721Data,
  getBridgeERC1155Data,
  getBridgeNativeData,
  getTransferERC20Data,
  getTransferERC721Data,
  getTransferERC1155Data,
  getTransferNativeData,
  getWrapNativeData,
  getUnwrapNativeData,
  getMulticallData,
  getUniswapV2SwapExactTokensForTokensData,
  getUniswapV2SwapTokensForExactTokensData,
  getUniswapV2SwapExactETHForTokensData,
  getUniswapV2SwapTokensForExactETHData,
  getUniswapV2SwapExactTokensForETHData,
  getUniswapV2SwapETHForExactTokensData,
  getExactInputData,
  getExactOutputData,
  getTradeJoeSwapExactTokensForTokensData,
  getTradeJoeSwapTokensForExactTokensData,
  getTradeJoeSwapExactAVAXForTokensData,
  getTradeJoeSwapTokensForExactAVAXData,
  getTradeJoeSwapExactTokensForAVAXData,
  getTradeJoeSwapAVAXForExactTokensData,
  getMakeData,
  getOnERC1155BatchReceivedData,
  getOnERC1155ReceivedData,
  getOnERC721ReceivedData,
  getSupportsInterfaceData,
  getSetBridgeAddressData,
  getSetWETH9AddressData,
  getSetUniswapV3RouterAddressData,
  getSetUniswapV2RouterAddressData,
  getSetTraderJoeRouterAddressData,
};
