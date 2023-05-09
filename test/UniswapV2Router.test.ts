import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  SwapDiamond,
  MasterRouter,
  TransferRouter,
  UniswapV2Router,
  UniswapV2RouterMock,
  ERC20MintableBurnable,
  WrappedNativeMock,
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
    WETH: WrappedNativeMock;
    USDT: ERC20MintableBurnable;
    BTC: ERC20MintableBurnable;
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
    const WrappedNativeMock = await ethers.getContractFactory("WrappedNativeMock");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    transfer = await TransferRouter.deploy();
    uniswapV2 = await UniswapV2Router.deploy();
    uniswapV2Proxy = await UniswapV2Router.attach(diamond.address);
    tokens = {
      WETH: await WrappedNativeMock.deploy(),
      USDT: await ERC20MintableBurnable.deploy("USDT", "USDT", 6, OWNER.address),
      BTC: await ERC20MintableBurnable.deploy("BTC", "BTC", 8, OWNER.address),
    };
    uniswapV2Mock = await UniswapV2RouterMock.deploy(tokens.WETH.address);

    await setBalance(CALLER.address, wei("100000000"));

    await tokens.WETH.connect(CALLER).deposit({ value: wei(34000 * 2) });
    await tokens.BTC.mintTo(CALLER.address, weiBTC(1700 * 2));
    await tokens.USDT.mintTo(CALLER.address, weiUSDT(34000 * 1700 * 2));

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
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapExactTokensForTokensV2", [RECEIVER.address, weiUSDT("34000"), weiBTC("1"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: UniswapV2Router: invalid path");
    });

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

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("34000"));
      await expect(tx).changeTokenBalance(tokens.BTC, RECEIVER, weiBTC("1"));
    });

    it("should swap exact tokens for tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.WETH.address, wei("20")]).payload(),
          builder("swapExactTokensForTokensV2", [
            CALLER_ADDRESS,
            wei("20"),
            weiBTC("1"),
            [tokens.WETH.address, tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.WETH, CALLER, -wei("20"));
      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, weiBTC("1"));
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

      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, -weiBTC("1"));
      await expect(tx).changeTokenBalance(tokens.WETH, diamond, wei("20"));
    });
  });

  describe("#swapTokensForExactTokensV2", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapTokensForExactTokensV2", [RECEIVER.address, weiBTC("1"), weiUSDT("34000"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: UniswapV2Router: invalid path");
    });

    it("should swap tokens for exact tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapTokensForExactTokensV2", [
            RECEIVER.address,
            weiBTC("1"),
            weiUSDT("34000"),
            [tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("34000"));
      await expect(tx).changeTokenBalance(tokens.BTC, RECEIVER, weiBTC("1"));
    });

    it("should swap tokens for exact tokens with refund (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("40000")]).payload(),
          builder("swapTokensForExactTokensV2", [
            RECEIVER.address,
            weiBTC("1"),
            weiUSDT("40000"),
            [tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("34000"));
      await expect(tx).changeTokenBalance(tokens.BTC, RECEIVER, weiBTC("1"));
    });

    it("should swap tokens for exact tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.WETH.address, wei("20")]).payload(),
          builder("swapTokensForExactTokensV2", [
            CALLER_ADDRESS,
            weiBTC("1"),
            wei("20"),
            [tokens.WETH.address, tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.WETH, CALLER, -wei("20"));
      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, weiBTC("1"));
    });

    it("should swap tokens for exact tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("1")]).payload(),
          builder("swapTokensForExactTokensV2", [
            THIS_ADDRESS,
            wei("20"),
            weiBTC("1"),
            [tokens.BTC.address, tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, -weiBTC("1"));
      await expect(tx).changeTokenBalance(tokens.WETH, diamond, wei("20"));
    });
  });

  describe("#swapExactETHForTokens", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([builder("swapExactETHForTokens", [RECEIVER.address, wei("1"), weiUSDT("1700"), []]).payload()], {
          value: wei("1"),
        });

      await expect(tx).to.be.revertedWith("MasterRouter: UniswapV2Router: invalid path");
    });

    it("should swap exact ETH for tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapExactETHForTokens", [
              RECEIVER.address,
              wei("1"),
              weiUSDT("1700"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, RECEIVER, weiUSDT("1700"));
    });

    it("should swap exact ETH for tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapExactETHForTokens", [
              CALLER_ADDRESS,
              wei("1"),
              weiUSDT("1700"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, weiUSDT("1700"));
    });

    it("should swap exact ETH for tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapExactETHForTokens", [
              THIS_ADDRESS,
              wei("1"),
              weiUSDT("1700"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, diamond, weiUSDT("1700"));
    });
  });

  describe("#swapTokensForExactETH", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactETH", [RECEIVER.address, wei("1"), weiUSDT("1700"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: UniswapV2Router: invalid path");
    });

    it("should swap tokens for exact ETH (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactETH", [
            RECEIVER.address,
            wei("1"),
            weiUSDT("1700"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(RECEIVER, wei("1"));
    });

    it("should swap tokens for exact ETH with refund (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("2000")]).payload(),
          builder("swapTokensForExactETH", [
            RECEIVER.address,
            wei("1"),
            weiUSDT("2000"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(RECEIVER, wei("1"));
    });

    it("should swap tokens for exact ETH (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactETH", [
            CALLER_ADDRESS,
            wei("1"),
            weiUSDT("1700"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(CALLER, wei("1"));
    });

    it("should swap tokens for exact ETH (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactETH", [
            THIS_ADDRESS,
            wei("1"),
            weiUSDT("1700"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(diamond, wei("1"));
    });
  });

  describe("#swapExactTokensForETH", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForETH", [RECEIVER.address, weiUSDT("1700"), wei("1"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: UniswapV2Router: invalid path");
    });

    it("should swap exact tokens for ETH (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForETH", [
            RECEIVER.address,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(RECEIVER, wei("1"));
    });

    it("should swap exact tokens for ETH (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForETH", [
            CALLER_ADDRESS,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(CALLER, wei("1"));
    });

    it("should swap exact tokens for ETH (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForETH", [
            THIS_ADDRESS,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WETH.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(diamond, wei("1"));
    });
  });

  describe("#swapETHForExactTokens", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([builder("swapETHForExactTokens", [RECEIVER.address, weiUSDT("1700"), wei("1"), []]).payload()], {
          value: wei("1"),
        });

      await expect(tx).to.be.revertedWith("MasterRouter: UniswapV2Router: invalid path");
    });

    it("should swap ETH for exact tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapETHForExactTokens", [
              RECEIVER.address,
              weiUSDT("1700"),
              wei("1"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, RECEIVER, weiUSDT("1700"));
    });

    it("should swap ETH for exact tokens with refund (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapETHForExactTokens", [
              RECEIVER.address,
              weiUSDT("1700"),
              wei("2"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("2") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, RECEIVER, weiUSDT("1700"));
    });

    it("should swap ETH for exact tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapETHForExactTokens", [
              CALLER_ADDRESS,
              weiUSDT("1700"),
              wei("1"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, weiUSDT("1700"));
    });

    it("should swap ETH for exact tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapETHForExactTokens", [
              THIS_ADDRESS,
              weiUSDT("1700"),
              wei("1"),
              [tokens.WETH.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, diamond, weiUSDT("1700"));
    });
  });
});
