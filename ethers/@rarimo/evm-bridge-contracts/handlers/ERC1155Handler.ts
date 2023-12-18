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

export declare namespace IERC1155Handler {
  export type DepositERC1155ParametersStruct = {
    token: PromiseOrValue<string>;
    tokenId: PromiseOrValue<BigNumberish>;
    amount: PromiseOrValue<BigNumberish>;
    bundle: IBundler.BundleStruct;
    network: PromiseOrValue<string>;
    receiver: PromiseOrValue<string>;
    isWrapped: PromiseOrValue<boolean>;
  };

  export type DepositERC1155ParametersStructOutput = [
    string,
    BigNumber,
    BigNumber,
    IBundler.BundleStructOutput,
    string,
    string,
    boolean
  ] & {
    token: string;
    tokenId: BigNumber;
    amount: BigNumber;
    bundle: IBundler.BundleStructOutput;
    network: string;
    receiver: string;
    isWrapped: boolean;
  };

  export type WithdrawERC1155ParametersStruct = {
    token: PromiseOrValue<string>;
    tokenId: PromiseOrValue<BigNumberish>;
    tokenURI: PromiseOrValue<string>;
    amount: PromiseOrValue<BigNumberish>;
    bundle: IBundler.BundleStruct;
    originHash: PromiseOrValue<BytesLike>;
    receiver: PromiseOrValue<string>;
    proof: PromiseOrValue<BytesLike>;
    isWrapped: PromiseOrValue<boolean>;
  };

  export type WithdrawERC1155ParametersStructOutput = [
    string,
    BigNumber,
    string,
    BigNumber,
    IBundler.BundleStructOutput,
    string,
    string,
    string,
    boolean
  ] & {
    token: string;
    tokenId: BigNumber;
    tokenURI: string;
    amount: BigNumber;
    bundle: IBundler.BundleStructOutput;
    originHash: string;
    receiver: string;
    proof: string;
    isWrapped: boolean;
  };
}

export interface ERC1155HandlerInterface extends utils.Interface {
  functions: {
    "__Bundler_init(address,address)": FunctionFragment;
    "bundleExecutorImplementation()": FunctionFragment;
    "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))": FunctionFragment;
    "determineProxyAddress(bytes32)": FunctionFragment;
    "facade()": FunctionFragment;
    "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)": FunctionFragment;
    "onERC1155Received(address,address,uint256,uint256,bytes)": FunctionFragment;
    "supportsInterface(bytes4)": FunctionFragment;
    "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))": FunctionFragment;
    "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "__Bundler_init"
      | "__Bundler_init(address,address)"
      | "bundleExecutorImplementation"
      | "bundleExecutorImplementation()"
      | "depositERC1155"
      | "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))"
      | "determineProxyAddress"
      | "determineProxyAddress(bytes32)"
      | "facade"
      | "facade()"
      | "onERC1155BatchReceived"
      | "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"
      | "onERC1155Received"
      | "onERC1155Received(address,address,uint256,uint256,bytes)"
      | "supportsInterface"
      | "supportsInterface(bytes4)"
      | "withdrawERC1155"
      | "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"
      | "withdrawERC1155Bundle"
      | "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"
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
    functionFragment: "depositERC1155",
    values: [IERC1155Handler.DepositERC1155ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))",
    values: [IERC1155Handler.DepositERC1155ParametersStruct]
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
    functionFragment: "onERC1155BatchReceived",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>[],
      PromiseOrValue<BigNumberish>[],
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>[],
      PromiseOrValue<BigNumberish>[],
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onERC1155Received",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onERC1155Received(address,address,uint256,uint256,bytes)",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "supportsInterface",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "supportsInterface(bytes4)",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC1155",
    values: [IERC1155Handler.WithdrawERC1155ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))",
    values: [IERC1155Handler.WithdrawERC1155ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC1155Bundle",
    values: [IERC1155Handler.WithdrawERC1155ParametersStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))",
    values: [IERC1155Handler.WithdrawERC1155ParametersStruct]
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
    functionFragment: "depositERC1155",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))",
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
    functionFragment: "onERC1155BatchReceived",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onERC1155Received",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onERC1155Received(address,address,uint256,uint256,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "supportsInterface",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "supportsInterface(bytes4)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC1155",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC1155Bundle",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))",
    data: BytesLike
  ): Result;

  events: {
    "DepositedERC1155(address,uint256,uint256,bytes32,bytes,string,string,bool)": EventFragment;
    "Initialized(uint8)": EventFragment;
    "WithdrawnERC1155(address,uint256,string,uint256,bytes32,bytes,bytes32,address,bytes,bool)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "DepositedERC1155"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "DepositedERC1155(address,uint256,uint256,bytes32,bytes,string,string,bool)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized(uint8)"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "WithdrawnERC1155"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "WithdrawnERC1155(address,uint256,string,uint256,bytes32,bytes,bytes32,address,bytes,bool)"
  ): EventFragment;
}

