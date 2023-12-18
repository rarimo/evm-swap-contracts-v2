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

export declare namespace IFeeManager {
  export type AddFeeTokenParametersStruct = {
    feeTokens: PromiseOrValue<string>[];
    feeAmounts: PromiseOrValue<BigNumberish>[];
    signature: PromiseOrValue<BytesLike>;
  };

  export type AddFeeTokenParametersStructOutput = [
    string[],
    BigNumber[],
    string
  ] & { feeTokens: string[]; feeAmounts: BigNumber[]; signature: string };

  export type RemoveFeeTokenParametersStruct = {
    feeTokens: PromiseOrValue<string>[];
    feeAmounts: PromiseOrValue<BigNumberish>[];
    signature: PromiseOrValue<BytesLike>;
  };

  export type RemoveFeeTokenParametersStructOutput = [
    string[],
    BigNumber[],
    string
  ] & { feeTokens: string[]; feeAmounts: BigNumber[]; signature: string };

  export type UpdateFeeTokenParametersStruct = {
    feeTokens: PromiseOrValue<string>[];
    feeAmounts: PromiseOrValue<BigNumberish>[];
    signature: PromiseOrValue<BytesLike>;
  };

  export type UpdateFeeTokenParametersStructOutput = [
    string[],
    BigNumber[],
    string
  ] & { feeTokens: string[]; feeAmounts: BigNumber[]; signature: string };

  export type WithdrawFeeTokenParametersStruct = {
    receiver: PromiseOrValue<string>;
    feeTokens: PromiseOrValue<string>[];
    amounts: PromiseOrValue<BigNumberish>[];
    signature: PromiseOrValue<BytesLike>;
  };

  export type WithdrawFeeTokenParametersStructOutput = [
    string,
    string[],
    BigNumber[],
    string
  ] & {
    receiver: string;
    feeTokens: string[];
    amounts: BigNumber[];
    signature: string;
  };
}

