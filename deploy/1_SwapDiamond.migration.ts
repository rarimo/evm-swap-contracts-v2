import { Deployer, Logger } from "@dlsl/hardhat-migrate";
import { artifacts } from "hardhat";
import config from "./config/config.json";

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
  const contracts = new Array<[string, string]>();

  contracts.push(["SwapDiamond", (await deployer.deploy(SwapDiamond)).address]);
  contracts.push(["MasterRouter", (await deployer.deploy(MasterRouter)).address]);

  if (config.facets.bridge) {
    contracts.push(["BridgeRouter", (await deployer.deploy(BridgeRouter)).address]);
  }

  if (config.facets.multicall) {
    contracts.push(["MulticallRouter", (await deployer.deploy(MulticallRouter)).address]);
  }

  if (config.facets.transfer) {
    contracts.push(["TransferRouter", (await deployer.deploy(TransferRouter)).address]);
  }

  if (config.facets.wrap) {
    contracts.push(["WrapRouter", (await deployer.deploy(WrapRouter)).address]);
  }

  if (config.facets.uniswapV2) {
    contracts.push(["UniswapV2Router", (await deployer.deploy(UniswapV2Router)).address]);
  }

  if (config.facets.uniswapV3) {
    contracts.push(["UniswapV3Router", (await deployer.deploy(UniswapV3Router)).address]);
  }

  if (config.facets.traderJoe) {
    contracts.push(["TraderJoeRouter", (await deployer.deploy(TraderJoeRouter)).address]);
  }

  logger.logContracts(...contracts);
};
