import { Deployer, Logger } from "@dlsl/hardhat-migrate";
import { artifacts } from "hardhat";
import config from "./config/config.json";

const SwapDiamond = artifacts.require("SwapDiamond");

export = async (deployer: Deployer, logger: Logger) => {
  logger.logTransaction(
    await (await SwapDiamond.deployed()).transferOwnership(config.accounts.OWNER),
    "Transfer SwapDiamond ownership"
  );
};
