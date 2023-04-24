const { accounts } = require("./../scripts/utils/utils");

const Reverter = require("./helpers/reverter");
const { SelectorType } = require("./utils/contants");
const {
  getTransferERC20Data,
  getSetWETH9AddressData,
  getUnwrapNativeData,
  getWrapNativeData,
  getExactInputData,
  getExactOutputData,
  getSetUniswapV3RouterAddressData,
  getSetUniswapV2RouterAddressData,
  getUniswapV2SwapETHForExactTokensData,
  getUniswapV2SwapExactETHForTokensData,
  getUniswapV2SwapExactTokensForETHData,
  getUniswapV2SwapExactTokensForTokensData,
  getUniswapV2SwapTokensForExactETHData,
  getUniswapV2SwapTokensForExactTokensData,
  getSetTraderJoeRouterAddressData,
  getTradeJoeSwapAVAXForExactTokensData,
  getTradeJoeSwapExactAVAXForTokensData,
  getTradeJoeSwapExactTokensForAVAXData,
  getTradeJoeSwapExactTokensForTokensData,
  getTradeJoeSwapTokensForExactAVAXData,
  getTradeJoeSwapTokensForExactTokensData,
  getTransferERC721Data,
  getTransferERC1155Data,
  getTransferNativeData,
  getMulticallData,
  getSetBridgeAddressData,
  getBridgeERC20Data,
  getBridgeERC721Data,
  getBridgeERC1155Data,
  getBridgeNativeData,
  getMakeData,
  getOnERC721ReceivedData,
  getOnERC1155ReceivedData,
  getOnERC1155BatchReceivedData,
  getSupportsInterfaceData,
} = require("./utils/swap-diamond-utils");

const SwapDiamond = artifacts.require("SwapDiamond");
const MasterRouter = artifacts.require("MasterRouter");
const BridgeRouter = artifacts.require("BridgeRouter");
const MulticallRouter = artifacts.require("MulticallRouter");
const TraderJoeRouter = artifacts.require("TraderJoeRouter");
const TransferRouter = artifacts.require("TransferRouter");
const UniswapV2Router = artifacts.require("UniswapV2Router");
const UniswapV3Router = artifacts.require("UniswapV3Router");
const WrapRouter = artifacts.require("WrapRouter");

describe("SwapDiamond", async () => {
  const reverter = new Reverter();

  let OWNER;
  let CALLER;

  let diamond;

  before("setup", async () => {
    OWNER = await accounts(0);
    CALLER = await accounts(1);

    diamond = await SwapDiamond.new();

    const masterRouter = await MasterRouter.new();
    const bridgeRouter = await BridgeRouter.new();
    const multicallRouter = await MulticallRouter.new();
    const traderJoeRouter = await TraderJoeRouter.new();
    const transferRouter = await TransferRouter.new();
    const uniswapV2Router = await UniswapV2Router.new();
    const uniswapV3Router = await UniswapV3Router.new();
    const wrapRouter = await WrapRouter.new();

    await diamond.addFacet(
      masterRouter.address,
      [
        "0x00000000", // "receive"
        getMakeData().asSelector(),
        getOnERC721ReceivedData().asSelector(),
        getOnERC1155ReceivedData().asSelector(),
        getOnERC1155BatchReceivedData().asSelector(),
        getSupportsInterfaceData().asSelector(),
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
        SelectorType.SwapDiamond,
      ]
    );

    await diamond.addFacet(
      bridgeRouter.address,
      [
        getSetBridgeAddressData().asSelector(),
        getBridgeERC20Data().asSelector(),
        getBridgeERC721Data().asSelector(),
        getBridgeERC1155Data().asSelector(),
        getBridgeNativeData().asSelector(),
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );

    await diamond.addFacet(multicallRouter.address, getMulticallData().asSelector()[SelectorType.MasterRouter]);

    await diamond.addFacet(
      traderJoeRouter.address,
      [
        getSetTraderJoeRouterAddressData().asSelector(),
        getTradeJoeSwapAVAXForExactTokensData().asSelector(),
        getTradeJoeSwapExactAVAXForTokensData().asSelector(),
        getTradeJoeSwapExactTokensForAVAXData().asSelector(),
        getTradeJoeSwapExactTokensForTokensData().asSelector(),
        getTradeJoeSwapTokensForExactAVAXData().asSelector(),
        getTradeJoeSwapTokensForExactTokensData().asSelector(),
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );

    await diamond.addFacet(
      transferRouter.address,
      [
        getTransferERC20Data().asSelector(),
        getTransferERC721Data().asSelector(),
        getTransferERC1155Data().asSelector(),
        getTransferNativeData().asSelector(),
      ],
      [SelectorType.MasterRouter, SelectorType.MasterRouter, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await diamond.addFacet(
      uniswapV2Router.address,
      [
        getSetUniswapV2RouterAddressData().asSelector(),
        getUniswapV2SwapETHForExactTokensData().asSelector(),
        getUniswapV2SwapExactETHForTokensData().asSelector(),
        getUniswapV2SwapExactTokensForETHData().asSelector(),
        getUniswapV2SwapExactTokensForTokensData().asSelector(),
        getUniswapV2SwapTokensForExactETHData().asSelector(),
        getUniswapV2SwapTokensForExactTokensData().asSelector(),
      ],
      [
        SelectorType.SwapDiamond,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
        SelectorType.MasterRouter,
      ]
    );

    await diamond.addFacet(
      uniswapV3Router.address,
      [
        getExactInputData().asSelector(),
        getExactOutputData().asSelector(),
        getSetUniswapV3RouterAddressData().asSelector(),
      ],
      [SelectorType.MasterRouter, SelectorType.MasterRouter, SelectorType.SwapDiamond]
    );

    await diamond.addFacet(
      wrapRouter.address,
      [getSetWETH9AddressData().asSelector(), getUnwrapNativeData().asSelector(), getWrapNativeData().asSelector()],
      [SelectorType.SwapDiamond, SelectorType.MasterRouter, SelectorType.MasterRouter]
    );

    await reverter.snapshot();
  });

  beforeEach(async () => {});

  afterEach(reverter.revert);
});
