// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Invariant: xA*xB + xB*xB + xC*xC=k

import "./OutcomeToken.sol";
import "./MockUSDC.sol";

contract PredictionMarket{
    OutcomeToken public outcomeA;
    OutcomeToken public outcomeB;
    OutcomeToken public outcomeC;

    MockUSDC public usdc;

    uint public xA;
    uint public xB;
    uint public xC;

    uint public invariantK;
    uint256 constant ONE = 1e18;


    constructor(address _usdc){
        usdc = MockUSDC(_usdc);

        outcomeA = new OutcomeToken("Outcome A","A", address(this));
        outcomeB = new OutcomeToken("Outcome B", "B", address(this));
        outcomeC = new OutcomeToken("Outcome C", "c", address(this));

        xA= 100* ONE;
        xB= 100* ONE;
        xC= 100* ONE;

        invariantK = xA*xB+xB*xB+xC*xC;
    }

     /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/


    function getProbabilities() external view returns(uint256 pA, uint256 pB, uint256 pC){
        uint256 sum = xA+xB+xC;
        pA=(xA*ONE)/sum;
        pB = (xB * ONE) / sum;
        pC = (xC * ONE) / sum;
    }


    /*//////////////////////////////////////////////////////////////
                              TRADING LOGIC
    //////////////////////////////////////////////////////////////*/

    function buyA(uint256 usdcIn) external{
        require(usdcIn>0, "Zero trade");

        usdc.transferFrom(msg.sender, address(this), usdcIn);

        uint256 oldxA = xA;

        xA += usdcIn * ONE;

        uint256 remaining= invariantK-(xA*xA);

        uint256 bcSquared = xB*xB+xC*xC;

        uint scale= sqrt((remaining*ONE)/bcSquared);

        xB=(xB*scale)/ONE;
        xC=(xC*scale)/ONE;

        uint256 sharesMinted = xA - oldxA;
        outcomeA.mint(msg.sender, sharesMinted);


    }


    /*//////////////////////////////////////////////////////////////
                          MATH HELPERS
    //////////////////////////////////////////////////////////////*/

    function sqrt(uint256 y) internal pure returns (uint256 z) {
        if (y > 3) {
            z = y;
            uint256 x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
    }





}
