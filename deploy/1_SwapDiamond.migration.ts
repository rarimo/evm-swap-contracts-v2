import { Deployer, Logger } from "@dlsl/hardhat-migrate";
import { artifacts } from "hardhat";
import { DEX_TYPE, DexType } from "./config/env-utils";

const SwapDiamond = artifacts.require("SwapDiamond");
const MasterRouter = artifacts.require("MasterRouter");
const BridgeRouter = artifacts.require("BridgeRouter");
const MulticallRouter = artifacts.require("MulticallRouter");
const TransferRouter = artifacts.require("TransferRouter");
const WrapRouter = artifacts.require("WrapRouter");
const UniswapV2Router = artifacts.require("UniswapV2Router");
const UniswapV3Router = artifacts.require("UniswapV3Router");
const TraderJoeRouter = artifacts.require("TraderJoeRouter");

export = async (deployer: Deployer, logger: Logger) => {
  const dexType = DEX_TYPE();

  const contracts = new Array<[string, string]>();

  contracts.push(["SwapDiamond", (await deployer.deploy(SwapDiamond)).address]);
  contracts.push(["MasterRouter", (await deployer.deploy(MasterRouter)).address]);
  contracts.push(["BridgeRouter", (await deployer.deploy(BridgeRouter)).address]);
  contracts.push(["MulticallRouter", (await deployer.deploy(MulticallRouter)).address]);
  contracts.push(["TransferRouter", (await deployer.deploy(TransferRouter)).address]);
  contracts.push(["WrapRouter", (await deployer.deploy(WrapRouter)).address]);

  if (dexType === DexType.Uniswap) {
    contracts.push(["UniswapV2Router", (await deployer.deploy(UniswapV2Router)).address]);
    contracts.push(["UniswapV3Router", (await deployer.deploy(UniswapV3Router)).address]);
  } else {
    contracts.push(["TraderJoeRouter", (await deployer.deploy(TraderJoeRouter)).address]);
  }

  logger.logContracts(...contracts);
};
