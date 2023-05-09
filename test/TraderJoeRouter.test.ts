import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  SwapDiamond,
  MasterRouter,
  TransferRouter,
  ERC20MintableBurnable,
  WrappedNativeMock,
  TraderJoeRouter,
  TraderJoeMock,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CALLER_ADDRESS, SelectorType, THIS_ADDRESS } from "./utils/contants";
import { wei, weiBTC, weiUSDT } from "../scripts/utils/utils";
import { setBalance } from "@nomicfoundation/hardhat-network-helpers";

describe("TraderJoeRouter", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let RECEIVER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let transfer: TransferRouter;
  let traderJoe: TraderJoeRouter;
  let traderJoeProxy: TraderJoeRouter;
  let tokens: {
    WAVAX: WrappedNativeMock;
    USDT: ERC20MintableBurnable;
    BTC: ERC20MintableBurnable;
  };
  let traderJoeMock: TraderJoeMock;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const TransferRouter = await ethers.getContractFactory("TransferRouter");
    const TraderJoeRouter = await ethers.getContractFactory("TraderJoeRouter");
    const TraderJoeMock = await ethers.getContractFactory("TraderJoeMock");
    const WrappedNativeMock = await ethers.getContractFactory("WrappedNativeMock");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    transfer = await TransferRouter.deploy();
    traderJoe = await TraderJoeRouter.deploy();
    traderJoeProxy = await TraderJoeRouter.attach(diamond.address);
    tokens = {
      WAVAX: await WrappedNativeMock.deploy(),
      USDT: await ERC20MintableBurnable.deploy("USDT", "USDT", 6, OWNER.address),
      BTC: await ERC20MintableBurnable.deploy("BTC", "BTC", 8, OWNER.address),
    };
    traderJoeMock = await TraderJoeMock.deploy(tokens.WAVAX.address);

    await setBalance(CALLER.address, wei("100000000"));

    await tokens.WAVAX.connect(CALLER).deposit({ value: wei(34000 * 2) });
    await tokens.BTC.mintTo(CALLER.address, weiBTC(1700 * 2));
    await tokens.USDT.mintTo(CALLER.address, weiUSDT(34000 * 1700 * 2));

    await tokens.WAVAX.connect(CALLER).approve(diamond.address, wei("34000"));
    await tokens.BTC.connect(CALLER).approve(diamond.address, weiBTC("1700"));
    await tokens.USDT.connect(CALLER).approve(diamond.address, weiUSDT(34000 * 1700));

    await tokens.WAVAX.connect(CALLER).approve(traderJoeMock.address, wei("34000"));
    await tokens.BTC.connect(CALLER).approve(traderJoeMock.address, weiBTC("1700"));
    await tokens.USDT.connect(CALLER).approve(traderJoeMock.address, weiUSDT(34000 * 1700));

    await traderJoeMock.connect(CALLER).setReserve(tokens.WAVAX.address, wei("34000"));
    await traderJoeMock.connect(CALLER).setReserve(tokens.BTC.address, weiBTC("1700"));
    await traderJoeMock.connect(CALLER).setReserve(tokens.USDT.address, weiUSDT(34000 * 1700));

    await traderJoeMock.enablePair(tokens.WAVAX.address, tokens.USDT.address);
    await traderJoeMock.enablePair(tokens.BTC.address, tokens.USDT.address);

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
      traderJoe.address,
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

    await traderJoeProxy.setTraderJoeRouterAddress(traderJoeMock.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#setTraderJoeRouterAddress", () => {
    it("should not set if caller is not the owner", async () => {
      await expect(traderJoeProxy.connect(CALLER).setTraderJoeRouterAddress(traderJoeMock.address)).to.be.revertedWith(
        "ODStorage: not an owner"
      );
    });
  });

  describe("#swapExactTokensForTokensTJ", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapExactTokensForTokensTJ", [RECEIVER.address, weiUSDT("34000"), weiBTC("1"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: TraderJoeRouter: invalid path");
    });

    it("should swap exact tokens for tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapExactTokensForTokensTJ", [
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
          builder("transferFromERC20", [tokens.WAVAX.address, wei("20")]).payload(),
          builder("swapExactTokensForTokensTJ", [
            CALLER_ADDRESS,
            wei("20"),
            weiBTC("1"),
            [tokens.WAVAX.address, tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.WAVAX, CALLER, -wei("20"));
      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, weiBTC("1"));
    });

    it("should swap exact tokens for tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("1")]).payload(),
          builder("swapExactTokensForTokensTJ", [
            THIS_ADDRESS,
            weiBTC("1"),
            wei("20"),
            [tokens.BTC.address, tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, -weiBTC("1"));
      await expect(tx).changeTokenBalance(tokens.WAVAX, diamond, wei("20"));
    });
  });

  describe("#swapTokensForExactTokensTJ", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapTokensForExactTokensTJ", [RECEIVER.address, weiBTC("1"), weiUSDT("34000"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: TraderJoeRouter: invalid path");
    });

    it("should swap tokens for exact tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("34000")]).payload(),
          builder("swapTokensForExactTokensTJ", [
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
          builder("swapTokensForExactTokensTJ", [
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
          builder("transferFromERC20", [tokens.WAVAX.address, wei("20")]).payload(),
          builder("swapTokensForExactTokensTJ", [
            CALLER_ADDRESS,
            weiBTC("1"),
            wei("20"),
            [tokens.WAVAX.address, tokens.USDT.address, tokens.BTC.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.WAVAX, CALLER, -wei("20"));
      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, weiBTC("1"));
    });

    it("should swap tokens for exact tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.BTC.address, weiBTC("1")]).payload(),
          builder("swapTokensForExactTokensTJ", [
            THIS_ADDRESS,
            wei("20"),
            weiBTC("1"),
            [tokens.BTC.address, tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.BTC, CALLER, -weiBTC("1"));
      await expect(tx).changeTokenBalance(tokens.WAVAX, diamond, wei("20"));
    });
  });

  describe("#swapExactAVAXForTokens", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([builder("swapExactAVAXForTokens", [RECEIVER.address, wei("1"), weiUSDT("1700"), []]).payload()], {
          value: wei("1"),
        });

      await expect(tx).to.be.revertedWith("MasterRouter: TraderJoeRouter: invalid path");
    });

    it("should swap exact AVAX for tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapExactAVAXForTokens", [
              RECEIVER.address,
              wei("1"),
              weiUSDT("1700"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, RECEIVER, weiUSDT("1700"));
    });

    it("should swap exact AVAX for tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapExactAVAXForTokens", [
              CALLER_ADDRESS,
              wei("1"),
              weiUSDT("1700"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, weiUSDT("1700"));
    });

    it("should swap exact AVAX for tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapExactAVAXForTokens", [
              THIS_ADDRESS,
              wei("1"),
              weiUSDT("1700"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, diamond, weiUSDT("1700"));
    });
  });

  describe("#swapTokensForExactAVAX", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactAVAX", [RECEIVER.address, wei("1"), weiUSDT("1700"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: TraderJoeRouter: invalid path");
    });

    it("should swap tokens for exact AVAX (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactAVAX", [
            RECEIVER.address,
            wei("1"),
            weiUSDT("1700"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(RECEIVER, wei("1"));
    });

    it("should swap tokens for exact AVAX with refund (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("2000")]).payload(),
          builder("swapTokensForExactAVAX", [
            RECEIVER.address,
            wei("1"),
            weiUSDT("2000"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(RECEIVER, wei("1"));
    });

    it("should swap tokens for exact AVAX (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactAVAX", [
            CALLER_ADDRESS,
            wei("1"),
            weiUSDT("1700"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(CALLER, wei("1"));
    });

    it("should swap tokens for exact AVAX (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapTokensForExactAVAX", [
            THIS_ADDRESS,
            wei("1"),
            weiUSDT("1700"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(diamond, wei("1"));
    });
  });

  describe("#swapExactTokensForAVAX", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForAVAX", [RECEIVER.address, weiUSDT("1700"), wei("1"), []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: TraderJoeRouter: invalid path");
    });

    it("should swap exact tokens for AVAX (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForAVAX", [
            RECEIVER.address,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(RECEIVER, wei("1"));
    });

    it("should swap exact tokens for AVAX (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForAVAX", [
            CALLER_ADDRESS,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(CALLER, wei("1"));
    });

    it("should swap exact tokens for AVAX (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [tokens.USDT.address, weiUSDT("1700")]).payload(),
          builder("swapExactTokensForAVAX", [
            THIS_ADDRESS,
            weiUSDT("1700"),
            wei("1"),
            [tokens.USDT.address, tokens.WAVAX.address],
          ]).payload(),
        ]);

      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, -weiUSDT("1700"));
      await expect(tx).changeEtherBalance(diamond, wei("1"));
    });
  });

  describe("#swapAVAXForExactTokens", () => {
    it("should revert if invalid path", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([builder("swapAVAXForExactTokens", [RECEIVER.address, weiUSDT("1700"), wei("1"), []]).payload()], {
          value: wei("1"),
        });

      await expect(tx).to.be.revertedWith("MasterRouter: TraderJoeRouter: invalid path");
    });

    it("should swap AVAX for exact tokens (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapAVAXForExactTokens", [
              RECEIVER.address,
              weiUSDT("1700"),
              wei("1"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, RECEIVER, weiUSDT("1700"));
    });

    it("should swap AVAX for exact tokens with refund (arbitrary receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapAVAXForExactTokens", [
              RECEIVER.address,
              weiUSDT("1700"),
              wei("2"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("2") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, RECEIVER, weiUSDT("1700"));
    });

    it("should swap AVAX for exact tokens (CALLER_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapAVAXForExactTokens", [
              CALLER_ADDRESS,
              weiUSDT("1700"),
              wei("1"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, CALLER, weiUSDT("1700"));
    });

    it("should swap AVAX for exact tokens (THIS_ADDRESS receiver)", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("swapAVAXForExactTokens", [
              THIS_ADDRESS,
              weiUSDT("1700"),
              wei("1"),
              [tokens.WAVAX.address, tokens.USDT.address],
            ]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalance(CALLER, -wei("1"));
      await expect(tx).changeTokenBalance(tokens.USDT, diamond, weiUSDT("1700"));
    });
  });
});
