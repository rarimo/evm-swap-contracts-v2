import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  SwapDiamond,
  MasterRouter,
  TransferRouter,
  ERC20MintableBurnable,
  WrappedNativeMock,
  UniswapV3RouterMock,
  UniswapV3Router,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CALLER_ADDRESS, SelectorType, THIS_ADDRESS } from "./utils/contants";
import { wei, weiBTC, weiUSDT } from "../scripts/utils/utils";
import { setBalance } from "@nomicfoundation/hardhat-network-helpers";

describe("UniswapV3Router", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let RECEIVER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let transfer: TransferRouter;
  let uniswapV3: UniswapV3Router;
  let uniswapV3Proxy: UniswapV3Router;
  let tokens: {
    WETH: WrappedNativeMock;
    USDT: ERC20MintableBurnable;
    BTC: ERC20MintableBurnable;
  };
  let uniswapV3Mock: UniswapV3RouterMock;

  function v3Path(...tokens: string[]) {
    let types = new Array<string>();
    let values = new Array<string | number>();

    for (let i = 0; i < tokens.length; ++i) {
      types.push("address");
      values.push(tokens[i]);

      if (i !== tokens.length - 1) {
        types.push("uint24");
        values.push(3000);
      }
    }

    return ethers.utils.solidityPack(types, values);
  }

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const TransferRouter = await ethers.getContractFactory("TransferRouter");
    const UniswapV3Router = await ethers.getContractFactory("UniswapV3Router");
    const UniswapV3RouterMock = await ethers.getContractFactory("UniswapV3RouterMock");
    const WrappedNativeMock = await ethers.getContractFactory("WrappedNativeMock");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    transfer = await TransferRouter.deploy();
    uniswapV3 = await UniswapV3Router.deploy();
    uniswapV3Proxy = await UniswapV3Router.attach(diamond.address);
    tokens = {
      WETH: await WrappedNativeMock.deploy(),
      USDT: await ERC20MintableBurnable.deploy("USDT", "USDT", 6, OWNER.address),
      BTC: await ERC20MintableBurnable.deploy("BTC", "BTC", 8, OWNER.address),
    };
    uniswapV3Mock = await UniswapV3RouterMock.deploy(tokens.WETH.address);

    await setBalance(CALLER.address, wei("100000000"));

    await tokens.WETH.connect(CALLER).deposit({ value: wei(34000 * 2) });
    await tokens.BTC.mintTo(CALLER.address, weiBTC(1700 * 2));
    await tokens.USDT.mintTo(CALLER.address, weiUSDT(34000 * 1700 * 2));

    await tokens.WETH.connect(CALLER).approve(diamond.address, wei("34000"));
    await tokens.BTC.connect(CALLER).approve(diamond.address, weiBTC("1700"));
    await tokens.USDT.connect(CALLER).approve(diamond.address, weiUSDT(34000 * 1700));

    await tokens.WETH.connect(CALLER).approve(uniswapV3Mock.address, wei("34000"));
    await tokens.BTC.connect(CALLER).approve(uniswapV3Mock.address, weiBTC("1700"));
    await tokens.USDT.connect(CALLER).approve(uniswapV3Mock.address, weiUSDT(34000 * 1700));

    await uniswapV3Mock.connect(CALLER).setReserve(tokens.WETH.address, wei("34000"));
    await uniswapV3Mock.connect(CALLER).setReserve(tokens.BTC.address, weiBTC("1700"));
    await uniswapV3Mock.connect(CALLER).setReserve(tokens.USDT.address, weiUSDT(34000 * 1700));

    await uniswapV3Mock.enablePair(tokens.WETH.address, tokens.USDT.address);
    await uniswapV3Mock.enablePair(tokens.BTC.address, tokens.USDT.address);

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
      uniswapV3.address,
      [builder("setUniswapV3RouterAddress").selector, builder("exactInput").selector, builder("exactOutput").selector],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await uniswapV3Proxy.setUniswapV3RouterAddress(uniswapV3Mock.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#setUniswapV3RouterAddress", () => {
    it("should not set if caller is not the owner", async () => {
      await expect(uniswapV3Proxy.connect(CALLER).setUniswapV3RouterAddress(uniswapV3Mock.address)).to.be.revertedWith(
        "ODStorage: not an owner"
      );
    });
  });

  describe("#exactInput", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.WETH.address, wei("20")]).payload(),
          builder("exactInput", [false, RECEIVER.address, wei("20"), weiBTC("1"), v3Path()]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: BytesHelper: invalid path length");
    });

    it("should swap exact tokens for tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.WETH.address, wei("20")]).payload(),
          builder("exactInput", [
            false,
            RECEIVER.address,
            wei("20"),
            weiBTC("1"),
            v3Path(tokens.WETH.address, tokens.USDT.address, tokens.BTC.address),
          ]).payload(),
        ]);

      await expect(tx).to.changeTokenBalance(tokens.WETH, CALLER, -wei("20"));
      await expect(tx).to.changeTokenBalance(tokens.BTC, RECEIVER, weiBTC("1"));
    });

    it("should swap exact ETH for tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("exactInput", [
              true,
              CALLER_ADDRESS,
              wei("10"),
              weiBTC("0.5"),
              v3Path(tokens.WETH.address, tokens.USDT.address, tokens.BTC.address),
            ]).payload(),
          ],
          { value: wei("10") }
        );

      await expect(tx).to.changeEtherBalance(CALLER, -wei("10"));
      await expect(tx).to.changeTokenBalance(tokens.BTC, CALLER, weiBTC("0.5"));
    });

    it("should swap exact ETH for tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("exactInput", [
              true,
              THIS_ADDRESS,
              wei("10"),
              weiBTC("0.1"),
              v3Path(tokens.WETH.address, tokens.USDT.address, tokens.BTC.address),
            ]).payload(),
          ],
          { value: wei("10") }
        );

      await expect(tx).to.changeEtherBalance(CALLER, -wei("10"));
      await expect(tx).to.changeTokenBalance(tokens.BTC, diamond, weiBTC("0.5"));
    });
  });

  describe("#exactOutput", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("2")]).payload(),
          builder("exactOutput", [false, RECEIVER.address, wei("40"), weiBTC("2"), v3Path()]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: BytesHelper: invalid path length");
    });

    it("should swap tokens for exact tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("2")]).payload(),
          builder("exactOutput", [
            false,
            RECEIVER.address,
            wei("40"),
            weiBTC("2"),
            v3Path(tokens.WETH.address, tokens.USDT.address, tokens.BTC.address),
          ]).payload(),
        ]);

      await expect(tx).to.changeTokenBalance(tokens.BTC, CALLER, -weiBTC("2"));
      await expect(tx).to.changeTokenBalance(tokens.WETH, RECEIVER, wei("40"));
    });

    it("should swap tokens for exact tokens with token refund (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("3")]).payload(),
          builder("exactOutput", [
            false,
            RECEIVER.address,
            wei("40"),
            weiBTC("3"),
            v3Path(tokens.WETH.address, tokens.USDT.address, tokens.BTC.address),
          ]).payload(),
        ]);

      await expect(tx).to.changeTokenBalance(tokens.BTC, CALLER, -weiBTC("2"));
      await expect(tx).to.changeTokenBalance(tokens.WETH, RECEIVER, wei("40"));
    });

    it("should swap eth for exact tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("exactOutput", [
              true,
              CALLER_ADDRESS,
              weiBTC("2"),
              wei("40"),
              v3Path(tokens.BTC.address, tokens.USDT.address, tokens.WETH.address),
            ]).payload(),
          ],
          { value: wei("40") }
        );

      await expect(tx).to.changeEtherBalance(CALLER, -wei("40"));
      await expect(tx).to.changeTokenBalance(tokens.BTC, CALLER, weiBTC("2"));
    });

    it("should swap eth for exact tokens with eth refund (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("exactOutput", [
              true,
              THIS_ADDRESS,
              weiBTC("2"),
              wei("60"),
              v3Path(tokens.BTC.address, tokens.USDT.address, tokens.WETH.address),
            ]).payload(),
          ],
          { value: wei("60") }
        );

      await expect(tx).to.changeEtherBalance(CALLER, -wei("40"));
      await expect(tx).to.changeTokenBalance(tokens.BTC, diamond, weiBTC("2"));
    });
  });
});
