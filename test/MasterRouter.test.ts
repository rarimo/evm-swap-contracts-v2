import { Reverter } from "./helpers/reverter";
import { Builder, getBuilder } from "./utils/builder";
import {
  BridgeRouter__factory,
  MasterRouter__factory,
  MulticallRouter__factory,
  SwapDiamond__factory,
  TraderJoeRouter__factory,
  TransferRouter__factory,
  UniswapV2Router__factory,
  UniswapV3Router__factory,
  WETH9Mock__factory,
  WrapRouter__factory,
  SwapDiamond,
  MasterRouter,
} from "../generated-types/ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { SelectorType } from "./utils/contants";

describe("SwapDiamond", async () => {
  const reverter = new Reverter();

  let builder: Builder;

  let OWNER: SignerWithAddress;
  let CALLER: SignerWithAddress;

  let diamond: SwapDiamond;
  let master: MasterRouter;

  before("setup", async () => {
    builder = await getBuilder();

    [OWNER, CALLER] = await ethers.getSigners();

    const SwapDiamond = await ethers.getContractFactory("SwapDiamond");
    const MasterRouter = await ethers.getContractFactory("MasterRouter");

    diamond = await SwapDiamond.deploy();
    master = await MasterRouter.deploy();

    await diamond["addFacet(address,bytes4[],uint8[])"](
      master.address,
      [
        "0x00000000",
        builder("make").selector,
        builder("onERC721Received").selector,
        builder("onERC1155Received").selector,
        builder("onERC1155BatchReceived").selector,
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
      ]
    );

    master = await MasterRouter.attach(diamond.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#onERC721Received", function () {});

  describe("#onERC721Received", function () {});

  describe("#onERC721Received", function () {});

  describe("#supportsInterface", function () {});
});
