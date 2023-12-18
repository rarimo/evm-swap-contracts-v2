/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../common";

export declare namespace IBundler {
  export type BundleStruct = {
    salt: PromiseOrValue<BytesLike>;
    bundle: PromiseOrValue<BytesLike>;
  };

  export type BundleStructOutput = [string, string] & {
    salt: string;
    bundle: string;
  };
}

export declare namespace INativeHandler {
  export type DepositNativeParametersStruct = {
    amount: PromiseOrValue<BigNumberish>;
    bundle: IBundler.BundleStruct;
    network: PromiseOrValue<string>;
    receiver: PromiseOrValue<string>;
  };

  export type DepositNativeParametersStructOutput = [
    BigNumber,
    IBundler.BundleStructOutput,
    string,
    string
  ] & {
    amount: BigNumber;
    bundle: IBundler.BundleStructOutput;
    network: string;
    receiver: string;
  };

  export type WithdrawNativeParametersStruct = {
    amount: PromiseOrValue<BigNumberish>;
    bundle: IBundler.BundleStruct;
    originHash: PromiseOrValue<BytesLike>;
    receiver: PromiseOrValue<string>;
    proof: PromiseOrValue<BytesLike>;
  };

  export type WithdrawNativeParametersStructOutput = [
    BigNumber,
    IBundler.BundleStructOutput,
    string,
    string,
    string
  ] & {
    amount: BigNumber;
    bundle: IBundler.BundleStructOutput;
    originHash: string;
    receiver: string;
    proof: string;
  };
}

export interface INativeHandlerInterface extends utils.Interface {
  functions: {
    "depositNative((uint256,(bytes32,bytes),string,string))": FunctionFragment;
    "determineProxyAddress(bytes32)": FunctionFragment;
    "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))": FunctionFragment;
    "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "depositNative"
      | "depositNative((uint256,(bytes32,bytes),string,string))"
      | "determineProxyAddress"
      | "determineProxyAddress(bytes32)"
      | "withdrawNative"
      | "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))"
      | "withdrawNativeBundle"
      | "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "depositNative",
    values: [INativeHandler.DepositNativeParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "depositNative((uint256,(bytes32,bytes),string,string))",
    values: [INativeHandler.DepositNativeParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "determineProxyAddress",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "determineProxyAddress(bytes32)",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawNative",
    values: [INativeHandler.WithdrawNativeParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))",
    values: [INativeHandler.WithdrawNativeParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawNativeBundle",
    values: [INativeHandler.WithdrawNativeParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))",
    values: [INativeHandler.WithdrawNativeParametersStruct]
  ): string;

  decodeFunctionResult(
    functionFragment: "depositNative",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositNative((uint256,(bytes32,bytes),string,string))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "determineProxyAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "determineProxyAddress(bytes32)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawNative",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawNativeBundle",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))",
    data: BytesLike
  ): Result;

  events: {
    "DepositedNative(uint256,bytes32,bytes,string,string)": EventFragment;
    "WithdrawnNative(uint256,bytes32,bytes,bytes32,address,bytes)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "DepositedNative"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "DepositedNative(uint256,bytes32,bytes,string,string)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "WithdrawnNative"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "WithdrawnNative(uint256,bytes32,bytes,bytes32,address,bytes)"
  ): EventFragment;
}

export interface DepositedNativeEventObject {
  amount: BigNumber;
  salt: string;
  bundle: string;
  network: string;
  receiver: string;
}
export type DepositedNativeEvent = TypedEvent<
  [BigNumber, string, string, string, string],
  DepositedNativeEventObject
>;

export type DepositedNativeEventFilter = TypedEventFilter<DepositedNativeEvent>;

export interface WithdrawnNativeEventObject {
  amount: BigNumber;
  salt: string;
  bundle: string;
  originHash: string;
  receiver: string;
  proof: string;
}
export type WithdrawnNativeEvent = TypedEvent<
  [BigNumber, string, string, string, string, string],
  WithdrawnNativeEventObject
>;

export type WithdrawnNativeEventFilter = TypedEventFilter<WithdrawnNativeEvent>;

export interface INativeHandler extends BaseContract {
  contractName: "INativeHandler";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: INativeHandlerInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    depositNative(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "depositNative((uint256,(bytes32,bytes),string,string))"(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    withdrawNative(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawNativeBundle(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  depositNative(
    params_: INativeHandler.DepositNativeParametersStruct,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "depositNative((uint256,(bytes32,bytes),string,string))"(
    params_: INativeHandler.DepositNativeParametersStruct,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  determineProxyAddress(
    salt_: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  "determineProxyAddress(bytes32)"(
    salt_: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  withdrawNative(
    params_: INativeHandler.WithdrawNativeParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))"(
    params_: INativeHandler.WithdrawNativeParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawNativeBundle(
    params_: INativeHandler.WithdrawNativeParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))"(
    params_: INativeHandler.WithdrawNativeParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    depositNative(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "depositNative((uint256,(bytes32,bytes),string,string))"(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    withdrawNative(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    withdrawNativeBundle(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "DepositedNative(uint256,bytes32,bytes,string,string)"(
      amount?: null,
      salt?: null,
      bundle?: null,
      network?: null,
      receiver?: null
    ): DepositedNativeEventFilter;
    DepositedNative(
      amount?: null,
      salt?: null,
      bundle?: null,
      network?: null,
      receiver?: null
    ): DepositedNativeEventFilter;

    "WithdrawnNative(uint256,bytes32,bytes,bytes32,address,bytes)"(
      amount?: null,
      salt?: null,
      bundle?: null,
      originHash?: null,
      receiver?: null,
      proof?: null
    ): WithdrawnNativeEventFilter;
    WithdrawnNative(
      amount?: null,
      salt?: null,
      bundle?: null,
      originHash?: null,
      receiver?: null,
      proof?: null
    ): WithdrawnNativeEventFilter;
  };

  estimateGas: {
    depositNative(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "depositNative((uint256,(bytes32,bytes),string,string))"(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    withdrawNative(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawNativeBundle(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    depositNative(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "depositNative((uint256,(bytes32,bytes),string,string))"(
      params_: INativeHandler.DepositNativeParametersStruct,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    withdrawNative(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawNativeBundle(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes))"(
      params_: INativeHandler.WithdrawNativeParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
