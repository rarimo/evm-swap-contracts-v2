import { expect } from "chai";
import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  SwapDiamond,
  MasterRouter,
  ERC721MintableBurnable,
  ERC1155MintableBurnable,
  ERC20MintableBurnable,
  BridgeRouter,
  BundleExecutorImplementation,
  Bridge,
  TransferRouter,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CONTRACT_BALANCE, SelectorType } from "./utils/contants";
import { wei } from "../scripts/utils/utils";

describe("BridgeRouter", () => {
  const chainName = "ethereum";
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000001";
  const receiver = "receiver";

  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let RECEIVER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let transfer: TransferRouter;
  let bridge: BridgeRouter;
  let bridgeProxy: BridgeRouter;
  let rarimoBridge: Bridge;
  let bundle: BundleExecutorImplementation;
  let erc20: ERC20MintableBurnable;
  let erc721: ERC721MintableBurnable;
  let erc1155: ERC1155MintableBurnable;

  function encodeSalt(origin: SignerWithAddress, salt_: string = salt): string {
    return ethers.utils.solidityKeccak256(["bytes32", "address"], [salt_, origin.address]);
  }

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER, RECEIVER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const TransferRouter = await ethers.getContractFactory("TransferRouter");
    const BridgeRouter = await ethers.getContractFactory("BridgeRouter");
    const RarimoBridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");
    const ERC721MintableBurnable = await ethers.getContractFactory("ERC721MintableBurnable");
    const ERC1155MintableBurnable = await ethers.getContractFactory("ERC1155MintableBurnable");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();
    masterProxy = await MasterRouter.attach(diamond.address);
    transfer = await TransferRouter.deploy();
    bridge = await BridgeRouter.deploy();
    bridgeProxy = await BridgeRouter.attach(diamond.address);
    rarimoBridge = await RarimoBridge.deploy();
    bundle = await BundleExecutorImplementation.deploy();

    await rarimoBridge.__Bridge_init(OWNER.address, bundle.address, chainName);

    erc20 = await ERC20MintableBurnable.deploy("ERC20Mock", "ERC20Mock", OWNER.address);

    await erc20.mintTo(CALLER.address, wei("100"));
    await erc20.connect(CALLER).approve(diamond.address, wei("100"));

    erc721 = await ERC721MintableBurnable.deploy("ERC721Mock", "ERC721Mock", OWNER.address, "");

    await erc721.mintTo(CALLER.address, 1, "");

    await erc721.connect(CALLER).setApprovalForAll(diamond.address, true);

    erc1155 = await ERC1155MintableBurnable.deploy(OWNER.address, "");

    await erc1155.mintTo(CALLER.address, 1, wei("100"), "");
    await erc1155.connect(CALLER).setApprovalForAll(diamond.address, true);

    await erc20.transferOwnership(rarimoBridge.address);
    await erc721.transferOwnership(rarimoBridge.address);
    await erc1155.transferOwnership(rarimoBridge.address);

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
        builder("transferFromERC20").selector,
        builder("transferFromERC721").selector,
        builder("transferFromERC1155").selector,
      ],
      [SelectorType.MasterRouter, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );
    await diamond["addFacet(address,bytes4[],uint8[])"](
      bridge.address,
      [
        builder("setBridgeAddress").selector,
        builder("bridgeERC20").selector,
        builder("bridgeERC721").selector,
        builder("bridgeERC1155").selector,
        builder("bridgeNative").selector,
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );

    await bridgeProxy.setBridgeAddress(rarimoBridge.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#setBridgeAddress", () => {
    it("should not set if caller is not the owner", async () => {
      await expect(bridgeProxy.connect(CALLER).setBridgeAddress(rarimoBridge.address)).to.be.revertedWith(
        "ODStorage: not an owner"
      );
    });
  });

  describe("#bridgeERC20", () => {
    it("should bridge erc20 & resolve amount", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC20", [erc20.address, wei("100")]).payload(),
          builder("bridgeERC20", [
            erc20.address,
            wei("1"),
            { salt: salt, bundle: "0x" },
            chainName,
            receiver,
            true,
          ]).payload(),
          builder("bridgeERC20", [
            erc20.address,
            CONTRACT_BALANCE,
            { salt: salt, bundle: "0x" },
            chainName,
            receiver,
            true,
          ]).payload(),
        ]);

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("1"), encodeSalt(CALLER), "0x", chainName, receiver, true);
      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("99"), encodeSalt(CALLER), "0x", chainName, receiver, true);
    });
  });

  describe("#bridgeERC721", () => {
    it("should bridge erc721", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC721", [erc721.address, [1]]).payload(),
          builder("bridgeERC721", [
            erc721.address,
            1,
            { salt: salt, bundle: "0x" },
            chainName,
            receiver,
            true,
          ]).payload(),
        ]);

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC721")
        .withArgs(erc721.address, 1, encodeSalt(CALLER), "0x", chainName, receiver, true);
    });
  });

  describe("#bridgeERC1155", () => {
    it("should bridge erc1155 & resolve amount", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make([
          builder("transferFromERC1155", [erc1155.address, [1], [wei("100")]]).payload(),
          builder("bridgeERC1155", [
            erc1155.address,
            1,
            wei("1"),
            { salt: salt, bundle: "0x" },
            chainName,
            receiver,
            true,
          ]).payload(),
          builder("bridgeERC1155", [
            erc1155.address,
            1,
            CONTRACT_BALANCE,
            { salt: salt, bundle: "0x" },
            chainName,
            receiver,
            true,
          ]).payload(),
        ]);

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC1155")
        .withArgs(erc1155.address, 1, wei("1"), encodeSalt(CALLER), "0x", chainName, receiver, true);
      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC1155")
        .withArgs(erc1155.address, 1, wei("99"), encodeSalt(CALLER), "0x", chainName, receiver, true);
    });
  });

  describe("#bridgeNative", () => {
    it("should bridge native & resolve amount", async () => {
      const tx = masterProxy
        .connect(CALLER)
        .make(
          [
            builder("bridgeNative", [wei("0.1"), { salt: salt, bundle: "0x" }, chainName, receiver]).payload(),
            builder("bridgeNative", [CONTRACT_BALANCE, { salt: salt, bundle: "0x" }, chainName, receiver]).payload(),
          ],
          { value: wei("1") }
        );

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedNative")
        .withArgs(wei("0.1"), encodeSalt(CALLER), "0x", chainName, receiver);
      await expect(tx)
        .to.emit(rarimoBridge, "DepositedNative")
        .withArgs(wei("0.9"), encodeSalt(CALLER), "0x", chainName, receiver);
    });
  });
});
