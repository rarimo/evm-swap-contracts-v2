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
            require(ok_, "UniswapRouterV3Mock: failed to refund ETH");
        }
    }

    function exactInput(
        ISwapRouter.ExactInputParams calldata params_
    ) external payable returns (uint256 amountOut_) {
        address[] memory v2Path_ = _toV2Path(params_.path);

        if (msg.value > 0) {
            require(v2Path_[0] == WRAPPED_NATIVE, "UniswapRouterV2Mock: wrong input token");
        }

        uint256[] memory amounts_ = _getAmountsOut(params_.amountIn, v2Path_);
        amountOut_ = amounts_[amounts_.length - 1];

        require(
            amountOut_ >= params_.amountOutMinimum,
            "UniswapRouterV3Mock: insufficient amount out"
        );

        _swap(amounts_, v2Path_, params_.recipient);
    }

    function exactOutput(
        ISwapRouter.ExactOutputParams calldata params_
    ) external payable returns (uint256 amountIn_) {
        address[] memory v2Path_ = _toV2Path(params_.path).reverse();

        if (msg.value > 0) {
            require(
                v2Path_[v2Path_.length - 1] == WRAPPED_NATIVE,
                "UniswapRouterV2Mock: wrong output token"
            );
        }

        uint256[] memory amounts_ = _getAmountsIn(params_.amountOut, v2Path_);
        amountIn_ = amounts_[0];

        require(
            amounts_[0] <= params_.amountInMaximum,
            "UniswapRouterV3Mock: excessive input amount"
        );

        _swap(amounts_, v2Path_, params_.recipient);
    }

    function _toV2Path(bytes calldata path_) internal pure returns (address[] memory v2Path_) {
        require((path_.length + 3) % 23 == 0, "UniswapRouterV3Mock: wrong path length");

        v2Path_ = new address[]((path_.length + 3) / 23);

        for (uint256 i = 0; i < v2Path_.length; ++i) {
            v2Path_[i] = path_.toAddress(23 * i);
        }
    }
}
