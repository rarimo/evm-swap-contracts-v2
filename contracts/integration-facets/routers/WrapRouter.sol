// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@solarity/solidity-lib/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "../../libs/Resolver.sol";
import "../../interfaces/tokens/IWrappedNative.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/WrapRouterRouterStorage.sol";
import "./TransferRouter.sol";

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
