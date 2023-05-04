import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  SwapDiamond,
  MasterRouter,
  TransferRouter,
  WETH9Mock,
  UniswapV2Router,
  UniswapV2RouterMock,
  ERC20Mock,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CALLER_ADDRESS, SelectorType, THIS_ADDRESS } from "./utils/contants";
import { wei, weiBTC, weiUSDT } from "../scripts/utils/utils";
import { setBalance } from "@nomicfoundation/hardhat-network-helpers";

describe("UniswapV2Router", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let RECEIVER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let transfer: TransferRouter;
  let uniswapV2: UniswapV2Router;
  let uniswapV2Proxy: UniswapV2Router;
  let tokens: {
    WETH: WETH9Mock;
    USDT: ERC20Mock;
    BTC: ERC20Mock;
  };
  let uniswapV2Mock: UniswapV2RouterMock;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const TransferRouter = await ethers.getContractFactory("TransferRouter");
    const UniswapV2Router = await ethers.getContractFactory("UniswapV2Router");
    const UniswapV2RouterMock = await ethers.getContractFactory("UniswapV2RouterMock");
    const WETH9Mock = await ethers.getContractFactory("WETH9Mock");
    const ERC20Mock = await ethers.getContractFactory("ERC20Mock");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    transfer = await TransferRouter.deploy();
    uniswapV2 = await UniswapV2Router.deploy();
    uniswapV2Proxy = await UniswapV2Router.attach(diamond.address);
    tokens = {
      WETH: await WETH9Mock.deploy(),
      USDT: await ERC20Mock.deploy("USDT", "USDT", 6),
      BTC: await ERC20Mock.deploy("BTC", "BTC", 8),
    };
    uniswapV2Mock = await UniswapV2RouterMock.deploy(tokens.WETH.address);

    await setBalance(CALLER.address, wei("100000000"));

    await tokens.WETH.connect(CALLER).deposit({ value: wei(34000 * 2) });
    await tokens.BTC.mint(CALLER.address, weiBTC(1700 * 2));
    await tokens.USDT.mint(CALLER.address, weiUSDT(34000 * 1700 * 2));

    await tokens.WETH.connect(CALLER).approve(diamond.address, wei("34000"));
    await tokens.BTC.connect(CALLER).approve(diamond.address, weiBTC("1700"));
    await tokens.USDT.connect(CALLER).approve(diamond.address, weiUSDT(34000 * 1700));

    await tokens.WETH.connect(CALLER).approve(uniswapV2Mock.address, wei("34000"));
    await tokens.BTC.connect(CALLER).approve(uniswapV2Mock.address, weiBTC("1700"));
    await tokens.USDT.connect(CALLER).approve(uniswapV2Mock.address, weiUSDT(34000 * 1700));

    await uniswapV2Mock.connect(CALLER).setReserve(tokens.WETH.address, wei("34000"));
    await uniswapV2Mock.connect(CALLER).setReserve(tokens.BTC.address, weiBTC("1700"));
    await uniswapV2Mock.connect(CALLER).setReserve(tokens.USDT.address, weiUSDT(34000 * 1700));

    await uniswapV2Mock.enablePair(tokens.WETH.address, tokens.USDT.address);
    await uniswapV2Mock.enablePair(tokens.BTC.address, tokens.USDT.address);

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      [builder("make").selector, builder("getCallerAddress").selector],
      [SelectorType.SwapDiamond, SelectorType.SwapDiamond]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      transfer.address,
      [builder("transferFromERC20").selector],
      [SelectorType.MasterRouter]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      uniswapV2.address,
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

    await uniswapV2Proxy.setUniswapV2RouterAddress(uniswapV2Mock.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#setUniswapV2RouterAddress", () => {
    it("should not set if caller is not the owner", async () => {
      await expect(uniswapV2Proxy.connect(CALLER).setUniswapV2RouterAddress(uniswapV2Mock.address)).to.be.revertedWith(
        "ODStorage: not an owner"
      );
    });
  });

  describe("#swapExactTokensForTokensV2", () => {
    it("should swap exact tokens for tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapExactTokensForTokensV2", [
            RECEIVER.address,
            weiUSDT("34000"),
            weiBTC("1"),
            [tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.BTC, RECEIVER, weiBTC("1"));
    });

    it("should swap exact tokens for tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForTokensV2", [
            CALLER_ADDRESS,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.WETH, CALLER, wei("1"));
    });

    it("should swap exact tokens for tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("1")]).payload(),
          builder("swapExactTokensForTokensV2", [
            THIS_ADDRESS,
            weiBTC("1"),
            wei("20"),
            [tokens.BTC.address, tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.WETH, diamond, wei("20"));
    });
  });

  describe("#swapTokensForExactTokensV2", () => {
    it("", async () => {});
  });

  describe("#swapExactETHForTokens", () => {
    it("", async () => {});
  });

  describe("#swapTokensForExactETH", () => {
    it("", async () => {});
  });

  describe("#swapExactTokensForETH", () => {
    it("", async () => {});
  });

  describe("#swapETHForExactTokens", () => {
    it("", async () => {});
  });
});
