import * as dotenv from "dotenv";
dotenv.config();

export enum DexType {
  Uniswap = 1,
  TraderJoe,
}

export function DEX_TYPE() {
  if (process.env.DEX_TYPE === undefined || process.env.DEX_TYPE === "UNISWAP") {
    return DexType.Uniswap;
  }

  if (process.env.DEX_TYPE === "TRADER_JOE") {
    return DexType.TraderJoe;
  }

  throw Error("wrong dex type");
}
