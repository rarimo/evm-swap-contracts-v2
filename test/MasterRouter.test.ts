import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  MasterRouter__factory,
  SwapDiamond__factory,
  SwapDiamond,
  MasterRouter,
  ERC721MintableBurnable,
  ERC721MintableBurnable__factory,
  ERC1155MintableBurnable__factory,
  ERC1155MintableBurnable,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { SelectorType } from "./utils/contants";
import { wei } from "../scripts/utils/utils";

describe("SwapDiamond", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let SwapDiamondFactory: SwapDiamond__factory;
  let MasterRouterFactory: MasterRouter__factory;
  let ERC721MintableBurnableFactory: ERC721MintableBurnable__factory;
  let ERC1155MintableBurnableFactory: ERC1155MintableBurnable__factory;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let erc721: ERC721MintableBurnable;
  let erc1155: ERC1155MintableBurnable;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER] = await ethers.getSigners();

    SwapDiamondFactory = await ethers.getContractFactory("SwapDiamond");
    MasterRouterFactory = await ethers.getContractFactory("MasterRouter");
    ERC721MintableBurnableFactory = await ethers.getContractFactory("ERC721MintableBurnable");
    ERC1155MintableBurnableFactory = await ethers.getContractFactory("ERC1155MintableBurnable");

    diamond = await SwapDiamondFactory.deploy();
    master = await MasterRouterFactory.deploy();

    erc721 = await ERC721MintableBurnableFactory.deploy("ERC721Mock", "ERC721Mock", OWNER.address, "");
    await erc721.mintTo(OWNER.address, 1, "");

    erc1155 = await ERC1155MintableBurnableFactory.deploy(OWNER.address, "");
    await erc1155.mintTo(OWNER.address, 1, 1, "");
    await erc1155.mintTo(OWNER.address, 2, 2, "");

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#receive", () => {
    it("should not receive if selector is not added", async () => {
      const tx = OWNER.sendTransaction({
        to: diamond.address,
        value: wei("1"),
      });

      await expect(tx).to.be.revertedWith("Diamond: selector is not registered");
    });

    it("should receive if selector is added", async () => {
      await diamond["addFacet(address,bytes4[],uint8[])"](master.address, ["0x00000000"], [SelectorType.SwapDiamond]);

      const tx = OWNER.sendTransaction({
        to: diamond.address,
        value: wei("1"),
      });

      await expect(tx).to.changeEtherBalances([OWNER, diamond], [-wei("1"), wei("1")]);
    });
  });

  describe("#onERC721Received", () => {
    it("should not safeTransferFrom if selector is not added", async () => {
      const tx = erc721["safeTransferFrom(address,address,uint256)"](OWNER.address, diamond.address, 1);

      await expect(tx).to.be.revertedWith("Diamond: selector is not registered");
    });

    it("should not safeTransferFrom with data if selector is not added", async () => {
      const tx = erc721["safeTransferFrom(address,address,uint256,bytes)"](OWNER.address, diamond.address, 1, "0x1337");

      await expect(tx).to.be.revertedWith("Diamond: selector is not registered");
    });

    context("selector is added", () => {
      beforeEach(async () => {
        await diamond["addFacet(address,bytes4[],uint8[])"](
          master.address,
          [builder("onERC721Received").selector],
          [SelectorType.SwapDiamond]
        );
      });

      it("should safeTransferFrom if selector is added", async () => {
        const tx = erc721["safeTransferFrom(address,address,uint256)"](OWNER.address, diamond.address, 1);

        await expect(tx).to.changeTokenBalances(erc721, [OWNER, diamond], [-1, 1]);
      });

      it("should safeTransferFrom with data if selector is added", async () => {
        const tx = erc721["safeTransferFrom(address,address,uint256)"](OWNER.address, diamond.address, 1);

        await expect(tx).to.changeTokenBalances(erc721, [OWNER, diamond], [-1, 1]);
      });
    });
  });

  describe("#onERC1155Received", () => {
    it("should not safeTransferFrom if selector is not added", async () => {
      const tx = erc1155.safeTransferFrom(OWNER.address, diamond.address, 1, 1, "0x");

      await expect(tx).to.be.revertedWith("Diamond: selector is not registered");
    });

    it("should safeTransferFrom if selector is added", async () => {
      await diamond["addFacet(address,bytes4[],uint8[])"](
        master.address,
        [builder("onERC1155Received").selector],
        [SelectorType.SwapDiamond]
      );

      await erc1155.safeTransferFrom(OWNER.address, diamond.address, 1, 1, "0x");

      expect(await erc1155.balanceOf(OWNER.address, 1)).to.eq(0n);
      expect(await erc1155.balanceOf(diamond.address, 1)).to.eq(1n);
    });
  });

  describe("#onERC1155BatchReceived", () => {
    it("should not safeBatchTransferFrom if selector is not added", async () => {
      const tx = erc1155.safeBatchTransferFrom(OWNER.address, diamond.address, [1, 2], [1, 2], "0x");

      await expect(tx).to.be.revertedWith("Diamond: selector is not registered");
    });

    it("should safeBatchTransferFrom if selector is added", async () => {
      await diamond["addFacet(address,bytes4[],uint8[])"](
        master.address,
        [builder("onERC1155BatchReceived").selector],
        [SelectorType.SwapDiamond]
      );

      await erc1155.safeBatchTransferFrom(OWNER.address, diamond.address, [1, 2], [1, 2], "0x");

      expect(await erc1155.balanceOf(OWNER.address, 1)).to.eq(0);
      expect(await erc1155.balanceOf(OWNER.address, 2)).to.eq(0);
      expect(await erc1155.balanceOf(diamond.address, 1)).to.eq(1);
      expect(await erc1155.balanceOf(diamond.address, 2)).to.eq(2);
    });
  });
});
