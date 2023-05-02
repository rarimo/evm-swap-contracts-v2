import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  SwapDiamond,
  MasterRouter,
  ERC721MintableBurnable,
  ERC1155MintableBurnable,
  ERC20MintableBurnable,
  TransferRouter,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CALLER_ADDRESS, CONTRACT_BALANCE, SelectorType, THIS_ADDRESS } from "./utils/contants";
import { wei } from "../scripts/utils/utils";

describe("TransferRouter", () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let RECEIVER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let transfer: TransferRouter;
  let erc20: ERC20MintableBurnable;
  let erc721: ERC721MintableBurnable;
  let erc1155: ERC1155MintableBurnable;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const TransferRouter = await ethers.getContractFactory("TransferRouter");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");
    const ERC721MintableBurnable = await ethers.getContractFactory("ERC721MintableBurnable");
    const ERC1155MintableBurnable = await ethers.getContractFactory("ERC1155MintableBurnable");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    transfer = await TransferRouter.deploy();

    erc20 = await ERC20MintableBurnable.deploy("ERC20Mock", "ERC20Mock", OWNER.address);

    await erc20.mintTo(CALLER.address, wei("100"));
    await erc20.connect(CALLER).approve(diamond.address, wei("100"));

    erc721 = await ERC721MintableBurnable.deploy("ERC721Mock", "ERC721Mock", OWNER.address, "");

    for (let i = 1; i <= 5; i++) {
      await erc721.mintTo(CALLER.address, i, "");
    }

    await erc721.connect(CALLER).setApprovalForAll(diamond.address, true);

    erc1155 = await ERC1155MintableBurnable.deploy(OWNER.address, "");

    await erc1155.mintTo(CALLER.address, 1, wei("100"), "");
    await erc1155.connect(CALLER).setApprovalForAll(diamond.address, true);

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      [
        builder("make").selector,
        builder("onERC721Received").selector,
        builder("onERC1155BatchReceived").selector,
        builder("getCallerAddress").selector,
      ],
      [SelectorType.SwapDiamond, SelectorType.SwapDiamond, SelectorType.SwapDiamond, SelectorType.SwapDiamond]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      transfer.address,
      [
        builder("transferERC20").selector,
        builder("transferERC721").selector,
        builder("transferERC1155").selector,
        builder("transferNative").selector,
        builder("transferFromERC20").selector,
        builder("transferFromERC721").selector,
        builder("transferFromERC1155").selector,
      ],
      [
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#transferFromERC20 & #transferERC20", () => {
    it("should transfer erc20 and resolve receiver & amount properly", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [erc20.address, wei("100")]).payload(),
          builder("transferERC20", [erc20.address, RECEIVER.address, wei("20")]).payload(),
          builder("transferERC20", [erc20.address, CALLER_ADDRESS, CONTRACT_BALANCE]).payload(),
        ]);

      await expect(tx).to.changeTokenBalances(erc20, [RECEIVER, CALLER], [wei("20"), -wei("20")]);
    });

    it("should not transfer erc20 from this", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [erc20.address, wei("100")]).payload(),
          builder("transferERC20", [erc20.address, THIS_ADDRESS, wei("50")]).payload(),
        ]);

      await expect(tx).to.changeTokenBalances(erc20, [CALLER, diamond], [-wei("100"), wei("100")]);
    });
  });

  describe("#transferFromERC721 & #transferERC721", () => {
    it("should transfer erc721 and resolve receiver & amount properly", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC721", [erc721.address, [1, 2, 3, 4, 5]]).payload(),
          builder("transferERC721", [erc721.address, RECEIVER.address, [1, 2, 3]]).payload(),
          builder("transferERC721", [erc721.address, CALLER_ADDRESS, [4]]).payload(),
          builder("transferERC721", [erc721.address, THIS_ADDRESS, [5]]).payload(),
        ]);

      await expect(tx).to.changeTokenBalances(erc721, [RECEIVER, CALLER, diamond], [3, -4, 1]);
    });
  });

  describe("#transferFromERC1155 & #transferERC1155", () => {
    it("should transfer erc1155 and resolve receiver & amount properly", async () => {
      await masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC1155", [erc1155.address, [1], [wei("100")]]).payload(),
          builder("transferERC1155", [erc1155.address, RECEIVER.address, [1], [wei("20")]]).payload(),
          builder("transferERC1155", [erc1155.address, CALLER_ADDRESS, [1], [CONTRACT_BALANCE]]).payload(),
        ]);

      expect(await erc1155.balanceOf(RECEIVER.address, 1)).to.eq(wei("20"));
      expect(await erc1155.balanceOf(CALLER.address, 1)).to.eq(wei("80"));
    });

    it("should not transfer erc1155 from this", async () => {
      await masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC1155", [erc1155.address, [1], [wei("100")]]).payload(),
          builder("transferERC1155", [erc1155.address, THIS_ADDRESS, [1], [wei("50")]]).payload(),
        ]);

      expect(await erc1155.balanceOf(CALLER.address, 1)).to.eq(0);
      expect(await erc1155.balanceOf(diamond.address, 1)).to.eq(wei("100"));
    });

    it("should not transfer erc1155 if ids and amounts length mismatch", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC1155", [erc1155.address, [1], [wei("100")]]).payload(),
          builder("transferERC1155", [erc1155.address, THIS_ADDRESS, [1], []]).payload(),
        ]);

      await expect(tx).to.be.revertedWith("MasterRouter: TransferRouter: lengths mismatch");
    });
  });

  describe("#transferNative", () => {
    it("should transfer native and resolve receiver & amount properly", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("transferNative", [RECEIVER.address, wei("0.2")]).payload(),
            builder("transferNative", [CALLER_ADDRESS, CONTRACT_BALANCE]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx).to.changeEtherBalances([RECEIVER, CALLER], [wei("0.2"), -wei("0.2")]);
    });

    it("should not transfer native from this", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([builder("transferNative", [THIS_ADDRESS, wei("0.5")]).payload()], { value: wei("1") });

      await expect(tx).to.changeEtherBalances([CALLER, diamond], [-wei("1"), wei("1")]);
    });

    it("should not transfer native if bad receiver", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([builder("transferNative", [transfer.address, wei("0.5")]).payload()], { value: wei("1") });

      await expect(tx).to.be.revertedWith("MasterRouter: TransferRouter: failed to transfer native");
    });
  });
});
