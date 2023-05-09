// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

library ErrorHelper {
    string internal constant ERROR_DELIMITER = ": ";

    function toStringReason(bytes memory data_) internal pure returns (string memory) {
        if (data_.length < 68) {
            return "ErrorHelper: command reverted silently";
        }

        assembly {
            data_ := add(data_, 0x04)
        }

        return abi.decode(data_, (string));
    }

    function wrap(
        string memory error_,
        string memory prefix_
    ) internal pure returns (string memory) {
        return string(abi.encodePacked(prefix_, ERROR_DELIMITER, error_));
    }
}
