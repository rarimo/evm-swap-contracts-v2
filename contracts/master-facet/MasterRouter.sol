// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC721Holder} from "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import {ERC1155Holder} from "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

import {DiamondStorage} from "@solarity/solidity-lib/diamond/DiamondStorage.sol";

import {Commands} from "../libs/Commands.sol";
import {ErrorHelper} from "../libs/ErrorHelper.sol";
import {MasterRouterStorage} from "./MasterRouterStorage.sol";
import {SwapDiamondStorage} from "../SwapDiamondStorage.sol";
import {BridgeRouter} from "../integration-facets/routers/BridgeRouter.sol";
import {TransferRouter} from "../integration-facets/routers/TransferRouter.sol";
import {WrapRouter} from "../integration-facets/routers/WrapRouter.sol";
import {MulticallRouter} from "../integration-facets/routers/MulticallRouter.sol";
import {UniswapV2Router} from "../integration-facets/routers/UniswapV2Router.sol";
import {UniswapV3Router} from "../integration-facets/routers/UniswapV3Router.sol";
import {TraderJoeRouter} from "../integration-facets/routers/TraderJoeRouter.sol";

contract MasterRouter is
    DiamondStorage,
    MasterRouterStorage,
    SwapDiamondStorage,
    ERC721Holder,
    ERC1155Holder
{
    using ErrorHelper for *;

    struct Payload {
        uint256 command;
        bool skipRevert;
        bytes data;
    }

    function make(Payload[] calldata payloads_) external payable onlyCaller {
        for (uint256 i = 0; i < payloads_.length; ++i) {
            _handle(payloads_[i]);
        }
    }

    function _handle(Payload calldata payload_) internal {
        bytes4 funcSelector_ = _getSelector(payload_.command);

        require(
            getSelectorType(funcSelector_) == SelectorType.MasterRouter,
            "MasterRouter: invalid command"
        );

        (bool ok_, bytes memory data_) = getFacetBySelector(funcSelector_).delegatecall(
            abi.encodePacked(funcSelector_, payload_.data)
        );

        require(ok_ || payload_.skipRevert, data_.toStringReason().wrap("MasterRouter"));
    }

    function _getSelector(uint256 command_) internal pure returns (bytes4 funcSelector_) {
        if (command_ < 10) {
            funcSelector_ = _getBridgeSelector(command_);
        } else if (command_ < 20) {
            funcSelector_ = _getTransferSelector(command_);
        } else if (command_ < 25) {
            funcSelector_ = _getWrapSelector(command_);
        } else if (command_ < 30) {
            funcSelector_ = _getMulticallSelector(command_);
        } else if (command_ < 60) {
            funcSelector_ = _getUniswapV2Selector(command_);
        } else if (command_ < 70) {
            funcSelector_ = _getUniswapV3Selector(command_);
        } else if (command_ < 80) {
            funcSelector_ = _getTraderJoeSelector(command_);
        }
    }

    function _getBridgeSelector(uint256 command_) internal pure returns (bytes4 bridgeSelector_) {
        if (command_ == Commands.BRIDGE_ERC20) {
            bridgeSelector_ = BridgeRouter.bridgeERC20.selector;
        } else if (command_ == Commands.BRIDGE_ERC721) {
            bridgeSelector_ = BridgeRouter.bridgeERC721.selector;
        } else if (command_ == Commands.BRIDGE_ERC1155) {
            bridgeSelector_ = BridgeRouter.bridgeERC1155.selector;
        } else if (command_ == Commands.BRIDGE_NATIVE) {
            bridgeSelector_ = BridgeRouter.bridgeNative.selector;
        }
    }

    function _getTransferSelector(
        uint256 command_
    ) internal pure returns (bytes4 transferSelector_) {
        if (command_ == Commands.TRANSFER_ERC20) {
            transferSelector_ = TransferRouter.transferERC20.selector;
        } else if (command_ == Commands.TRANSFER_ERC721) {
            transferSelector_ = TransferRouter.transferERC721.selector;
        } else if (command_ == Commands.TRANSFER_ERC1155) {
            transferSelector_ = TransferRouter.transferERC1155.selector;
        } else if (command_ == Commands.TRANSFER_NATIVE) {
            transferSelector_ = TransferRouter.transferNative.selector;
        } else if (command_ == Commands.TRANSFER_FROM_ERC20) {
            transferSelector_ = TransferRouter.transferFromERC20.selector;
        } else if (command_ == Commands.TRANSFER_FROM_ERC721) {
            transferSelector_ = TransferRouter.transferFromERC721.selector;
        } else if (command_ == Commands.TRANSFER_FROM_ERC1155) {
            transferSelector_ = TransferRouter.transferFromERC1155.selector;
        }
    }

    function _getWrapSelector(uint256 command_) internal pure returns (bytes4 wrapSelector_) {
        if (command_ == Commands.WRAP_NATIVE) {
            wrapSelector_ = WrapRouter.wrap.selector;
        } else if (command_ == Commands.UNWRAP_NATIVE) {
            wrapSelector_ = WrapRouter.unwrap.selector;
        }
    }

    function _getMulticallSelector(
        uint256 command_
    ) internal pure returns (bytes4 multicallSelector_) {
        if (command_ == Commands.MULTICALL) {
            multicallSelector_ = MulticallRouter.multicall.selector;
        }
    }

    function _getUniswapV2Selector(
        uint256 command_
    ) internal pure returns (bytes4 uniswapV2Selector_) {
        if (command_ == Commands.SWAP_EXACT_TOKENS_FOR_TOKENS_V2) {
            uniswapV2Selector_ = UniswapV2Router.swapExactTokensForTokensV2.selector;
        } else if (command_ == Commands.SWAP_TOKENS_FOR_EXACT_TOKENS_V2) {
            uniswapV2Selector_ = UniswapV2Router.swapTokensForExactTokensV2.selector;
        } else if (command_ == Commands.SWAP_EXACT_ETH_FOR_TOKENS) {
            uniswapV2Selector_ = UniswapV2Router.swapExactETHForTokens.selector;
        } else if (command_ == Commands.SWAP_TOKENS_FOR_EXACT_ETH) {
            uniswapV2Selector_ = UniswapV2Router.swapTokensForExactETH.selector;
        } else if (command_ == Commands.SWAP_EXACT_TOKENS_FOR_ETH) {
            uniswapV2Selector_ = UniswapV2Router.swapExactTokensForETH.selector;
        } else if (command_ == Commands.SWAP_ETH_FOR_EXACT_TOKENS) {
            uniswapV2Selector_ = UniswapV2Router.swapETHForExactTokens.selector;
        }
    }

    function _getUniswapV3Selector(
        uint256 command_
    ) internal pure returns (bytes4 uniswapV3Selector_) {
        if (command_ == Commands.EXACT_INPUT) {
            uniswapV3Selector_ = UniswapV3Router.exactInput.selector;
        } else if (command_ == Commands.EXACT_OUTPUT) {
            uniswapV3Selector_ = UniswapV3Router.exactOutput.selector;
        }
    }

    function _getTraderJoeSelector(
        uint256 command_
    ) internal pure returns (bytes4 traderJoeSelector_) {
        if (command_ == Commands.SWAP_EXACT_TOKENS_FOR_TOKENS_TJ) {
            traderJoeSelector_ = TraderJoeRouter.swapExactTokensForTokensTJ.selector;
        } else if (command_ == Commands.SWAP_TOKENS_FOR_EXACT_TOKENS_TJ) {
            traderJoeSelector_ = TraderJoeRouter.swapTokensForExactTokensTJ.selector;
        } else if (command_ == Commands.SWAP_EXACT_AVAX_FOR_TOKENS) {
            traderJoeSelector_ = TraderJoeRouter.swapExactAVAXForTokens.selector;
        } else if (command_ == Commands.SWAP_TOKENS_FOR_EXACT_AVAX) {
            traderJoeSelector_ = TraderJoeRouter.swapTokensForExactAVAX.selector;
        } else if (command_ == Commands.SWAP_EXACT_TOKENS_FOR_AVAX) {
            traderJoeSelector_ = TraderJoeRouter.swapExactTokensForAVAX.selector;
        } else if (command_ == Commands.SWAP_AVAX_FOR_EXACT_TOKENS) {
            traderJoeSelector_ = TraderJoeRouter.swapAVAXForExactTokens.selector;
        }
    }
}
