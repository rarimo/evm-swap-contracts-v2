import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  MasterRouter__factory,
  SwapDiamond__factory,
  SwapDiamond,
  MasterRouter,
  WETH9Mock__factory,
  WETH9Mock,
  WrapRouter__factory,
  WrapRouter,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CALLER_ADDRESS, CONTRACT_BALANCE, SelectorType, THIS_ADDRESS } from "./utils/contants";
import { wei } from "../scripts/utils/utils";

describe("SwapDiamond", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let SwapDiamondFactory: SwapDiamond__factory;
  let MasterRouterFactory: MasterRouter__factory;
  let WrapRouterFactory: WrapRouter__factory;
  let WETH9MockFactory: WETH9Mock__factory;

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterDiamond: MasterRouter;
  let wrap: WrapRouter;
  let wrapDiamond: WrapRouter;
  let weth9: WETH9Mock;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, SECOND] = await ethers.getSigners();

    SwapDiamondFactory = await ethers.getContractFactory("SwapDiamond");
    MasterRouterFactory = await ethers.getContractFactory("MasterRouter");
    WrapRouterFactory = await ethers.getContractFactory("WrapRouter");
    WETH9MockFactory = await ethers.getContractFactory("WETH9Mock");

    diamond = await SwapDiamondFactory.deploy();
    master = await MasterRouterFactory.deploy();
    masterDiamond = await MasterRouterFactory.attach(diamond.address);
    wrap = await WrapRouterFactory.deploy();
    wrapDiamond = await WrapRouterFactory.attach(diamond.address);
    weth9 = await WETH9MockFactory.deploy();

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      ["0x00000000", builder("make").selector, builder("getCallerAddress").selector],
      [SelectorType.SwapDiamond, SelectorType.SwapDiamond, SelectorType.SwapDiamond]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      wrap.address,
      [builder("setWETH9Address").selector, builder("wrap").selector, builder("unwrap").selector],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await wrapDiamond.setWETH9Address(weth9.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#setWETH9Address", () => {
    it("should not set if caller is not the owner", async () => {
      await expect(wrapDiamond.connect(SECOND).setWETH9Address(weth9.address)).to.be.revertedWith(
        "ODStorage: not an owner"
      );
    });
  });

  describe("#wrap", () => {
    it("should wrap and resolve amount & receiver properly", async () => {
      const tx = masterDiamond.make(
        [
          builder("wrap", [THIS_ADDRESS, wei("0.2")]).payload(),
          builder("wrap", [CALLER_ADDRESS, wei("0.1")]).payload(),
          builder("wrap", [SECOND.address, CONTRACT_BALANCE]).payload(),
        ],
        { value: wei("1") }
      );

      await expect(tx).changeTokenBalances(weth9, [diamond, OWNER, SECOND], [wei("0.2"), wei("0.1"), wei("0.7")]);
    });
  });

  describe("#unwrap", () => {
    it("should wrap and resolve amount & receiver properly", async () => {
      const tx = masterDiamond.make(
        [
          builder("wrap", [THIS_ADDRESS, wei("1")]).payload(),
          builder("unwrap", [THIS_ADDRESS, wei("0.2")]).payload(),
          builder("unwrap", [CALLER_ADDRESS, wei("0.1")]).payload(),
          builder("unwrap", [SECOND.address, CONTRACT_BALANCE]).payload(),
        ],
        { value: wei("1") }
      );

      await expect(tx).changeEtherBalances([diamond, OWNER, SECOND], [wei("0.2"), -wei("0.9"), wei("0.7")]);
    });
  });
});
