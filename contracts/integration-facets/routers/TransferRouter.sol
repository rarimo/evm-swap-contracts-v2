// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

import "../../libs/Payer.sol";

contract TransferRouter {
    using Payer for *;

    function transferERC20(address token_, address receiver_, uint256 amount_) public {
        IERC20(token_).pay(receiver_, amount_);
    }

    function transferERC721(
        address token_,
        address receiver_,
        uint256[] calldata nftIds_
    ) external {
        IERC721(token_).pay(receiver_, nftIds_);
    }

    function transferERC1155(
        address token_,
        address receiver_,
        uint256[] calldata tokenIds_,
        uint256[] calldata amounts_
    ) external {
        IERC1155(token_).pay(receiver_, tokenIds_, amounts_);
    }

    function transferNative(address receiver_, uint256 amount_) public {
        receiver_.pay(amount_);
    }
}
