// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {IERC1155} from "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

library Approver {
    function approveMax(IERC20 erc20_, address to_) internal {
        if (erc20_.allowance(address(this), to_) == 0) {
            erc20_.approve(to_, type(uint256).max);
        }
    }

    function approveMax(IERC721 erc721_, address to_) internal {
        erc721_.setApprovalForAll(to_, true);
    }

    function approveMax(IERC1155 erc1155_, address to_) internal {
        erc1155_.setApprovalForAll(to_, true);
    }
}
