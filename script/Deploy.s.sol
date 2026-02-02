pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/MockUSDC.sol";
import "../src/OutcomeToken.sol";
import "../src/PredictionMarket.sol";

contract Deploy is Script{
    function run() external{
        vm.startBroadcast();

        MockUSDC usdc = new MockUSDC();
        console.log("MockUSDC deployed at :", address(usdc));

        PredictionMarket market = new PredictionMarket(address(usdc));
        console.log("PredictionMarket deployed at:", address(market));

        vm.stopBroadcast();
    }
}