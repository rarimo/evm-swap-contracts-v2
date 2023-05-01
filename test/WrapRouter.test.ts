import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import { SwapDiamond, MasterRouter, WETH9Mock, WrapRouter } from "../generated-types/ethers";
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
  let wrap: WrapRouter;
  let wrapProxy: WrapRouter;
  let weth9: WETH9Mock;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const WrapRouter = await ethers.getContractFactory("WrapRouter");
    const WETH9Mock = await ethers.getContractFactory("WETH9Mock");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    wrap = await WrapRouter.deploy();
    wrapProxy = await WrapRouter.attach(diamond.address);
    weth9 = await WETH9Mock.deploy();

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      [builder("make").selector, builder("getCallerAddress").selector],
      [SelectorType.SwapDiamond, SelectorType.SwapDiamond]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      wrap.address,
      [builder("setWETH9Address").selector, builder("wrap").selector, builder("unwrap").selector],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await wrapProxy.setWETH9Address(weth9.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#setWETH9Address", () => {
    it("should not set if caller is not the owner", async () => {
      await expect(wrapProxy.connect(CALLER).setWETH9Address(weth9.address)).to.be.revertedWith(
        "ODStorage: not an owner"
      );
    });
  });

  describe("#wrap", () => {
    it("should wrap and resolve amount & receiver properly", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("wrap", [THIS_ADDRESS, wei("0.2")]).payload(),
            builder("wrap", [RECEIVER.address, wei("0.1")]).payload(),
            builder("wrap", [CALLER_ADDRESS, CONTRACT_BALANCE]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeTokenBalances(weth9, [diamond, RECEIVER, CALLER], [wei("0.2"), wei("0.1"), wei("0.7")]);
    });
  });

  describe("#unwrap", () => {
    it("should wrap and resolve amount & receiver properly", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("wrap", [THIS_ADDRESS, wei("1")]).payload(),
            builder("unwrap", [THIS_ADDRESS, wei("0.2")]).payload(),
            builder("unwrap", [RECEIVER.address, wei("0.1")]).payload(),
            builder("unwrap", [CALLER_ADDRESS, CONTRACT_BALANCE]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).changeEtherBalances([diamond, RECEIVER, CALLER], [wei("0.2"), wei("0.1"), -wei("0.3")]);
    });
  });
});
