import { ethers } from "hardhat";
import { BigNumber } from "ethers";

export function wei(value: string | number | bigint, decimal: number = 18): bigint {
  if (typeof value == "number" || typeof value == "bigint") {
    value = value.toString();
  }

  return ethers.utils.parseUnits(value as string, decimal).toBigInt();
}

export function fromWei(value: string | number | bigint, decimal: number = 18) {
  return BigNumber.from(value).div(BigNumber.from(10).pow(decimal));
}

export async function account(index: number) {
  return (await ethers.getSigners())[index];
}
