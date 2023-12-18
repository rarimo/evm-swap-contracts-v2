/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  ERC20Handler,
  ERC20HandlerInterface,
} from "../../../../@rarimo/evm-bridge-contracts/handlers/ERC20Handler";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "salt",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "bundle",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "string",
        name: "network",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "receiver",
        type: "string",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "isWrapped",
        type: "bool",
      },
    ],
    name: "DepositedERC20",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint8",
        name: "version",
        type: "uint8",
      },
    ],
    name: "Initialized",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "salt",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "bundle",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "originHash",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "proof",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "isWrapped",
        type: "bool",
      },
    ],
    name: "WithdrawnERC20",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "bundleExecutorImplementation_",
        type: "address",
      },
      {
        internalType: "address",
        name: "facade_",
        type: "address",
      },
    ],
    name: "__Bundler_init",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "bundleExecutorImplementation",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "token",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "amount",
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
            name: "bundle",
            type: "tuple",
          },
          {
            internalType: "string",
            name: "network",
            type: "string",
          },
          {
            internalType: "string",
            name: "receiver",
            type: "string",
          },
          {
            internalType: "bool",
            name: "isWrapped",
            type: "bool",
          },
        ],
        internalType: "struct IERC20Handler.DepositERC20Parameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "depositERC20",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "salt_",
        type: "bytes32",
      },
    ],
    name: "determineProxyAddress",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "facade",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "token",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "amount",
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
            name: "bundle",
            type: "tuple",
          },
          {
            internalType: "bytes32",
            name: "originHash",
            type: "bytes32",
          },
          {
            internalType: "address",
            name: "receiver",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "proof",
            type: "bytes",
          },
          {
            internalType: "bool",
            name: "isWrapped",
            type: "bool",
          },
        ],
        internalType: "struct IERC20Handler.WithdrawERC20Parameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawERC20",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "token",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "amount",
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
            name: "bundle",
            type: "tuple",
          },
          {
            internalType: "bytes32",
            name: "originHash",
            type: "bytes32",
          },
          {
            internalType: "address",
            name: "receiver",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "proof",
            type: "bytes",
          },
          {
            internalType: "bool",
            name: "isWrapped",
            type: "bool",
          },
        ],
        internalType: "struct IERC20Handler.WithdrawERC20Parameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawERC20Bundle",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

export class ERC20Handler__factory {
  static readonly abi = _abi;
  static createInterface(): ERC20HandlerInterface {
    return new utils.Interface(_abi) as ERC20HandlerInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ERC20Handler {
    return new Contract(address, _abi, signerOrProvider) as ERC20Handler;
  }
}
