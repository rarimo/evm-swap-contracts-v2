// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "../../libs/Resolver.sol";
import "../../interfaces/tokens/IWETH9.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/WrapRouterRouterStorage.sol";
import "./TransferRouter.sol";

contract WrapRouter is OwnableDiamondStorage, WrapRouterStorage, TransferRouter {
    using Resolver for uint256;

    function setWETH9Address(address weth9_) external onlyOwner {
        _getWrapRouterStorage().weth9 = weth9_;
    }

    function wrap(address receiver_, uint256 amount_) external {
        amount_ = amount_.resolve();

        address weth9_ = getWETH9Address();

        IWETH9(weth9_).deposit{value: amount_}();
        transferERC20(weth9_, receiver_, amount_);
    }

    function unwrap(address receiver_, uint256 amount_) external {
        address weth9_ = getWETH9Address();

        amount_ = amount_.resolve(IERC20(weth9_));

        IWETH9(weth9_).withdraw(amount_);
        transferNative(receiver_, amount_);
    }
}
