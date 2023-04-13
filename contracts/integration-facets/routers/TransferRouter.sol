// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../../libs/Constants.sol";

contract TransferRouter {
    using SafeERC20 for IERC20;

    function transferERC20(address token_, address to_, uint256 amount_) external {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            amount_ = IERC20(token_).balanceOf(address(this));
        }

        IERC20(token_).safeTransfer(to_, amount_);
    }

    function transferERC721(address token_, address to_, uint256[] calldata nftIds_) external {
        for (uint256 i = 0; i < nftIds_.length; ++i) {
            IERC721(token_).transferFrom(address(this), to_, nftIds_[i]);
        }
    }

    function transferERC1155(
        address token_,
        address to_,
        uint256[] calldata ids_,
        uint256[] calldata amounts_
    ) external {
        require(ids_.length == amounts_.length, "TransferRouter: lengths mismatch");

        for (uint256 i = 0; i < ids_.length; ++i) {
            uint256 amount_ = amounts_[i];

            if (amount_ == Constants.CONTRACT_BALANCE) {
                amount_ = IERC1155(token_).balanceOf(address(this), ids_[i]);
            }

            IERC1155(token_).safeTransferFrom(address(this), to_, ids_[i], amount_, "");
        }
    }

    function transferNative(address to_, uint256 amount_) external {
        if (amount_ == Constants.CONTRACT_BALANCE) {
            amount_ = address(this).balance;
        }

        (bool ok_, ) = to_.call{value: amount_}("");
        require(ok_, "TransferRouter: failed to transfer native");
    }
}
