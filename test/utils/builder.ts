import { ethers } from "hardhat";
import { Commands } from "./contants";
import {
  BridgeRouter,
  MasterRouter,
  MulticallRouter,
  TraderJoeRouter,
  TransferRouter,
  UniswapV2Router,
  UniswapV3Router,
  WrapRouter,
} from "../../generated-types/ethers";
import { ContractFactory } from "ethers";

type Selector = keyof (MasterRouter["functions"] &
  BridgeRouter["functions"] &
  MulticallRouter["functions"] &
  TraderJoeRouter["functions"] &
  TransferRouter["functions"] &
  UniswapV2Router["functions"] &
  UniswapV3Router["functions"] &
  WrapRouter["functions"]);

function getCommand(selector: Selector): Commands {
  switch (selector) {
    case "bridgeERC20":
      return Commands.BRIDGE_ERC20;
    case "bridgeERC721":
      return Commands.BRIDGE_ERC721;
    case "bridgeERC1155":
      return Commands.BRIDGE_ERC1155;
    case "bridgeNative":
      return Commands.BRIDGE_NATIVE;
    case "transferERC20":
      return Commands.TRANSFER_ERC20;
    case "transferERC721":
      return Commands.TRANSFER_ERC721;
    case "transferERC1155":
      return Commands.TRANSFER_ERC1155;
    case "transferNative":
      return Commands.TRANSFER_NATIVE;
    case "wrap":
      return Commands.WRAP_NATIVE;
    case "unwrap":
      return Commands.UNWRAP_NATIVE;
    case "multicall":
      return Commands.MULTICALL;
    case "swapExactTokensForTokensV2":
      return Commands.SWAP_EXACT_TOKENS_FOR_TOKENS_V2;
    case "swapTokensForExactTokensV2":
      return Commands.SWAP_TOKENS_FOR_EXACT_TOKENS_V2;
    case "swapExactETHForTokens":
      return Commands.SWAP_EXACT_ETH_FOR_TOKENS;
    case "swapTokensForExactETH":
      return Commands.SWAP_TOKENS_FOR_EXACT_ETH;
    case "swapExactTokensForETH":
      return Commands.SWAP_EXACT_TOKENS_FOR_ETH;
    case "swapETHForExactTokens":
      return Commands.SWAP_ETH_FOR_EXACT_TOKENS;
    case "exactInput":
      return Commands.EXACT_INPUT;
    case "exactOutput":
      return Commands.EXACT_OUTPUT;
    case "swapExactTokensForTokensTJ":
      return Commands.SWAP_EXACT_TOKENS_FOR_TOKENS_TJ;
    case "swapTokensForExactTokensTJ":
      return Commands.SWAP_TOKENS_FOR_EXACT_TOKENS_TJ;
    case "swapExactAVAXForTokens":
      return Commands.SWAP_EXACT_AVAX_FOR_TOKENS;
    case "swapTokensForExactAVAX":
      return Commands.SWAP_TOKENS_FOR_EXACT_AVAX;
    case "swapExactTokensForAVAX":
      return Commands.SWAP_EXACT_TOKENS_FOR_AVAX;
    case "swapAVAXForExactTokens":
      return Commands.SWAP_AVAX_FOR_EXACT_TOKENS;
    default:
      return Commands.UNDEFINED;
  }
}

export async function getBuilder() {
  const factories: ContractFactory[] = await Promise.all([
    ethers.getContractFactory("MasterRouter"),
    ethers.getContractFactory("BridgeRouter"),
    ethers.getContractFactory("MulticallRouter"),
    ethers.getContractFactory("TraderJoeRouter"),
    ethers.getContractFactory("TransferRouter"),
    ethers.getContractFactory("UniswapV2Router"),
    ethers.getContractFactory("UniswapV3Router"),
    ethers.getContractFactory("WrapRouter"),
  ]);

  return function (selector: Selector, values: any[] = []) {
    selector = selector.split("(")[0] as Selector;

    const abstractFactory = factories.find((i: ContractFactory) => {
      return Object.keys(i.interface.functions).some((f) => {
        return selector === f.split("(")[0];
      });
    });

    if (abstractFactory === undefined) {
      throw Error("unexpected selector");
    }

    const abstractInterface = abstractFactory.interface;

    let functionData = "";

    const fragment = abstractInterface.getFunction(selector);

    if (fragment?.inputs.length === values.length) {
      functionData = abstractInterface.encodeFunctionData(fragment, values);
    }

    return {
      selector: abstractInterface?.getSighash(selector),
      functionData: functionData,
      payload: (callerPayer: boolean = true, skipRevert: boolean = false) => ({
        command: getCommand(selector),
        skipRevert: skipRevert,
        callerPayer: callerPayer,
        data: "0x" + functionData.slice(10),
      }),
    };
  };
}
