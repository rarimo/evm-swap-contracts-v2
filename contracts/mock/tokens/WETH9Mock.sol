// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../interfaces/tokens/IWETH9.sol";
import "./ERC20Mock.sol";

contract WETH9Mock is IWETH9, ERC20Mock {
    constructor() ERC20Mock("WETH9Mock", "WETH9M", 18) {}

    function deposit() external payable override {
        _mint(msg.sender, msg.value);
    }

    function withdraw(uint256 amount_) external override {
        _burn(msg.sender, amount_);

        payable(msg.sender).transfer(amount_);
    }
}
