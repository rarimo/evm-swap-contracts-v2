// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "./Resolver.sol";

library Payer {
    using Resolver for *;
    using SafeERC20 for IERC20;

    function pay(address to_, uint256 amount_) internal {
        to_ = to_.resolve();

        if (to_ == address(this)) {
            return;
        }

        (bool ok_, ) = to_.call{value: amount_.resolve()}("");
        require(ok_, "Payer: failed to transfer native");
    }

    function pay(IERC20 erc20_, address to_, uint256 amount_) internal {
        to_ = to_.resolve();

        if (to_ == address(this)) {
            return;
        }

        erc20_.safeTransfer(to_, amount_.resolve(erc20_));
    }

    function pay(IERC721 erc721_, address to_, uint256[] calldata nftIds_) internal {
        to_ = to_.resolve();

        if (to_ == address(this)) {
            return;
        }

        for (uint256 i = 0; i < nftIds_.length; ++i) {
            erc721_.safeTransferFrom(address(this), to_, nftIds_[i], "");
        }
    }

    function pay(
        IERC1155 erc1155_,
        address to_,
        uint256[] calldata tokenIds_,
        uint256[] calldata amounts_
    ) internal {
        require(tokenIds_.length == amounts_.length, "Payer: lengths mismatch");

        to_ = to_.resolve();

        if (to_ == address(this)) {
            return;
        }

        for (uint256 i = 0; i < tokenIds_.length; ++i) {
            uint256 tokenId_ = tokenIds_[i];

            erc1155_.safeTransferFrom(
                address(this),
                to_,
                tokenId_,
                amounts_[i].resolve(erc1155_, tokenId_),
                ""
            );
        }
    }
}
