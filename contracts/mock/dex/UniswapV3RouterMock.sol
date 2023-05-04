// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@dlsl/dev-modules/libs/arrays/ArrayHelper.sol";

import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

import "./AbstractSwapRouterMock.sol";
import "../../libs/BytesHelper.sol";

contract UniswapV3RouterMock is AbstractSwapRouterMock {
    using ArrayHelper for address[];
    using BytesHelper for bytes;

    constructor(address weth9_) AbstractSwapRouterMock(weth9_) {}

    function refundETH() external payable {
        if (address(this).balance > 0) {
            (bool ok_, ) = msg.sender.call{value: address(this).balance}("");
            require(ok_, "UniswapV3RouterMock: failed to refund ETH");
        }
    }

    function exactInput(
        ISwapRouter.ExactInputParams calldata params_
    ) external payable returns (uint256 amountOut_) {
        address[] memory v2Path_ = _toV2Path(params_.path);

        if (msg.value > 0) {
            require(v2Path_[0] == WRAPPED_NATIVE, "UniswapV3RouterMock: wrong input token");
        }

        uint256[] memory amounts_ = _exactIn(
            params_.amountIn,
            params_.amountOutMinimum,
            v2Path_,
            params_.recipient,
            false
        );

        return amounts_[amounts_.length - 1];
    }

    function exactOutput(
        ISwapRouter.ExactOutputParams calldata params_
    ) external payable returns (uint256 amountIn_) {
        address[] memory v2Path_ = _toV2Path(params_.path).reverse();

        if (msg.value > 0) {
            require(
                v2Path_[v2Path_.length - 1] == WRAPPED_NATIVE,
                "UniswapV3RouterMock: wrong output token"
            );
        }

        amountIn_ = _exactOut(
            params_.amountOut,
            params_.amountInMaximum,
            v2Path_,
            params_.recipient,
            false
        )[0];
    }

    function _toV2Path(bytes calldata path_) internal pure returns (address[] memory v2Path_) {
        require((path_.length + 3) % 23 == 0, "UniswapRouterV3Mock: wrong path length");

        v2Path_ = new address[]((path_.length + 3) / 23);

        for (uint256 i = 0; i < v2Path_.length; ++i) {
            v2Path_[i] = path_.toAddress(23 * i);
        }
    }
}
