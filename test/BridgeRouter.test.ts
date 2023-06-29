import { expect } from "chai";
import { Reverter } from "@/test/helpers/reverter";
import { Builder, getBuilder } from "@/test/utils/builder";
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
  BridgeFacade,
} from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import {
  COMMISSION_ADDRESS,
  CONTRACT_BALANCE,
  ETHEREUM_ADDRESS,
  FacadeMethodId,
  SelectorType,
} from "@/test/utils/contants";
import { wei } from "@/scripts/utils/utils";
import { ZERO_ADDR } from "@/scripts/utils/constants";
import { BigNumberish } from "ethers";
import { Wallet } from "@ethersproject/wallet";

describe("BridgeRouter", () => {
  const chainName = "ethereum";
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000001";
  const receiver = "receiver";

  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;
  let SIGNER: Wallet;

  let diamond: SwapDiamond;
  let master: MasterRouter;
  let masterProxy: MasterRouter;
  let transfer: TransferRouter;
  let bridge: BridgeRouter;
  let bridgeProxy: BridgeRouter;
  let rarimoFacade: BridgeFacade;
  let rarimoBridge: Bridge;
  let bundle: BundleExecutorImplementation;
  let feeToken: ERC20MintableBurnable;
  let feeTokens: Array<string>;
  let feeAmounts: Array<bigint>;
  let erc20: ERC20MintableBurnable;
  let erc721: ERC721MintableBurnable;
  let erc1155: ERC1155MintableBurnable;

  function encodeSalt(origin: SignerWithAddress, salt_: string = salt): string {
    return ethers.utils.solidityKeccak256(["bytes32", "address"], [salt_, origin.address]);
  }

  function signAddFeeToken(feeTokens: Array<string>, feeAmounts: Array<BigNumberish>, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint256", "address", "string", "uint8", "address[]", "uint256[]"],
      [nonce, rarimoFacade.address, chainName, FacadeMethodId.AddFeeToken, feeTokens, feeAmounts]
    );

    return ethers.utils.joinSignature(SIGNER._signingKey().signDigest(hash));
  }

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER] = await ethers.getSigners();
    SIGNER = ethers.Wallet.createRandom();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");
    const TransferRouter = await ethers.getContractFactory("TransferRouter");
    const BridgeRouter = await ethers.getContractFactory("BridgeRouter");
    const RarimoFacade = await ethers.getContractFactory("BridgeFacade");
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
    rarimoFacade = await RarimoFacade.deploy();
    rarimoBridge = await RarimoBridge.deploy();
    bundle = await BundleExecutorImplementation.deploy(ZERO_ADDR);

    await rarimoBridge.__Bridge_init(SIGNER.address, bundle.address, chainName, rarimoFacade.address);
    await rarimoFacade.__BridgeFacade_init(rarimoBridge.address);

    erc20 = await ERC20MintableBurnable.deploy("ERC20Mock", "ERC20Mock", 18, OWNER.address);

    await erc20.mintTo(CALLER.address, wei("100"));
    await erc20.connect(CALLER).approve(diamond.address, wei("100"));

    erc721 = await ERC721MintableBurnable.deploy("ERC721Mock", "ERC721Mock", OWNER.address, "");

    await erc721.mintTo(CALLER.address, 1, "");
    await erc721.mintTo(CALLER.address, 2, "");
    await erc721.mintTo(CALLER.address, 3, "");
    await erc721.connect(CALLER).setApprovalForAll(diamond.address, true);

    erc1155 = await ERC1155MintableBurnable.deploy(OWNER.address, "");

    await erc1155.mintTo(CALLER.address, 1, wei("100"), "");
    await erc1155.mintTo(CALLER.address, 2, wei("200"), "");
    await erc1155.connect(CALLER).setApprovalForAll(diamond.address, true);

    feeToken = await ERC20MintableBurnable.deploy("ERC20FeeMock", "ERC20FeeMock", 18, OWNER.address);

    await feeToken.mintTo(CALLER.address, wei("2"));
    await feeToken.connect(CALLER).approve(diamond.address, wei("2"));

    feeTokens = [ETHEREUM_ADDRESS, feeToken.address, COMMISSION_ADDRESS, erc20.address];
    feeAmounts = [wei("1"), wei("2"), 1n, wei("10")];

    const signature = signAddFeeToken(feeTokens, feeAmounts, 0);

    await rarimoFacade.addFeeToken({ feeTokens, feeAmounts, signature });

    await erc20.transferOwnership(rarimoBridge.address);
    await erc721.transferOwnership(rarimoBridge.address);
    await erc1155.transferOwnership(rarimoBridge.address);

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      [builder("make").selector, builder("onERC721Received").selector, builder("onERC1155BatchReceived").selector],
      [SelectorType.SwapDiamond, SelectorType.SwapDiamond, SelectorType.SwapDiamond]
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

    await bridgeProxy.setBridgeAddress(rarimoFacade.address);

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
    it("should revert if no funds to pay ERC20 commission", async () => {
      const tx = masterProxy.connect(CALLER).make([
        builder("transferFromERC20", [erc20.address, wei("1")]).payload(),
        builder("bridgeERC20", [
          {
            feeToken: feeTokens[3],
          },
          {
            token: erc20.address,
            amount: wei("1"),
            bundle: { salt: salt, bundle: "0x" },
            network: chainName,
            receiver: receiver,
            isWrapped: true,
          },
        ]).payload(),
      ]);

      await expect(tx).to.be.revertedWith("MasterRouter: ERC20: transfer amount exceeds balance");
    });

    it("should revert if commission exceeds balance", async () => {
      const tx = masterProxy.connect(CALLER).make([
        builder("transferFromERC20", [erc20.address, feeAmounts[3] - 1n]).payload(),
        builder("bridgeERC20", [
          {
            feeToken: feeTokens[3],
          },
          {
            token: erc20.address,
            amount: CONTRACT_BALANCE,
            bundle: { salt: salt, bundle: "0x" },
            network: chainName,
            receiver: receiver,
            isWrapped: true,
          },
        ]).payload(),
      ]);

      await expect(tx).to.be.revertedWith("MasterRouter: Resolver: commission exceeds balance");
    });

    it("should bridge erc20 with different commissions & resolve amount", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("transferFromERC20", [erc20.address, wei("100")]).payload(),
          builder("transferFromERC20", [feeToken.address, feeAmounts[1]]).payload(),
          builder("bridgeERC20", [
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc20.address,
              amount: wei("1"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC20", [
            {
              feeToken: feeTokens[1],
            },
            {
              token: erc20.address,
              amount: wei("2"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC20", [
            {
              feeToken: feeTokens[2],
            },
            {
              token: erc20.address,
              amount: wei("3"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC20", [
            {
              feeToken: feeTokens[3],
            },
            {
              token: erc20.address,
              amount: wei("4"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC20", [
            {
              feeToken: feeTokens[3],
            },
            {
              token: erc20.address,
              amount: CONTRACT_BALANCE,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
        ],
        { value: feeAmounts[1] }
      );

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("1"), encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("2"), encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("3"), encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("4"), encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC20")
        .withArgs(erc20.address, wei("70"), encodeSalt(CALLER), "0x", chainName, receiver, true);
    });
  });

  describe("#bridgeERC721", () => {
    it("should bridge erc721 with different commissions", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("transferFromERC20", [feeToken.address, feeAmounts[1]]).payload(),
          builder("transferFromERC721", [erc721.address, [1, 2, 3]]).payload(),
          builder("bridgeERC721", [
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc721.address,
              tokenId: 1,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC721", [
            {
              feeToken: feeTokens[1],
            },
            {
              token: erc721.address,
              tokenId: 2,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC721", [
            {
              feeToken: feeTokens[2],
            },
            {
              token: erc721.address,
              tokenId: 3,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
        ],
        { value: feeAmounts[0] }
      );

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC721")
        .withArgs(erc721.address, 1, encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC721")
        .withArgs(erc721.address, 2, encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC721")
        .withArgs(erc721.address, 3, encodeSalt(CALLER), "0x", chainName, receiver, true);
    });
  });

  describe("#bridgeERC1155", () => {
    it("should bridge erc1155 with different commissions & resolve amount", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("transferFromERC20", [feeToken.address, feeAmounts[1]]).payload(),
          builder("transferFromERC1155", [erc1155.address, [1, 2], [wei("100"), wei("200")]]).payload(),
          builder("bridgeERC1155", [
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc1155.address,
              tokenId: 1,
              amount: wei("25"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC1155", [
            {
              feeToken: feeTokens[1],
            },
            {
              token: erc1155.address,
              tokenId: 1,
              amount: CONTRACT_BALANCE,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
          builder("bridgeERC1155", [
            {
              feeToken: feeTokens[2],
            },
            {
              token: erc1155.address,
              tokenId: 2,
              amount: CONTRACT_BALANCE,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
              isWrapped: true,
            },
          ]).payload(),
        ],
        { value: feeAmounts[1] }
      );

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedERC1155")
        .withArgs(erc1155.address, 1, wei("25"), encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC1155")
        .withArgs(erc1155.address, 1, wei("75"), encodeSalt(CALLER), "0x", chainName, receiver, true)
        .to.emit(rarimoBridge, "DepositedERC1155")
        .withArgs(erc1155.address, 2, wei("200"), encodeSalt(CALLER), "0x", chainName, receiver, true);
    });
  });

  describe("#bridgeNative", () => {
    it("should revert if no funds to pay native commission", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("bridgeNative", [
            {
              feeToken: feeTokens[0],
            },
            {
              amount: wei("10"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
            },
          ]).payload(),
        ],
        { value: wei("10") }
      );

      await expect(tx).to.be.revertedWith("MasterRouter: ErrorHelper: command reverted silently");
    });

    it("should revert if commission exceeds balance", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("bridgeNative", [
            {
              feeToken: feeTokens[0],
            },
            {
              amount: CONTRACT_BALANCE,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
            },
          ]).payload(),
        ],
        { value: feeAmounts[0] - 1n }
      );

      await expect(tx).to.be.revertedWith("MasterRouter: Resolver: commission exceeds balance");
    });

    it("should bridge native with different commissions & resolve amount", async () => {
      const tx = masterProxy.connect(CALLER).make(
        [
          builder("transferFromERC20", [feeToken.address, feeAmounts[1]]).payload(),
          builder("bridgeNative", [
            {
              feeToken: feeTokens[1],
            },
            {
              amount: wei("1"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
            },
          ]).payload(),
          builder("bridgeNative", [
            {
              feeToken: feeTokens[2],
            },
            {
              amount: wei("2"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
            },
          ]).payload(),
          builder("bridgeNative", [
            {
              feeToken: feeTokens[0],
            },
            {
              amount: wei("3"),
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
            },
          ]).payload(),
          builder("bridgeNative", [
            {
              feeToken: feeTokens[0],
            },
            {
              amount: CONTRACT_BALANCE,
              bundle: { salt: salt, bundle: "0x" },
              network: chainName,
              receiver: receiver,
            },
          ]).payload(),
        ],
        { value: wei("30") }
      );

      await expect(tx)
        .to.emit(rarimoBridge, "DepositedNative")
        .withArgs(wei("1"), encodeSalt(CALLER), "0x", chainName, receiver)
        .to.emit(rarimoBridge, "DepositedNative")
        .withArgs(wei("2"), encodeSalt(CALLER), "0x", chainName, receiver)
        .to.emit(rarimoBridge, "DepositedNative")
        .withArgs(wei("3"), encodeSalt(CALLER), "0x", chainName, receiver)
        .to.emit(rarimoBridge, "DepositedNative")
        .withArgs(wei("22"), encodeSalt(CALLER), "0x", chainName, receiver);
    });
  });
});
