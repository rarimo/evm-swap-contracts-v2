// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {OwnableDiamondStorage} from "@solarity/solidity-lib/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import {Resolver} from "../../libs/Resolver.sol";
import {IWrappedNative} from "../../interfaces/tokens/IWrappedNative.sol";
import {MasterRouterStorage} from "../../master-facet/MasterRouterStorage.sol";
import {WrapRouterStorage} from "../storages/WrapRouterStorage.sol";
import {TransferRouter} from "./TransferRouter.sol";

contract WrapRouter is OwnableDiamondStorage, WrapRouterStorage, TransferRouter {
    using Resolver for uint256;

    function setWrappedNativeAddress(address wrappedNative_) external onlyOwner {
        _getWrapRouterStorage().wrappedNative = wrappedNative_;
    }

    function wrap(address receiver_, uint256 amount_) external payable {
        address weth9_ = getWrappedNativeAddress();

        amount_ = amount_.resolve();

        IWrappedNative(weth9_).deposit{value: amount_}();
        transferERC20(weth9_, receiver_, amount_);
    }

    function unwrap(address receiver_, uint256 amount_) external payable {
        address weth9_ = getWrappedNativeAddress();

        amount_ = amount_.resolve(IERC20(weth9_));

        IWrappedNative(weth9_).withdraw(amount_);
        transferNative(receiver_, amount_);
    }
}
