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
            "UniswapRouterV2Mock: insufficient amount out"
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

        require(amounts_[0] <= amountInMax_, "UniswapRouterV2Mock: excessive input amount");

        _swap(amounts_, path_, receiver_);
    }

    function swapExactETHForTokens(
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) public payable returns (uint256[] memory amounts_) {
        require(
            path_.length > 2 && path_[0] == WRAPPED_NATIVE,
            "UniswapRouterV2Mock: wrong input token"
        );

        amounts_ = swapExactTokensForTokens(msg.value, amountOutMin_, path_, receiver_, deadline_);
    }

    function swapTokensForExactETH(
        uint256 amountOut_,
        uint256 amountInMax_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) external returns (uint256[] memory amounts_) {
        require(
            path_.length > 2 && path_[path_.length - 1] == WRAPPED_NATIVE,
            "UniswapRouterV2Mock: wrong output token"
        );

        amounts_ = swapTokensForExactTokens(amountOut_, amountInMax_, path_, receiver_, deadline_);
    }

    function swapExactTokensForETH(
        uint256 amountIn_,
        uint256 amountOutMin_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) external returns (uint256[] memory amounts_) {
        require(
            path_.length > 2 && path_[0] == WRAPPED_NATIVE,
            "UniswapRouterV2Mock: wrong input token"
        );

        amounts_ = swapExactTokensForTokens(amountIn_, amountOutMin_, path_, receiver_, deadline_);
    }

    function swapETHForExactTokens(
        uint256 amountOut_,
        address[] calldata path_,
        address receiver_,
        uint256 deadline_
    ) external payable returns (uint256[] memory amounts_) {
        require(
            path_.length > 2 && path_[path_.length - 1] == WRAPPED_NATIVE,
            "UniswapRouterV2Mock: wrong output token"
        );

        amounts_ = swapTokensForExactTokens(amountOut_, msg.value, path_, receiver_, deadline_);
    }
}
