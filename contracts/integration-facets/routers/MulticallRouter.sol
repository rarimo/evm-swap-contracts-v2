// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Resolver} from "../../libs/Resolver.sol";
import {ErrorHelper} from "../../libs/ErrorHelper.sol";

contract MulticallRouter {
    using Resolver for *;
    using ErrorHelper for *;

    function multicall(
        address[] calldata targets_,
        uint256[] calldata values_,
        bytes[] calldata data_
    ) external payable {
        require(
            targets_.length == values_.length && values_.length == data_.length,
            "MulticallRouter: lengths mismatch"
        );

        for (uint256 i = 0; i < targets_.length; ++i) {
            (bool ok_, bytes memory revertData_) = targets_[i].resolve().call{
                value: values_[i].resolve()
            }(data_[i]);

            require(ok_, revertData_.toStringReason().wrap("MulticallRouter"));
        }
    }
}
