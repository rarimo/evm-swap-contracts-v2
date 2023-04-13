// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "./Constants.sol";
import "../master-facet/MasterRouterStorage.sol";

library Payer {
    using SafeERC20 for IERC20;

    modifier ifNotTransferToThis(address to_) {
        if (to_ != Constants.THIS_ADDRESS) {
            _;
        }
    }

    function pay(address to_, uint256 amount_) internal ifNotTransferToThis(to_) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            amount_ = address(this).balance;
        }

        if (to_ == Constants.CALLER_ADDRESS) {
            to_ = _getCallerAddress();
        }

        (bool ok_, ) = to_.call{value: amount_}("");
        require(ok_, "Payer: failed to transfer native");
    }

    function pay(IERC20 erc20_, address to_, uint256 amount_) internal ifNotTransferToThis(to_) {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            amount_ = erc20_.balanceOf(address(this));
        }

        if (to_ == Constants.CALLER_ADDRESS) {
            to_ = _getCallerAddress();
        }

        erc20_.safeTransfer(to_, amount_);
    }

    function pay(
        IERC721 erc721_,
        address to_,
        uint256[] calldata nftIds_
    ) internal ifNotTransferToThis(to_) {
        if (to_ == Constants.CALLER_ADDRESS) {
            to_ = _getCallerAddress();
        }

        for (uint256 i = 0; i < nftIds_.length; ++i) {
            erc721_.safeTransferFrom(address(this), to_, nftIds_[i], "");
        }
    }

    function pay(
        IERC1155 erc1155_,
        address to_,
        uint256[] calldata ids_,
        uint256[] calldata amounts_
    ) external {
        require(ids_.length == amounts_.length, "Payer: lengths mismatch");

        for (uint256 i = 0; i < ids_.length; ++i) {
            uint256 amount_ = amounts_[i];

            if (amount_ == Constants.CONTRACT_BALANCE) {
                amount_ = erc1155_.balanceOf(address(this), ids_[i]);
            }

            if (to_ == Constants.CALLER_ADDRESS) {
                to_ = _getCallerAddress();
            }

            erc1155_.safeTransferFrom(address(this), to_, ids_[i], amount_, "");
        }
    }

    function _getCallerAddress() private view returns (address) {
        return MasterRouterStorage(address(this)).getCallerAddress();
    }
}
