import { expect } from "chai";
import { Reverter } from "@/test/helpers/reverter";
import { Builder, getBuilder } from "@/test/utils/builder";
import { SwapDiamond, MasterRouter, MulticallRouter } from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { SelectorType } from "@/test/utils/contants";

describe("SwapDiamond", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let multicall: MulticallRouter;
  let multicallProxy: MulticallRouter;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const MulticallRouter = await ethers.getContractFactory("MulticallRouter");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    multicall = await MulticallRouter.deploy();
    multicallProxy = await MulticallRouter.attach(diamond.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#addFacet", async () => {
    it("should add SwapDiamond facet properly", async () => {
      await diamond["addFacet(address,bytes4[])"](master.address, [builder("make").selector]);

      expect(await diamond.getFacetSelectors(master.address)).to.deep.eq([builder("make").selector]);
      expect(await diamond.getSelectorType(builder("make").selector)).to.eq(SelectorType.SwapDiamond);
    });

    it("should add arbitrary type facet properly", async () => {
      await diamond["addFacet(address,bytes4[],uint8[])"](
        multicall.address,
        [builder("multicall").selector],
        [SelectorType.MasterRouter]
      );

      expect(await diamond.getFacetSelectors(multicall.address)).to.deep.eq([builder("multicall").selector]);
      expect(await diamond.getSelectorType(builder("multicall").selector)).to.eq(SelectorType.MasterRouter);
    });

    it("should revert if arguments lengths mismatch", async () => {
      const tx = diamond["addFacet(address,bytes4[],uint8[])"](multicall.address, [builder("multicall").selector], []);

      await expect(tx).to.be.revertedWith("SwapDiamond: lengths mismatch");
    });
  });

  context("if selectors are added", () => {
    beforeEach(async () => {
      await diamond["addFacet(address,bytes4[])"](master.address, [builder("make").selector]);
      await diamond["addFacet(address,bytes4[],uint8[])"](
        multicall.address,
        [builder("multicall").selector],
        [SelectorType.MasterRouter]
      );
    });

    describe("#removeFacet", () => {
      it("should remove facet properly", async () => {
        await diamond.removeFacet(master.address, [builder("make").selector]);

        expect(await diamond.getFacetSelectors(master.address)).to.deep.eq([]);
        expect(await diamond.getSelectorType(builder("make").selector)).to.eq(SelectorType.Undefined);
      });
    });

    describe("#updateFacet", () => {
      it("should update selectors to SwapDiamond selectors properly", async () => {
        await diamond["updateFacet(address,bytes4[],bytes4[])"](
          master.address,
          [builder("make").selector],
          [builder("getCallerAddress").selector]
        );

        expect(await diamond.getFacetSelectors(master.address)).to.deep.eq([builder("getCallerAddress").selector]);
        expect(await diamond.getSelectorType(builder("make").selector)).to.eq(SelectorType.Undefined);
        expect(await diamond.getSelectorType(builder("getCallerAddress").selector)).to.eq(SelectorType.SwapDiamond);
      });

      it("should update selectors to arbitrary type selectors properly", async () => {
        await diamond["updateFacet(address,bytes4[],bytes4[],uint8[])"](
          master.address,
          [builder("make").selector],
          [builder("getCallerAddress").selector],
          [SelectorType.MasterRouter]
        );

        expect(await diamond.getFacetSelectors(master.address)).to.deep.eq([builder("getCallerAddress").selector]);
        expect(await diamond.getSelectorType(builder("make").selector)).to.eq(SelectorType.Undefined);
        expect(await diamond.getSelectorType(builder("getCallerAddress").selector)).to.eq(SelectorType.MasterRouter);
      });

      it("should revert if arguments lengths mismatch", async () => {
        const tx = diamond["updateFacet(address,bytes4[],bytes4[],uint8[])"](
          master.address,
          [builder("make").selector],
          [builder("getCallerAddress").selector],
          []
        );

        await expect(tx).to.be.revertedWith("SwapDiamond: lengths mismatch");
      });
    });

    describe("#beforeFallback", () => {
      it("should not call MasterRouter facet", async () => {
        await expect(multicallProxy.multicall([], [], [])).to.be.revertedWith("SwapDiamond: wrong selector type");
      });
    });
  });
});