export interface DepositedERC1155EventObject {
  token: string;
  tokenId: BigNumber;
  amount: BigNumber;
  salt: string;
  bundle: string;
  network: string;
  receiver: string;
  isWrapped: boolean;
}
export type DepositedERC1155Event = TypedEvent<
  [string, BigNumber, BigNumber, string, string, string, string, boolean],
  DepositedERC1155EventObject
>;

export type DepositedERC1155EventFilter =
  TypedEventFilter<DepositedERC1155Event>;

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface WithdrawnERC1155EventObject {
  token: string;
  tokenId: BigNumber;
  tokenURI: string;
  amount: BigNumber;
  salt: string;
  bundle: string;
  originHash: string;
  receiver: string;
  proof: string;
  isWrapped: boolean;
}
export type WithdrawnERC1155Event = TypedEvent<
  [
    string,
    BigNumber,
    string,
    BigNumber,
    string,
    string,
    string,
    string,
    string,
    boolean
  ],
  WithdrawnERC1155EventObject
>;

export type WithdrawnERC1155EventFilter =
  TypedEventFilter<WithdrawnERC1155Event>;

export interface ERC1155Handler extends BaseContract {
  contractName: "ERC1155Handler";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ERC1155HandlerInterface;

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

    depositERC1155(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
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

    onERC1155BatchReceived(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    onERC1155Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "onERC1155Received(address,address,uint256,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    supportsInterface(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    "supportsInterface(bytes4)"(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    withdrawERC1155(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdrawERC1155Bundle(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
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

  depositERC1155(
    params_: IERC1155Handler.DepositERC1155ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))"(
    params_: IERC1155Handler.DepositERC1155ParametersStruct,
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

  onERC1155BatchReceived(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>[],
    arg3: PromiseOrValue<BigNumberish>[],
    arg4: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>[],
    arg3: PromiseOrValue<BigNumberish>[],
    arg4: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  onERC1155Received(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>,
    arg3: PromiseOrValue<BigNumberish>,
    arg4: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "onERC1155Received(address,address,uint256,uint256,bytes)"(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<string>,
    arg2: PromiseOrValue<BigNumberish>,
    arg3: PromiseOrValue<BigNumberish>,
    arg4: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  supportsInterface(
    interfaceId: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  "supportsInterface(bytes4)"(
    interfaceId: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  withdrawERC1155(
    params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
    params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdrawERC1155Bundle(
    params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
    params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
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

    depositERC1155(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
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

    onERC1155BatchReceived(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    onERC1155Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    "onERC1155Received(address,address,uint256,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    supportsInterface(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    "supportsInterface(bytes4)"(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    withdrawERC1155(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    withdrawERC1155Bundle(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "DepositedERC1155(address,uint256,uint256,bytes32,bytes,string,string,bool)"(
      token?: null,
      tokenId?: null,
      amount?: null,
      salt?: null,
      bundle?: null,
      network?: null,
      receiver?: null,
      isWrapped?: null
    ): DepositedERC1155EventFilter;
    DepositedERC1155(
      token?: null,
      tokenId?: null,
      amount?: null,
      salt?: null,
      bundle?: null,
      network?: null,
      receiver?: null,
      isWrapped?: null
    ): DepositedERC1155EventFilter;

    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;

    "WithdrawnERC1155(address,uint256,string,uint256,bytes32,bytes,bytes32,address,bytes,bool)"(
      token?: null,
      tokenId?: null,
      tokenURI?: null,
      amount?: null,
      salt?: null,
      bundle?: null,
      originHash?: null,
      receiver?: null,
      proof?: null,
      isWrapped?: null
    ): WithdrawnERC1155EventFilter;
    WithdrawnERC1155(
      token?: null,
      tokenId?: null,
      tokenURI?: null,
      amount?: null,
      salt?: null,
      bundle?: null,
      originHash?: null,
      receiver?: null,
      proof?: null,
      isWrapped?: null
    ): WithdrawnERC1155EventFilter;
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

    depositERC1155(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
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

    onERC1155BatchReceived(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    onERC1155Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "onERC1155Received(address,address,uint256,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    supportsInterface(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "supportsInterface(bytes4)"(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    withdrawERC1155(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdrawERC1155Bundle(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
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

    depositERC1155(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool))"(
      params_: IERC1155Handler.DepositERC1155ParametersStruct,
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

    onERC1155BatchReceived(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>[],
      arg3: PromiseOrValue<BigNumberish>[],
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    onERC1155Received(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "onERC1155Received(address,address,uint256,uint256,bytes)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<string>,
      arg2: PromiseOrValue<BigNumberish>,
      arg3: PromiseOrValue<BigNumberish>,
      arg4: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    supportsInterface(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "supportsInterface(bytes4)"(
      interfaceId: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    withdrawERC1155(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdrawERC1155Bundle(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool))"(
      params_: IERC1155Handler.WithdrawERC1155ParametersStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
