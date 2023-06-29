// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

import "./Constants.sol";
import "../master-facet/MasterRouterStorage.sol";

library Resolver {
    function resolve(address address_) internal view returns (address) {
        if (address_ == Constants.THIS_ADDRESS) {
            return address(this);
        }

        if (address_ == Constants.CALLER_ADDRESS) {
            return MasterRouterStorage(address(this)).getCallerAddress();
        }

        return address_;
    }

    function resolve(uint256 amount_) internal view returns (uint256) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            return address(this).balance;
        }

        return amount_;
    }

    function resolve(uint256 amount_, uint256 commission_) internal view returns (uint256) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            require(address(this).balance >= commission_, "Resolver: commission exceeds balance");

            return address(this).balance - commission_;
        }

        return amount_;
    }

    function resolve(uint256 amount_, IERC20 erc20_) internal view returns (uint256) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            return erc20_.balanceOf(address(this));
        }

        return amount_;
    }

    function resolve(
        uint256 amount_,
        IERC20 erc20_,
        uint256 commission_
    ) internal view returns (uint256) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            uint256 erc20Balance_ = erc20_.balanceOf(address(this));

            require(erc20Balance_ >= commission_, "Resolver: commission exceeds balance");

            return erc20Balance_ - commission_;
        }

        return amount_;
    }

    function resolve(
        uint256 amount_,
        IERC1155 erc1155_,
        uint256 tokenId_
    ) internal view returns (uint256) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            return erc1155_.balanceOf(address(this), tokenId_);
        }

        return amount_;
    }
}
