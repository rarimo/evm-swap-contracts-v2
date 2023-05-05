import { Deployer, Logger } from "@dlsl/hardhat-migrate";
import { artifacts } from "hardhat";
import { DEX_TYPE, DexType } from "./config/env-utils";
import { SelectorType } from "../test/utils/contants";
import { getBuilder } from "../test/utils/builder";

const SwapDiamond = artifacts.require("SwapDiamond");
const MasterRouter = artifacts.require("MasterRouter");
const BridgeRouter = artifacts.require("BridgeRouter");
const MulticallRouter = artifacts.require("MulticallRouter");
const TransferRouter = artifacts.require("TransferRouter");
const WrapRouter = artifacts.require("WrapRouter");
const UniswapV2Router = artifacts.require("UniswapV2Router");
const UniswapV3Router = artifacts.require("UniswapV3Router");
const TraderJoeRouter = artifacts.require("TraderJoeRouter");

export = async (deployer: Deployer, logger: Logger) => {
  const dexType = DEX_TYPE();

  const builder = await getBuilder();

  const diamond = await SwapDiamond.deployed();

  const addFacet = async (facetName: string, facet: string, selectors: Array<string>, types: Array<SelectorType>) => {
    logger.logTransaction(
      await diamond.methods["addFacet(address,bytes4[],uint8[])"](facet, selectors, types),
      `Add ${facetName} facet`
    );
  };

  await addFacet(
    "MasterRouter",
    (
      await MasterRouter.deployed()
    ).address,
    [
      builder("make").selector,
      builder("onERC721Received").selector,
      builder("onERC1155Received").selector,
      builder("onERC1155BatchReceived").selector,
      builder("getCallerAddress").selector,
    ],
    [
      SelectorType.SwapDiamond,
      SelectorType.SwapDiamond,
      SelectorType.SwapDiamond,
      SelectorType.SwapDiamond,
      SelectorType.SwapDiamond,
    ]
  );

  await addFacet(
    "BridgeRouter",
    (
      await BridgeRouter.deployed()
    ).address,
    [
      builder("setBridgeAddress").selector,
      builder("bridgeERC20").selector,
      builder("bridgeERC721").selector,
      builder("bridgeERC1155").selector,
      builder("bridgeNative").selector,
    ],
    [
      SelectorType.SwapDiamond,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
    ]
  );

  await addFacet(
    "MulticallRouter",
    (
      await MulticallRouter.deployed()
    ).address,
    [builder("multicall").selector],
    [SelectorType.MasterRouter]
  );

  await addFacet(
    "TransferRouter",
    (
      await TransferRouter.deployed()
    ).address,
    [
      builder("transferERC20").selector,
      builder("transferERC721").selector,
      builder("transferERC1155").selector,
      builder("transferNative").selector,
      builder("transferFromERC20").selector,
      builder("transferFromERC721").selector,
      builder("transferFromERC1155").selector,
    ],
    [
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
      SelectorType.MasterRouter,
    ]
  );

  await addFacet(
    "WrapRouter",
    (
      await WrapRouter.deployed()
    ).address,
    [builder("setWrappedNativeAddress").selector, builder("wrap").selector, builder("unwrap").selector],
    [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
  );

  if (dexType == DexType.Uniswap) {
    await addFacet(
      "UniswapV2Router",
      (
        await UniswapV2Router.deployed()
      ).address,
      [
        builder("setUniswapV2RouterAddress").selector,
        builder("swapExactTokensForTokensV2").selector,
        builder("swapTokensForExactTokensV2").selector,
        builder("swapExactETHForTokens").selector,
        builder("swapTokensForExactETH").selector,
        builder("swapExactTokensForETH").selector,
        builder("swapETHForExactTokens").selector,
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );

    await addFacet(
      "UniswapV3Router",
      (
        await UniswapV3Router.deployed()
      ).address,
      [builder("setUniswapV3RouterAddress").selector, builder("exactInput").selector, builder("exactOutput").selector],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );
  } else {
    await addFacet(
      "TraderJoeRouter",
      (
        await TraderJoeRouter.deployed()
      ).address,
      [
        builder("setTraderJoeRouterAddress").selector,
        builder("swapExactTokensForTokensTJ").selector,
        builder("swapTokensForExactTokensTJ").selector,
        builder("swapExactAVAXForTokens").selector,
        builder("swapTokensForExactAVAX").selector,
        builder("swapExactTokensForAVAX").selector,
        builder("swapAVAXForExactTokens").selector,
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );
  }
};
