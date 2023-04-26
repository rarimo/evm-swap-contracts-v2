// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../libs/Resolver.sol";

contract MulticallRouter {
    using Resolver for *;

    function multicall(
        address[] calldata targets_,
        bytes[] calldata data_,
        uint256[] calldata values_
    ) external payable {
        for (uint256 i = 0; i < targets_.length; ++i) {
            (bool ok_, ) = targets_[i].resolve().call{value: values_[i].resolve()}(data_[i]);

            require(ok_, "MulticallRouter: failed to call");
        }
    }
}
