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
} from "../../../common";

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

export declare namespace IERC721Handler {
  export type DepositERC721ParametersStruct = {
    token: PromiseOrValue<string>;
    tokenId: PromiseOrValue<BigNumberish>;
    bundle: IBundler.BundleStruct;
    network: PromiseOrValue<string>;
    receiver: PromiseOrValue<string>;
    isWrapped: PromiseOrValue<boolean>;
  };

  export type DepositERC721ParametersStructOutput = [
    string,
    BigNumber,
    IBundler.BundleStructOutput,
    string,
    string,
    boolean
  ] & {
    token: string;
    tokenId: BigNumber;
    bundle: IBundler.BundleStructOutput;
    network: string;
    receiver: string;
    isWrapped: boolean;
  };

  export type WithdrawERC721ParametersStruct = {
    token: PromiseOrValue<string>;
    tokenId: PromiseOrValue<BigNumberish>;
    tokenURI: PromiseOrValue<string>;
    bundle: IBundler.BundleStruct;
    originHash: PromiseOrValue<BytesLike>;
    receiver: PromiseOrValue<string>;
    proof: PromiseOrValue<BytesLike>;
    isWrapped: PromiseOrValue<boolean>;
  };

  export type WithdrawERC721ParametersStructOutput = [
    string,
    BigNumber,
    string,
    IBundler.BundleStructOutput,
    string,
    string,
    string,
    boolean
  ] & {
    token: string;
    tokenId: BigNumber;
    tokenURI: string;
    bundle: IBundler.BundleStructOutput;
    originHash: string;
    receiver: string;
    proof: string;
    isWrapped: boolean;
  };
}

export interface ERC721HandlerInterface extends utils.Interface {
  functions: {
    "__Bundler_init(address,address)": FunctionFragment;
    "bundleExecutorImplementation()": FunctionFragment;
    "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))": FunctionFragment;
    "determineProxyAddress(bytes32)": FunctionFragment;
    "facade()": FunctionFragment;
    "onERC721Received(address,address,uint256,bytes)": FunctionFragment;
    "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))": FunctionFragment;
    "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "__Bundler_init"
      | "__Bundler_init(address,address)"
      | "bundleExecutorImplementation"
      | "bundleExecutorImplementation()"
      | "depositERC721"
      | "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))"
      | "determineProxyAddress"
      | "determineProxyAddress(bytes32)"
      | "facade"
      | "facade()"
      | "onERC721Received"
      | "onERC721Received(address,address,uint256,bytes)"
      | "withdrawERC721"
      | "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"
      | "withdrawERC721Bundle"
      | "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "__Bundler_init",
    values: [PromiseOrValue<string>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "__Bundler_init(address,address)",
    values: [PromiseOrValue<string>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "bundleExecutorImplementation",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "bundleExecutorImplementation()",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "depositERC721",
    values: [IERC721Handler.DepositERC721ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))",
    values: [IERC721Handler.DepositERC721ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "determineProxyAddress",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "determineProxyAddress(bytes32)",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(functionFragment: "facade", values?: undefined): string;
  encodeFunctionData(functionFragment: "facade()", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "onERC721Received",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onERC721Received(address,address,uint256,bytes)",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC721",
    values: [IERC721Handler.WithdrawERC721ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))",
    values: [IERC721Handler.WithdrawERC721ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC721Bundle",
    values: [IERC721Handler.WithdrawERC721ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))",
    values: [IERC721Handler.WithdrawERC721ParametersStruct]
  ): string;

  decodeFunctionResult(
    functionFragment: "__Bundler_init",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "__Bundler_init(address,address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "bundleExecutorImplementation",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "bundleExecutorImplementation()",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositERC721",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))",
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
  decodeFunctionResult(functionFragment: "facade", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "facade()", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "onERC721Received",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onERC721Received(address,address,uint256,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC721",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC721Bundle",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))",
    data: BytesLike
  ): Result;

  events: {
    "DepositedERC721(address,uint256,bytes32,bytes,string,string,bool)": EventFragment;
    "Initialized(uint8)": EventFragment;
    "WithdrawnERC721(address,uint256,string,bytes32,bytes,bytes32,address,bytes,bool)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "DepositedERC721"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "DepositedERC721(address,uint256,bytes32,bytes,string,string,bool)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized(uint8)"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "WithdrawnERC721"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "WithdrawnERC721(address,uint256,string,bytes32,bytes,bytes32,address,bytes,bool)"
  ): EventFragment;
}

export interface DepositedERC721EventObject {
  token: string;
  tokenId: BigNumber;
  salt: string;
  bundle: string;
  network: string;
  receiver: string;
  isWrapped: boolean;
}
export type DepositedERC721Event = TypedEvent<
  [string, BigNumber, string, string, string, string, boolean],
  DepositedERC721EventObject
