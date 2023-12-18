/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  OwnableDiamondStorage,
  OwnableDiamondStorageInterface,
} from "../../../../../../@solarity/solidity-lib/diamond/presets/OwnableDiamond/OwnableDiamondStorage";

const _abi = [
  {
    inputs: [],
    name: "OWNABLE_DIAMOND_STORAGE_SLOT",
    outputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
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
] as const;

export class OwnableDiamondStorage__factory {
  static readonly abi = _abi;
  static createInterface(): OwnableDiamondStorageInterface {
    return new utils.Interface(_abi) as OwnableDiamondStorageInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): OwnableDiamondStorage {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as OwnableDiamondStorage;
  }
}
