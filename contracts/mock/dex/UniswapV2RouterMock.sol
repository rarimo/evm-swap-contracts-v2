// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./AbstractSwapRouterMock.sol";

contract UniswapV2RouterMock is AbstractSwapRouterMock {
    constructor(address weth9_) AbstractSwapRouterMock(weth9_) {}

    function swapExactTokensForTokens(
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) public returns (uint256[] memory amounts_) {
        amounts_ = _exactIn(amountIn_, amountOutMin_, path_, receiver_, false);
    }

    function swapTokensForExactTokens(
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) public returns (uint256[] memory amounts_) {
        amounts_ = _exactOut(amountOut_, amountInMax_, path_, receiver_, false);
    }

    function swapExactETHForTokens(
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) public payable returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[0] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong input token"
        );

        amounts_ = _exactIn(msg.value, amountOutMin_, path_, receiver_, false);
    }

    function swapTokensForExactETH(
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) external returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[path_.length - 1] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong output token"
        );

        amounts_ = _exactOut(amountOut_, amountInMax_, path_, receiver_, true);
    }

    function swapExactTokensForETH(
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) external returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[0] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong input token"
        );

        amounts_ = _exactIn(amountIn_, amountOutMin_, path_, receiver_, true);
    }

    function swapETHForExactTokens(
        uint256 amountOut_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) external payable returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[path_.length - 1] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong output token"
        );

        amounts_ = _exactOut(amountOut_, msg.value, path_, receiver_, true);

        if (msg.value > amounts_[0]) {
            (bool ok_, ) = msg.sender.call{value: msg.value - amounts_[0]}("");
            require(ok_, "UniswapV2RouterMock: failed to transfer rest");
        }
    }
}
