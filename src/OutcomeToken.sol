// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract OutcomeToken is ERC20, Ownable {
    constructor(
        string memory name_,
        string memory symbol_,
        address market
    ) ERC20(name_, symbol_) Ownable(market) {}

    function mint(address to, uint256 amount) external onlyOwner{
        _mint(to,amount);
    }

    function burn(address from, uint256 amount) external onlyOwner {
        _burn(from, amount);
    }
}