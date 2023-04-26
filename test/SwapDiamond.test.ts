import { ethers } from "hardhat";

import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect } from "chai";
import { wei } from "../scripts/utils/utils";
import { Reverter } from "./helpers/reverter";
import {
  BridgeRouter__factory,
  MasterRouter__factory,
  MasterRouter,
  MulticallRouter__factory,
  SwapDiamond,
  SwapDiamond__factory,
  TraderJoeRouter__factory,
  TransferRouter__factory,
  UniswapV2Router__factory,
  UniswapV3Router__factory,
  WETH9Mock__factory,
  WrapRouter,
  WrapRouter__factory,
} from "../generated-types/ethers";
import { getBuilder } from "./utils/builder";
import { SelectorType, THIS_ADDRESS } from "./utils/contants";

describe("SwapDiamond", async () => {
  const reverter = new Reverter();

  let builder: Awaited<ReturnType<typeof getBuilder>>;

  let SwapDiamond: SwapDiamond__factory;
  let MasterRouter: MasterRouter__factory;
  let BridgeRouter: BridgeRouter__factory;
  let MulticallRouter: MulticallRouter__factory;
  let TraderJoeRouter: TraderJoeRouter__factory;
  let TransferRouter: TransferRouter__factory;
  let UniswapV2Router: UniswapV2Router__factory;
  let UniswapV3Router: UniswapV3Router__factory;
  let WrapRouter: WrapRouter__factory;
  let WETH9Mock: WETH9Mock__factory;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;

  let diamond: SwapDiamond;

  before("setup", async () => {
    builder = await getBuilder();
    [OWNER, CALLER] = await ethers.getSigners();

    SwapDiamond = await ethers.getContractFactory("SwapDiamond");

    MasterRouter = await ethers.getContractFactory("MasterRouter");
    BridgeRouter = await ethers.getContractFactory("BridgeRouter");
    MulticallRouter = await ethers.getContractFactory("MulticallRouter");
    TraderJoeRouter = await ethers.getContractFactory("TraderJoeRouter");
    TransferRouter = await ethers.getContractFactory("TransferRouter");
    UniswapV2Router = await ethers.getContractFactory("UniswapV2Router");
    UniswapV3Router = await ethers.getContractFactory("UniswapV3Router");
    WrapRouter = await ethers.getContractFactory("WrapRouter");

    WETH9Mock = await ethers.getContractFactory("WETH9Mock");

    diamond = await SwapDiamond.deploy();

    const masterRouter = await MasterRouter.deploy();
    const bridgeRouter = await BridgeRouter.deploy();
    const multicallRouter = await MulticallRouter.deploy();
    const traderJoeRouter = await TraderJoeRouter.deploy();
    const transferRouter = await TransferRouter.deploy();
    const uniswapV2Router = await UniswapV2Router.deploy();
    const uniswapV3Router = await UniswapV3Router.deploy();
    const wrapRouter = await WrapRouter.deploy();

    await diamond["addFacet(address,bytes4[],uint8[])"](
      masterRouter.address,
      [
        "0x00000000",
        builder("make").selector,
        builder("onERC721Received").selector,
        builder("onERC1155Received").selector,
        builder("onERC1155BatchReceived").selector,
        builder("supportsInterface").selector,
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
      ]
    );

    await diamond["addFacet(address,bytes4[],uint8[])"](
      bridgeRouter.address,
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

    await diamond["addFacet(address,bytes4[],uint8[])"](
      multicallRouter.address,
      [builder("multicall").selector],
      [SelectorType.MasterRouter]
    );

    await diamond["addFacet(address,bytes4[],uint8[])"](
      traderJoeRouter.address,
      [
        builder("setTraderJoeRouterAddress").selector,
        builder("swapAVAXForExactTokens").selector,
        builder("swapExactAVAXForTokens").selector,
        builder("swapExactTokensForAVAX").selector,
        builder("swapExactTokensForTokensTJ").selector,
        builder("swapTokensForExactAVAX").selector,
        builder("swapTokensForExactTokensTJ").selector,
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

    await diamond["addFacet(address,bytes4[],uint8[])"](
      transferRouter.address,
      [
        builder("transferERC20").selector,
        builder("transferERC721").selector,
        builder("transferERC1155").selector,
        builder("transferNative").selector,
      ],
      [SelectorType.MasterRouter, SelectorType.MasterRouter, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await diamond["addFacet(address,bytes4[],uint8[])"](
      uniswapV2Router.address,
      [
        builder("setUniswapV2RouterAddress").selector,
        builder("swapETHForExactTokens").selector,
        builder("swapExactETHForTokens").selector,
        builder("swapExactTokensForETH").selector,
        builder("swapExactTokensForTokensV2").selector,
        builder("swapTokensForExactETH").selector,
        builder("swapTokensForExactTokensV2").selector,
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

    await diamond["addFacet(address,bytes4[],uint8[])"](
      uniswapV3Router.address,
      [builder("setUniswapV3RouterAddress").selector, builder("exactInput").selector, builder("exactOutput").selector],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await diamond["addFacet(address,bytes4[],uint8[])"](
      wrapRouter.address,
      [builder("setWETH9Address").selector, builder("unwrap").selector, builder("wrap").selector],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await reverter.snapshot();
  });

  beforeEach(async () => {});

  afterEach(reverter.revert);

  describe("dev", function () {
    it("dev", async function () {
      const weth9 = await WETH9Mock.deploy();
      await weth9.mint(OWNER.address, wei("100"));

      await (await WrapRouter.attach(diamond.address)).setWETH9Address(weth9.address);

      const tx = await (
        await MasterRouter.attach(diamond.address)
      ).make(
        [
          builder("wrap", [THIS_ADDRESS, wei("2.5")]).payload(),
          builder("transferERC20", [weth9.address, OWNER.address, wei("1.5")]).payload(),
          builder("transferERC20", [weth9.address, CALLER.address, wei("1")]).payload(),
        ],
        { value: wei("2.5") }
      );

      console.log(await weth9.balanceOf(OWNER.address));
      console.log(await weth9.balanceOf(CALLER.address));

      await expect(tx).to.changeTokenBalances(weth9, [OWNER, CALLER], [wei("1.5"), wei("1")]);
    });
  });
});
