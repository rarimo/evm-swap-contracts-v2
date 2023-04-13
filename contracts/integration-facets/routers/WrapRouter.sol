// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@dlsl/dev-modules/diamond/presets/OwnableDiamond/OwnableDiamondStorage.sol";

import "../../libs/Constants.sol";
import "../../interfaces/tokens/IWETH9.sol";
import "../../master-facet/MasterRouterStorage.sol";
import "../storages/WrapRouterRouterStorage.sol";

contract WrapRouter is OwnableDiamondStorage, WrapRouterStorage {
    function setWETH9Address(address weth9_) external onlyOwner {
        getWrapRouterStorage().weth9 = weth9_;
    }

    function wrap(uint256 amount_) external {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            amount_ = address(this).balance;
        } else {
            require(amount_ <= address(this).balance, "WrapRouter: insufficient native balance");
        }

        IWETH9(getWETH9Address()).deposit{value: amount_}();
    }

    function unwrap(uint256 amount_) external {
        address weth9_ = getWETH9Address();

        if (amount_ == Constants.CONTRACT_BALANCE) {
            amount_ = IWETH9(weth9_).balanceOf(address(this));
        } else {
            require(
                IWETH9(weth9_).balanceOf(address(this)) >= amount_,
                "WrapRouter: insufficient WETH balance"
            );
        }

        IWETH9(weth9_).withdraw(amount_);
    }
}
