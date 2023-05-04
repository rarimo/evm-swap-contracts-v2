import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import { MasterRouter, MulticallRouter, SwapDiamond, WrappedNativeMock, WrapRouter } from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CALLER_ADDRESS, CONTRACT_BALANCE, SelectorType, THIS_ADDRESS } from "./utils/contants";
import { wei } from "../scripts/utils/utils";

describe("WrapRouter", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let RECEIVER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let multicall: MulticallRouter;
  let wrap: WrapRouter;
  let weth9: WrappedNativeMock;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const MulticallRouter = await ethers.getContractFactory("MulticallRouter");
    const WrapRouter = await ethers.getContractFactory("WrapRouter");
    const WrappedNativeMock = await ethers.getContractFactory("WrappedNativeMock");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    multicall = await MulticallRouter.deploy();
    wrap = await WrapRouter.deploy();
    weth9 = await WrappedNativeMock.deploy();

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      [builder("make").selector, builder("getCallerAddress").selector],
      [SelectorType.SwapDiamond, SelectorType.SwapDiamond]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      multicall.address,
      [builder("multicall").selector],
      [SelectorType.MasterRouter]
    );

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#multicall", () => {
    it("should multicall wrap and transfer", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("multicall", [
            [weth9.address, weth9.address, weth9.address, THIS_ADDRESS, RECEIVER.address, CALLER_ADDRESS],
            [
              weth9.interface.encodeFunctionData("deposit"),
              weth9.interface.encodeFunctionData("transfer", [RECEIVER.address, wei("0.5")]),
              weth9.interface.encodeFunctionData("withdraw", [wei("0.5")]),
              "0x",
              "0x",
              "0x",
            ],
            [CONTRACT_BALANCE, 0, 0, wei("0.5"), wei("0.25"), CONTRACT_BALANCE],
          ]).payload(),
        ],
        {
          value: wei("1"),
        }
      );

      await expect(tx).changeTokenBalance(weth9, RECEIVER, wei("0.5"));
      await expect(tx).changeEtherBalances([RECEIVER, CALLER], [wei("0.25"), -wei("0.75")]);
    });

    it("should not multicall if params lengths mismatch", async () => {
      await expect(
        masterProxy.connect(CALLER).make([builder("multicall", [[RECEIVER.address], [], []]).payload()])
      ).to.be.revertedWith("MasterRouter: MulticallRouter: lengths mismatch");

      await expect(
        masterProxy.connect(CALLER).make([builder("multicall", [[], ["0x"], []]).payload()])
      ).to.be.revertedWith("MasterRouter: MulticallRouter: lengths mismatch");

      await expect(
        masterProxy.connect(CALLER).make([builder("multicall", [[], [], [0]]).payload()])
      ).to.be.revertedWith("MasterRouter: MulticallRouter: lengths mismatch");
    });

    it("should not multicall if cannot call", async () => {
      await expect(
        masterProxy
          .connect(CALLER)
          .make([builder("multicall", [[multicall.address], ["0x"], [wei("1")]]).payload()], { value: wei("1") })
      ).to.be.revertedWith("MasterRouter: MulticallRouter: failed to call");
    });
  });
});
