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
        amounts_ = _getAmountsOut(amountIn_, path_);

        require(
            amounts_[amounts_.length - 1] >= amountOutMin_,
            "UniswapV2RouterMock: insufficient amount out"
        );

        _swap(amounts_, path_, receiver_);
    }

    function swapTokensForExactTokens(
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_,
        address receiver_,
        uint256
    ) public returns (uint256[] memory amounts_) {
        amounts_ = _getAmountsIn(amountOut_, path_);

        require(amounts_[0] <= amountInMax_, "UniswapV2RouterMock: excessive input amount");

        _swap(amounts_, path_, receiver_);
    }

    function swapExactETHForTokens(
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) public payable returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[0] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong input token"
        );

        amounts_ = _getAmountsOut(msg.value, path_);

        require(
            amounts_[amounts_.length - 1] >= amountOutMin_,
            "UniswapV2RouterMock: insufficient amount out"
        );

        _swap(amounts_, path_, receiver_);
    }

    function swapTokensForExactETH(
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) external returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[path_.length - 1] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong output token"
        );

        amounts_ = _getAmountsIn(amountOut_, path_);

        require(amounts_[0] <= amountInMax_, "UniswapV2RouterMock: excessive input amount");

        _swap(amounts_, path_, receiver_);
    }

    function swapExactTokensForETH(
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) external returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[0] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong input token"
        );

        amounts_ = _getAmountsOut(amountIn_, path_);

        require(
            amounts_[amounts_.length - 1] >= amountOutMin_,
            "UniswapV2RouterMock: insufficient amount out"
        );

        _swap(amounts_, path_, receiver_);
    }

    function swapETHForExactTokens(
        uint256 amountOut_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) external payable returns (uint256[] memory amounts_) {
        require(
            path_.length >= 2 && path_[path_.length - 1] == WRAPPED_NATIVE,
            "UniswapV2RouterMock: wrong output token"
        );

        amounts_ = _getAmountsIn(amountOut_, path_);

        require(amounts_[0] <= msg.value, "UniswapV2RouterMock: excessive input amount");

        _swap(amounts_, path_, receiver_);

        if (msg.value > amounts_[0]) {
            (bool ok_, ) = msg.sender.call{value: msg.value - amounts_[0]}("");
            require(ok_, "UniswapV2RouterMock: failed to transfer rest");
        }
    }
}
