// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../libs/Resolver.sol";
import "../../libs/ErrorHelper.sol";

contract MulticallRouter {
    using Resolver for *;
    using ErrorHelper for *;

    function multicall(
        address[] calldata targets_,
        bytes[] calldata data_,
        uint256[] calldata values_
    ) external payable {
        require(
            targets_.length == data_.length && data_.length == values_.length,
            "MulticallRouter: lengths mismatch"
        );

        for (uint256 i = 0; i < targets_.length; ++i) {
            (bool ok_, bytes memory data_) = targets_[i].resolve().call{
                value: values_[i].resolve()
            }(data_[i]);

            require(ok_, data_.toStringReason().wrap("MulticallRouter"));
        }
    }
}
