// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

library Approver {
    function approveMax(IERC20 token_, address to_) internal {
        if (token_.allowance(address(this), to_) == 0) {
            token_.approve(to_, type(uint256).max);
        }
    }

    function approveMax(IERC721 token_, address to_) internal {
        token_.setApprovalForAll(to_, true);
    }

    function approveMax(IERC1155 token_, address to_) internal {
        token_.setApprovalForAll(to_, true);
    }
}
