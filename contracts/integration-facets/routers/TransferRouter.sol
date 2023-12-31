// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {IERC1155} from "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {MasterRouterStorage} from "../../master-facet/MasterRouterStorage.sol";
import {Resolver} from "../../libs/Resolver.sol";

contract TransferRouter is MasterRouterStorage {
    using SafeERC20 for IERC20;
    using Resolver for *;

    function transferERC20(address token_, address receiver_, uint256 amount_) public payable {
        receiver_ = receiver_.resolve();

        if (receiver_ == address(this)) {
            return;
        }

        IERC20(token_).safeTransfer(receiver_, amount_.resolve(IERC20(token_)));
    }

    function transferERC721(
        address token_,
        address receiver_,
        uint256[] calldata nftIds_
    ) external payable {
        receiver_ = receiver_.resolve();

        if (receiver_ == address(this)) {
            return;
        }

        for (uint256 i = 0; i < nftIds_.length; ++i) {
            IERC721(token_).safeTransferFrom(address(this), receiver_, nftIds_[i], "");
        }
    }

    function transferERC1155(
        address token_,
        address receiver_,
        uint256[] calldata tokenIds_,
        uint256[] calldata amounts_
    ) external payable {
        require(tokenIds_.length == amounts_.length, "TransferRouter: lengths mismatch");

        receiver_ = receiver_.resolve();

        if (receiver_ == address(this)) {
            return;
        }

        for (uint256 i = 0; i < tokenIds_.length; ++i) {
            uint256 tokenId_ = tokenIds_[i];

            IERC1155(token_).safeTransferFrom(
                address(this),
                receiver_,
                tokenId_,
                amounts_[i].resolve(IERC1155(token_), tokenId_),
                ""
            );
        }
    }

    function transferFromERC20(address token_, uint256 amount_) external payable {
        IERC20(token_).safeTransferFrom(getCallerAddress(), address(this), amount_);
    }

    function transferFromERC721(address token_, uint256[] calldata nftIds_) external payable {
        for (uint256 i = 0; i < nftIds_.length; i++) {
            IERC721(token_).safeTransferFrom(getCallerAddress(), address(this), nftIds_[i], "");
        }
    }

    function transferFromERC1155(
        address token_,
        uint256[] calldata tokenIds_,
        uint256[] calldata amounts_
    ) external payable {
        IERC1155(token_).safeBatchTransferFrom(
            getCallerAddress(),
            address(this),
            tokenIds_,
            amounts_,
            ""
        );
    }

    function transferNative(address receiver_, uint256 amount_) public payable {
        receiver_ = receiver_.resolve();

        if (receiver_ == address(this)) {
            return;
        }

        (bool ok_, ) = receiver_.call{value: amount_.resolve()}("");
        require(ok_, "TransferRouter: failed to transfer native");
    }
}