export interface IFeeManagerInterface extends utils.Interface {
  functions: {
    "addFeeToken((address[],uint256[],bytes))": FunctionFragment;
    "getCommission(address)": FunctionFragment;
    "removeFeeToken((address[],uint256[],bytes))": FunctionFragment;
    "updateFeeToken((address[],uint256[],bytes))": FunctionFragment;
    "withdrawFeeToken((address,address[],uint256[],bytes))": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "addFeeToken"
      | "addFeeToken((address[],uint256[],bytes))"
      | "getCommission"
      | "getCommission(address)"
      | "removeFeeToken"
      | "removeFeeToken((address[],uint256[],bytes))"
      | "updateFeeToken"
      | "updateFeeToken((address[],uint256[],bytes))"
      | "withdrawFeeToken"
      | "withdrawFeeToken((address,address[],uint256[],bytes))"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "addFeeToken",
    values: [IFeeManager.AddFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "addFeeToken((address[],uint256[],bytes))",
    values: [IFeeManager.AddFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "getCommission",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "getCommission(address)",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "removeFeeToken",
    values: [IFeeManager.RemoveFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "removeFeeToken((address[],uint256[],bytes))",
    values: [IFeeManager.RemoveFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "updateFeeToken",
    values: [IFeeManager.UpdateFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "updateFeeToken((address[],uint256[],bytes))",
    values: [IFeeManager.UpdateFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawFeeToken",
    values: [IFeeManager.WithdrawFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawFeeToken((address,address[],uint256[],bytes))",
    values: [IFeeManager.WithdrawFeeTokenParametersStruct]
  ): string;

  decodeFunctionResult(
    functionFragment: "addFeeToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "addFeeToken((address[],uint256[],bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getCommission",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getCommission(address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "removeFeeToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "removeFeeToken((address[],uint256[],bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateFeeToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateFeeToken((address[],uint256[],bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawFeeToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawFeeToken((address,address[],uint256[],bytes))",
    data: BytesLike
  ): Result;

  events: {
    "AddedFeeToken(address,uint256)": EventFragment;
    "RemovedFeeToken(address,uint256)": EventFragment;
    "UpdatedFeeToken(address,uint256)": EventFragment;
    "WithdrawnFeeToken(address,address,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "AddedFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "AddedFeeToken(address,uint256)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "RemovedFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "RemovedFeeToken(address,uint256)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "UpdatedFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "UpdatedFeeToken(address,uint256)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "WithdrawnFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "WithdrawnFeeToken(address,address,uint256)"
  ): EventFragment;
}

export interface AddedFeeTokenEventObject {
  feeToken: string;
  feeAmount: BigNumber;
}
export type AddedFeeTokenEvent = TypedEvent<
  [string, BigNumber],
  AddedFeeTokenEventObject
>;

export type AddedFeeTokenEventFilter = TypedEventFilter<AddedFeeTokenEvent>;

export interface RemovedFeeTokenEventObject {
  feeToken: string;
  feeAmount: BigNumber;
}
export type RemovedFeeTokenEvent = TypedEvent<
  [string, BigNumber],
  RemovedFeeTokenEventObject
>;

export type RemovedFeeTokenEventFilter = TypedEventFilter<RemovedFeeTokenEvent>;

export interface UpdatedFeeTokenEventObject {
  feeToken: string;
  feeAmount: BigNumber;
}
export type UpdatedFeeTokenEvent = TypedEvent<
  [string, BigNumber],
  UpdatedFeeTokenEventObject
>;

export type UpdatedFeeTokenEventFilter = TypedEventFilter<UpdatedFeeTokenEvent>;

export interface WithdrawnFeeTokenEventObject {
  receiver: string;
  feeToken: string;
  amount: BigNumber;
}
export type WithdrawnFeeTokenEvent = TypedEvent<
  [string, string, BigNumber],
  WithdrawnFeeTokenEventObject
>;

export type WithdrawnFeeTokenEventFilter =
  TypedEventFilter<WithdrawnFeeTokenEvent>;

export interface IFeeManager extends BaseContract {
  contractName: "IFeeManager";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IFeeManagerInterface;

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
    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { commission_: BigNumber }>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { commission_: BigNumber }>;

    removeFeeToken(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "removeFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    updateFeeToken(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "updateFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawFeeToken(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawFeeToken((address,address[],uint256[],bytes))"(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  addFeeToken(
    params_: IFeeManager.AddFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "addFeeToken((address[],uint256[],bytes))"(
    params_: IFeeManager.AddFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  getCommission(
    feeToken_: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  "getCommission(address)"(
    feeToken_: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  removeFeeToken(
    params_: IFeeManager.RemoveFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "removeFeeToken((address[],uint256[],bytes))"(
    params_: IFeeManager.RemoveFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  updateFeeToken(
    params_: IFeeManager.UpdateFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "updateFeeToken((address[],uint256[],bytes))"(
    params_: IFeeManager.UpdateFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawFeeToken(
    params_: IFeeManager.WithdrawFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawFeeToken((address,address[],uint256[],bytes))"(
    params_: IFeeManager.WithdrawFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    removeFeeToken(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "removeFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    updateFeeToken(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "updateFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    withdrawFeeToken(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawFeeToken((address,address[],uint256[],bytes))"(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "AddedFeeToken(address,uint256)"(
      feeToken?: null,
      feeAmount?: null
    ): AddedFeeTokenEventFilter;
    AddedFeeToken(feeToken?: null, feeAmount?: null): AddedFeeTokenEventFilter;

    "RemovedFeeToken(address,uint256)"(
      feeToken?: null,
      feeAmount?: null
    ): RemovedFeeTokenEventFilter;
    RemovedFeeToken(
      feeToken?: null,
      feeAmount?: null
    ): RemovedFeeTokenEventFilter;

    "UpdatedFeeToken(address,uint256)"(
      feeToken?: null,
      feeAmount?: null
    ): UpdatedFeeTokenEventFilter;
    UpdatedFeeToken(
      feeToken?: null,
      feeAmount?: null
    ): UpdatedFeeTokenEventFilter;

    "WithdrawnFeeToken(address,address,uint256)"(
      receiver?: null,
      feeToken?: null,
      amount?: null
    ): WithdrawnFeeTokenEventFilter;
    WithdrawnFeeToken(
      receiver?: null,
      feeToken?: null,
      amount?: null
    ): WithdrawnFeeTokenEventFilter;
  };

  estimateGas: {
    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    removeFeeToken(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "removeFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    updateFeeToken(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "updateFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawFeeToken(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawFeeToken((address,address[],uint256[],bytes))"(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    removeFeeToken(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "removeFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.RemoveFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    updateFeeToken(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "updateFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.UpdateFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawFeeToken(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawFeeToken((address,address[],uint256[],bytes))"(
      params_: IFeeManager.WithdrawFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
