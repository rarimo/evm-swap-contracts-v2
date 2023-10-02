import { Deployer, Logger } from "@solarity/hardhat-migrate";
import { artifacts } from "hardhat";
import config from "@/deploy/config/config.json";

const SwapDiamond = artifacts.require("SwapDiamond");

export = async (deployer: Deployer, logger: Logger) => {
  logger.logTransaction(
    await (await SwapDiamond.deployed()).transferOwnership(config.accounts.owner),
    "Transfer SwapDiamond ownership"
  );
};
