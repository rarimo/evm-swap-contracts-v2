// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract MulticallRouter {
    function multicall(
        address[] calldata targets_,
        bytes[] calldata data_,
        uint256[] calldata values_
    ) external {
        for (uint256 i = 0; i < targets_.length; ++i) {
            (bool ok, ) = targets_[i].call{value: values_[i]}(data_[i]);

            require(ok, "MulticallRouter: failed to call");
        }
    }
}
