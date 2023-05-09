import { Deployer, Logger } from "@dlsl/hardhat-migrate";
import { artifacts } from "hardhat";
import config from "./config/config.json";

const SwapDiamond = artifacts.require("SwapDiamond");
const BridgeRouter = artifacts.require("BridgeRouter");
const WrapRouter = artifacts.require("WrapRouter");
const UniswapV2Router = artifacts.require("UniswapV2Router");
const UniswapV3Router = artifacts.require("UniswapV3Router");
const TraderJoeRouter = artifacts.require("TraderJoeRouter");

export = async (deployer: Deployer, logger: Logger) => {
  const diamondAddress = (await SwapDiamond.deployed()).address;

  if (config.facets.bridge) {
    logger.logTransaction(
      await (await BridgeRouter.at(diamondAddress)).setBridgeAddress(config.integrations.evmBridge),
      "Set bridge address"
    );
  }

  if (config.facets.wrap) {
    logger.logTransaction(
      await (await WrapRouter.at(diamondAddress)).setWrappedNativeAddress(config.integrations.wrappedNative),
      "Set wrapped native address"
    );
  }

  if (config.facets.uniswapV2) {
    logger.logTransaction(
      await (await UniswapV2Router.at(diamondAddress)).setUniswapV2RouterAddress(config.integrations.uniswapV2Router),
      "Set UniswapV2 router address"
    );
  }

  if (config.facets.uniswapV3) {
    logger.logTransaction(
      await (await UniswapV3Router.at(diamondAddress)).setUniswapV3RouterAddress(config.integrations.uniswapV3Router),
      "Set UniswapV3 router address"
    );
  }

  if (config.facets.traderJoe) {
    logger.logTransaction(
      await (await TraderJoeRouter.at(diamondAddress)).setTraderJoeRouterAddress(config.integrations.traderJoeRouter),
      "Set TraderJoe router address"
    );
  }
};