>;

export type DepositedERC721EventFilter = TypedEventFilter<DepositedERC721Event>;

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface WithdrawnERC721EventObject {
  token: string;
  tokenId: BigNumber;
  tokenURI: string;
  salt: string;
  bundle: string;
  originHash: string;
  receiver: string;
  proof: string;
  isWrapped: boolean;
}
export type WithdrawnERC721Event = TypedEvent<
  [string, BigNumber, string, string, string, string, string, string, boolean],
  WithdrawnERC721EventObject
>;

export type WithdrawnERC721EventFilter = TypedEventFilter<WithdrawnERC721Event>;

export interface ERC721Handler extends BaseContract {
  contractName: "ERC721Handler";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ERC721HandlerInterface;

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
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "__Bundler_init(address,address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    bundleExecutorImplementation(overrides?: CallOverrides): Promise<[string]>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<[string]>;

    depositERC721(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    facade(overrides?: CallOverrides): Promise<[string]>;

    "facade()"(overrides?: CallOverrides): Promise<[string]>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "onERC721Received(address,address,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawERC721(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawERC721Bundle(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  __Bundler_init(
    bundleExecutorImplementation_: PromiseOrValue<string>,
    facade_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "__Bundler_init(address,address)"(
    bundleExecutorImplementation_: PromiseOrValue<string>,
    facade_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  bundleExecutorImplementation(overrides?: CallOverrides): Promise<string>;

  "bundleExecutorImplementation()"(overrides?: CallOverrides): Promise<string>;

  depositERC721(
    params_: IERC721Handler.DepositERC721ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))"(
    params_: IERC721Handler.DepositERC721ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  determineProxyAddress(
    salt_: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  "determineProxyAddress(bytes32)"(
    salt_: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  facade(overrides?: CallOverrides): Promise<string>;

  "facade()"(overrides?: CallOverrides): Promise<string>;

  onERC721Received(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>,
    arg3: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "onERC721Received(address,address,uint256,bytes)"(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>,
    arg3: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawERC721(
    params_: IERC721Handler.WithdrawERC721ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
    params_: IERC721Handler.WithdrawERC721ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawERC721Bundle(
    params_: IERC721Handler.WithdrawERC721ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
    params_: IERC721Handler.WithdrawERC721ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "__Bundler_init(address,address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    bundleExecutorImplementation(overrides?: CallOverrides): Promise<string>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<string>;

    depositERC721(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC721Handler.DepositERC721ParametersStruct,
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

    facade(overrides?: CallOverrides): Promise<string>;

    "facade()"(overrides?: CallOverrides): Promise<string>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    "onERC721Received(address,address,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    withdrawERC721(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    withdrawERC721Bundle(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "DepositedERC721(address,uint256,bytes32,bytes,string,string,bool)"(
      token?: null,
      tokenId?: null,
      salt?: null,
      bundle?: null,
      network?: null,
      receiver?: null,
      isWrapped?: null
    ): DepositedERC721EventFilter;
    DepositedERC721(
      token?: null,
      tokenId?: null,
      salt?: null,
      bundle?: null,
      network?: null,
      receiver?: null,
      isWrapped?: null
    ): DepositedERC721EventFilter;

    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;

    "WithdrawnERC721(address,uint256,string,bytes32,bytes,bytes32,address,bytes,bool)"(
      token?: null,
      tokenId?: null,
      tokenURI?: null,
      salt?: null,
      bundle?: null,
      originHash?: null,
      receiver?: null,
      proof?: null,
      isWrapped?: null
    ): WithdrawnERC721EventFilter;
    WithdrawnERC721(
      token?: null,
      tokenId?: null,
      tokenURI?: null,
      salt?: null,
      bundle?: null,
      originHash?: null,
      receiver?: null,
      proof?: null,
      isWrapped?: null
    ): WithdrawnERC721EventFilter;
  };

  estimateGas: {
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "__Bundler_init(address,address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    bundleExecutorImplementation(overrides?: CallOverrides): Promise<BigNumber>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    depositERC721(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    facade(overrides?: CallOverrides): Promise<BigNumber>;

    "facade()"(overrides?: CallOverrides): Promise<BigNumber>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "onERC721Received(address,address,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawERC721(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawERC721Bundle(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "__Bundler_init(address,address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    bundleExecutorImplementation(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    depositERC721(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "depositERC721((address,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC721Handler.DepositERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    facade(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "facade()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    onERC721Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "onERC721Received(address,address,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawERC721(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawERC721Bundle(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC721Handler.WithdrawERC721ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
