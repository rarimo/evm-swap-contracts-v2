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
} from "../../../common";

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

export interface FeeManagerInterface extends utils.Interface {
  functions: {
    "__FeeManager_init(address)": FunctionFragment;
    "addFeeToken((address[],uint256[],bytes))": FunctionFragment;
    "bridge()": FunctionFragment;
    "getCommission(address)": FunctionFragment;
    "proxiableUUID()": FunctionFragment;
    "removeFeeToken((address[],uint256[],bytes))": FunctionFragment;
    "updateFeeToken((address[],uint256[],bytes))": FunctionFragment;
    "upgradeTo(address)": FunctionFragment;
    "upgradeToAndCall(address,bytes)": FunctionFragment;
    "upgradeToWithSig(address,bytes)": FunctionFragment;
    "withdrawFeeToken((address,address[],uint256[],bytes))": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "__FeeManager_init"
      | "__FeeManager_init(address)"
      | "addFeeToken"
      | "addFeeToken((address[],uint256[],bytes))"
      | "bridge"
      | "bridge()"
      | "getCommission"
      | "getCommission(address)"
      | "proxiableUUID"
      | "proxiableUUID()"
      | "removeFeeToken"
      | "removeFeeToken((address[],uint256[],bytes))"
      | "updateFeeToken"
      | "updateFeeToken((address[],uint256[],bytes))"
      | "upgradeTo"
      | "upgradeTo(address)"
      | "upgradeToAndCall"
      | "upgradeToAndCall(address,bytes)"
      | "upgradeToWithSig"
      | "upgradeToWithSig(address,bytes)"
      | "withdrawFeeToken"
      | "withdrawFeeToken((address,address[],uint256[],bytes))"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "__FeeManager_init",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "__FeeManager_init(address)",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "addFeeToken",
    values: [IFeeManager.AddFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "addFeeToken((address[],uint256[],bytes))",
    values: [IFeeManager.AddFeeTokenParametersStruct]
  ): string;
  encodeFunctionData(functionFragment: "bridge", values?: undefined): string;
  encodeFunctionData(functionFragment: "bridge()", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "getCommission",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "getCommission(address)",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "proxiableUUID",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "proxiableUUID()",
    values?: undefined
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
    functionFragment: "upgradeTo",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeTo(address)",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall",
    values: [PromiseOrValue<string>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToAndCall(address,bytes)",
    values: [PromiseOrValue<string>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToWithSig",
    values: [PromiseOrValue<string>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "upgradeToWithSig(address,bytes)",
    values: [PromiseOrValue<string>, PromiseOrValue<BytesLike>]
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
    functionFragment: "__FeeManager_init",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "__FeeManager_init(address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "addFeeToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "addFeeToken((address[],uint256[],bytes))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "bridge", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "bridge()", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getCommission",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getCommission(address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "proxiableUUID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "proxiableUUID()",
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
  decodeFunctionResult(functionFragment: "upgradeTo", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "upgradeTo(address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToAndCall(address,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToWithSig",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "upgradeToWithSig(address,bytes)",
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
    "AdminChanged(address,address)": EventFragment;
    "BeaconUpgraded(address)": EventFragment;
    "Initialized(uint8)": EventFragment;
    "RemovedFeeToken(address,uint256)": EventFragment;
    "UpdatedFeeToken(address,uint256)": EventFragment;
    "Upgraded(address)": EventFragment;
    "WithdrawnFeeToken(address,address,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "AddedFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "AddedFeeToken(address,uint256)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "AdminChanged"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "AdminChanged(address,address)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "BeaconUpgraded"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "BeaconUpgraded(address)"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized(uint8)"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "RemovedFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "RemovedFeeToken(address,uint256)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "UpdatedFeeToken"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "UpdatedFeeToken(address,uint256)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Upgraded"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Upgraded(address)"): EventFragment;
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

export interface AdminChangedEventObject {
  previousAdmin: string;
  newAdmin: string;
}
export type AdminChangedEvent = TypedEvent<
  [string, string],
  AdminChangedEventObject
>;

export type AdminChangedEventFilter = TypedEventFilter<AdminChangedEvent>;

export interface BeaconUpgradedEventObject {
  beacon: string;
}
export type BeaconUpgradedEvent = TypedEvent<
  [string],
  BeaconUpgradedEventObject
>;

export type BeaconUpgradedEventFilter = TypedEventFilter<BeaconUpgradedEvent>;

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

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

export interface UpgradedEventObject {
  implementation: string;
}
export type UpgradedEvent = TypedEvent<[string], UpgradedEventObject>;

export type UpgradedEventFilter = TypedEventFilter<UpgradedEvent>;

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

export interface FeeManager extends BaseContract {
  contractName: "FeeManager";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: FeeManagerInterface;

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
    __FeeManager_init(
      bridge_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "__FeeManager_init(address)"(
      bridge_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    bridge(overrides?: CallOverrides): Promise<[string]>;

    "bridge()"(overrides?: CallOverrides): Promise<[string]>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { commission_: BigNumber }>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { commission_: BigNumber }>;

    proxiableUUID(overrides?: CallOverrides): Promise<[string]>;

    "proxiableUUID()"(overrides?: CallOverrides): Promise<[string]>;

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

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "upgradeTo(address)"(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "upgradeToAndCall(address,bytes)"(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    upgradeToWithSig(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "upgradeToWithSig(address,bytes)"(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
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

  __FeeManager_init(
    bridge_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "__FeeManager_init(address)"(
    bridge_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  addFeeToken(
    params_: IFeeManager.AddFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "addFeeToken((address[],uint256[],bytes))"(
    params_: IFeeManager.AddFeeTokenParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  bridge(overrides?: CallOverrides): Promise<string>;

  "bridge()"(overrides?: CallOverrides): Promise<string>;

  getCommission(
    feeToken_: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  "getCommission(address)"(
    feeToken_: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  proxiableUUID(overrides?: CallOverrides): Promise<string>;

  "proxiableUUID()"(overrides?: CallOverrides): Promise<string>;

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

  upgradeTo(
    newImplementation: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "upgradeTo(address)"(
    newImplementation: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  upgradeToAndCall(
    newImplementation: PromiseOrValue<string>,
    data: PromiseOrValue<BytesLike>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "upgradeToAndCall(address,bytes)"(
    newImplementation: PromiseOrValue<string>,
    data: PromiseOrValue<BytesLike>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  upgradeToWithSig(
    newImplementation_: PromiseOrValue<string>,
    signature_: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "upgradeToWithSig(address,bytes)"(
    newImplementation_: PromiseOrValue<string>,
    signature_: PromiseOrValue<BytesLike>,
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
    __FeeManager_init(
      bridge_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "__FeeManager_init(address)"(
      bridge_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    bridge(overrides?: CallOverrides): Promise<string>;

    "bridge()"(overrides?: CallOverrides): Promise<string>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    proxiableUUID(overrides?: CallOverrides): Promise<string>;

    "proxiableUUID()"(overrides?: CallOverrides): Promise<string>;

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

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "upgradeTo(address)"(
      newImplementation: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "upgradeToAndCall(address,bytes)"(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    upgradeToWithSig(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "upgradeToWithSig(address,bytes)"(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
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

    "AdminChanged(address,address)"(
      previousAdmin?: null,
      newAdmin?: null
    ): AdminChangedEventFilter;
    AdminChanged(
      previousAdmin?: null,
      newAdmin?: null
    ): AdminChangedEventFilter;

    "BeaconUpgraded(address)"(
      beacon?: PromiseOrValue<string> | null
    ): BeaconUpgradedEventFilter;
    BeaconUpgraded(
      beacon?: PromiseOrValue<string> | null
    ): BeaconUpgradedEventFilter;

    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;

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

    "Upgraded(address)"(
      implementation?: PromiseOrValue<string> | null
    ): UpgradedEventFilter;
    Upgraded(
      implementation?: PromiseOrValue<string> | null
    ): UpgradedEventFilter;

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
    __FeeManager_init(
      bridge_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "__FeeManager_init(address)"(
      bridge_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    bridge(overrides?: CallOverrides): Promise<BigNumber>;

    "bridge()"(overrides?: CallOverrides): Promise<BigNumber>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    proxiableUUID(overrides?: CallOverrides): Promise<BigNumber>;

    "proxiableUUID()"(overrides?: CallOverrides): Promise<BigNumber>;

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

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "upgradeTo(address)"(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "upgradeToAndCall(address,bytes)"(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    upgradeToWithSig(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "upgradeToWithSig(address,bytes)"(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
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
    __FeeManager_init(
      bridge_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "__FeeManager_init(address)"(
      bridge_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    addFeeToken(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "addFeeToken((address[],uint256[],bytes))"(
      params_: IFeeManager.AddFeeTokenParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    bridge(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "bridge()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getCommission(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "getCommission(address)"(
      feeToken_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    proxiableUUID(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "proxiableUUID()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

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

    upgradeTo(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "upgradeTo(address)"(
      newImplementation: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    upgradeToAndCall(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "upgradeToAndCall(address,bytes)"(
      newImplementation: PromiseOrValue<string>,
      data: PromiseOrValue<BytesLike>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    upgradeToWithSig(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "upgradeToWithSig(address,bytes)"(
      newImplementation_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
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
