// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/MockUSDC.sol";
import "../src/OutcomeToken.sol";
import "../src/PredictionMarket.sol";

contract PredictionMarketTest is Test {
    MockUSDC usdc;
    PredictionMarket market;

    function setUp() public {
        usdc = new MockUSDC();
        market = new PredictionMarket(address(usdc));
    }

    function testInitialReservesAreEqual() public {
        (uint256 xA, uint256 xB, uint256 xC) = market.getReserves();

        // All three outcomes should start with the same reserve.
        assertEq(xA, xB, "xA != xB");
        assertEq(xB, xC, "xB != xC");
        assertGt(xA, 0, "reserves should be > 0");
    }

    function testSwapAtoB() public {
        // Set up a trader who holds some outcome A tokens.
        address trader = address(0x123);

        // Only the market (owner) can mint outcome tokens; use cheatcode to mint to trader.
        vm.startPrank(address(market));
        market.outcomeA().mint(trader, 10 ether);
        vm.stopPrank();

        // Approve the market to pull A from the trader.
        vm.startPrank(trader);
        market.outcomeA().approve(address(market), 10 ether);

        (uint256 beforeA, uint256 beforeB, ) = market.getReserves();

        uint256 amountIn = 1 ether;
        uint256 amountOut =
            market.swap(PredictionMarket.Outcome.A, PredictionMarket.Outcome.B, amountIn);

        assertGt(amountOut, 0, "no B received");

        (uint256 afterA, uint256 afterB, ) = market.getReserves();

        // Reserves should update: A up, B down.
        assertEq(afterA, beforeA + amountIn, "xA did not increase correctly");
        assertEq(afterB, beforeB - amountOut, "xB did not decrease correctly");

        vm.stopPrank();
    }
}

