import { BigNumber } from "ethers";

export const CONTRACT_BALANCE = BigNumber.from(2).pow(255);

export const THIS_ADDRESS = "0x0000000000000000000000000000000000000001";
export const CALLER_ADDRESS = "0x0000000000000000000000000000000000000002";
export const ETHEREUM_ADDRESS = "0x0000000000000000000000000000000000000000";
export const COMMISSION_ADDRESS = "0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF";

export enum SelectorType {
  Undefined,
  SwapDiamond,
  MasterRouter,
}

export enum Commands {
  UNDEFINED = 0,
  BRIDGE_ERC20 = 1,
  BRIDGE_ERC721 = 2,
  BRIDGE_ERC1155 = 3,
  BRIDGE_NATIVE = 4,
  TRANSFER_ERC20 = 10,
  TRANSFER_ERC721 = 11,
  TRANSFER_ERC1155 = 12,
  TRANSFER_NATIVE = 13,
  TRANSFER_FROM_ERC20 = 14,
  TRANSFER_FROM_ERC721 = 15,
  TRANSFER_FROM_ERC1155 = 16,
  WRAP_NATIVE = 20,
  UNWRAP_NATIVE = 21,
  MULTICALL = 25,
  SWAP_EXACT_TOKENS_FOR_TOKENS_V2 = 50,
  SWAP_TOKENS_FOR_EXACT_TOKENS_V2 = 51,
  SWAP_EXACT_ETH_FOR_TOKENS = 52,
  SWAP_TOKENS_FOR_EXACT_ETH = 53,
  SWAP_EXACT_TOKENS_FOR_ETH = 54,
  SWAP_ETH_FOR_EXACT_TOKENS = 55,
  EXACT_INPUT = 60,
  EXACT_OUTPUT = 61,
  SWAP_EXACT_TOKENS_FOR_TOKENS_TJ = 70,
  SWAP_TOKENS_FOR_EXACT_TOKENS_TJ = 71,
  SWAP_EXACT_AVAX_FOR_TOKENS = 72,
  SWAP_TOKENS_FOR_EXACT_AVAX = 73,
  SWAP_EXACT_TOKENS_FOR_AVAX = 74,
  SWAP_AVAX_FOR_EXACT_TOKENS = 75,
}

export enum FacadeMethodId {
  AddFeeToken,
  RemoveFeeToken,
  UpdateFeeToken,
  WithdrawFeeToken,
  AuthorizeUpgrade,
}
