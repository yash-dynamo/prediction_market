// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Sphere AMM–style multi‑outcome prediction market for three outcomes A, B, C.
// Reserves xA, xB, xC lie on the sphere:
//   ||r − x||² = r²  where  r = (r, r, r),  x = (xA, xB, xC)
//
// Local price between outcome tokens (from Orbital / Sphere AMM):
//   dxi/dxj = (r − xj) / (r − xi)

import "./OutcomeToken.sol";
import "./MockUSDC.sol";

/// @notice Minimal Chainlink price feed interface (v3).
interface AggregatorV3Interface {
    function decimals() external view returns (uint8);
    function description() external view returns (string memory);
    function latestRoundData()
        external
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        );
}

contract PredictionMarket {
    enum Outcome {
        A,
        B,
        C
    }

    OutcomeToken public outcomeA;
    OutcomeToken public outcomeB;
    OutcomeToken public outcomeC;

    // Optional collateral token (not wired into the Sphere math yet).
    MockUSDC public usdc;

    // AMM reserves of each outcome token.
    uint256 public xA;
    uint256 public xB;
    uint256 public xC;

    // Sphere parameter r (same units as reserves).
    uint256 public r;

    uint256 public constant ONE = 1e18;
    uint256 public constant N = 3;


    /*//////////////////////////////////////////////////////////////
        Oracle configuration (Chainlink) for price-based resolution
    //////////////////////////////////////////////////////////////*/

    AggregatorV3Interface public priceFeed; // e.g., BTC / USD
    uint256 public resolveTime;             // UNIX timestamp when market can be resolved
    uint256 public maxStaleTime;            // max allowed staleness for Chainlink answer

    // Optional 3-way price buckets: A, B, C decided by bounds in feed units.
    int256 public lowerBound; // e.g., price < lowerBound => Outcome A
    int256 public upperBound; // e.g., lowerBound <= price < upperBound => Outcome B, else C

    bool public resolved;
    Outcome public winningOutcome;
    int256 public resolvedPrice;

    address public owner;

    event Swap(Outcome indexed from, Outcome indexed to, uint256 amountIn, uint256 amountOut);

    constructor(address _usdc) {
        owner = msg.sender;
        usdc = MockUSDC(_usdc);

        outcomeA = new OutcomeToken("Outcome A", "A", address(this));
        outcomeB = new OutcomeToken("Outcome B", "B", address(this));
        outcomeC = new OutcomeToken("Outcome C", "C", address(this));

        // Initialize at the equal‑price point q for all three outcomes.
        // From Orbital:
        //   q = r * (1 − 1/n)  ⇒  r = q * n / (n − 1)
        uint256 q = 100 * ONE;
        r = (q * N) / (N - 1);

        xA = q;
        xB = q;
        xC = q;

        // Mint initial liquidity to the AMM itself.
        outcomeA.mint(address(this), q);
        outcomeB.mint(address(this), q);
        outcomeC.mint(address(this), q);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "not owner");
        _;
    }

    /// @notice Configure Chainlink feed and resolution time/staleness.
    /// @param feed Address of Chainlink AggregatorV3 feed (crypto, FX, index, etc.).
    /// @param _resolveTime UNIX timestamp when the market may be resolved.
    /// @param _maxStaleTime Max seconds between Chainlink update and resolution.
    function setOracleConfig(
        address feed,
        uint256 _resolveTime,
        uint256 _maxStaleTime
    ) external onlyOwner {
        require(feed != address(0), "feed=0");
        require(_resolveTime > block.timestamp, "resolve in past");
        require(_maxStaleTime > 0, "stale=0");

        priceFeed = AggregatorV3Interface(feed);
        resolveTime = _resolveTime;
        maxStaleTime = _maxStaleTime;
    }

    /// @notice Configure 3-way price buckets in the feed's native units.
    /// A: price < lowerBound
    /// B: lowerBound <= price < upperBound
    /// C: price >= upperBound
    function setPriceBuckets(int256 _lowerBound, int256 _upperBound) external onlyOwner {
        require(_lowerBound < _upperBound, "bounds");
        lowerBound = _lowerBound;
        upperBound = _upperBound;
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function getReserves() external view returns (uint256, uint256, uint256) {
        return (xA, xB, xC);
    }

    /// "Probabilities" derived from Sphere geometry: weights w_i = (r - x_i), normalized.
    function getProbabilities()
        external
        view
        returns (uint256 pA, uint256 pB, uint256 pC)
    {
        uint256 wA = r - xA;
        uint256 wB = r - xB;
        uint256 wC = r - xC;

        uint256 sumW = wA + wB + wC;
        require(sumW > 0, "invalid state");

        pA = (wA * ONE) / sumW;
        pB = (wB * ONE) / sumW;
        pC = (wC * ONE) / sumW;
    }

    /// Left‑hand side of the Sphere invariant ‖r⃗ − x⃗‖²; should be close to r².
    function getSphereInvariant() external view returns (uint256) {
        uint256 a = r - xA;
        uint256 b = r - xB;
        uint256 c = r - xC;
        return a * a + b * b + c * c;
    }

    /// @notice Returns the latest Chainlink price and timestamp (if configured).
    function getLatestPrice() external view returns (int256 price, uint256 updatedAt) {
        require(address(priceFeed) != address(0), "no feed");
        (, int256 answer, , uint256 upd, ) = priceFeed.latestRoundData();
        return (answer, upd);
    }

    /// @notice Resolve the market based on the Chainlink price at/after resolveTime.
    /// Uses configured price buckets to map the observed price to Outcome A/B/C.
    function resolve() external {
        require(!resolved, "resolved");
        require(address(priceFeed) != address(0), "no feed");
        require(resolveTime != 0, "no resolveTime");
        require(block.timestamp >= resolveTime, "too early");

        (, int256 answer, , uint256 updatedAt, ) = priceFeed.latestRoundData();
        require(answer > 0, "bad price");
        require(block.timestamp - updatedAt <= maxStaleTime, "stale price");

        resolvedPrice = answer;
        resolved = true;

        // Decide winning outcome using configured buckets.
        if (answer < lowerBound) {
            winningOutcome = Outcome.A;
        } else if (answer < upperBound) {
            winningOutcome = Outcome.B;
        } else {
            winningOutcome = Outcome.C;
        }
    }

    /*//////////////////////////////////////////////////////////////
                              TRADING LOGIC
    //////////////////////////////////////////////////////////////*/

    /// Swap from one outcome token into another while preserving the Sphere invariant.
    ///
    /// Mathematically, with a = r - x_from, b = r - x_to, dx = amountIn, dy = amountOut:
    ///   dy² + 2 b dy = 2 a dx − dx²
    /// so:
    ///   dy = -b + sqrt(b² + 2 a dx − dx²).
    function swap(Outcome from, Outcome to, uint256 amountIn)
        external
        returns (uint256 amountOut)
    {
        require(from != to, "same outcome");
        require(amountIn > 0, "zero amount");

        // Select reserves and token contracts.
        OutcomeToken fromToken;
        OutcomeToken toToken;
        uint256 fromReserve;
        uint256 toReserve;

        if (from == Outcome.A) {
            fromToken = outcomeA;
            fromReserve = xA;
        } else if (from == Outcome.B) {
            fromToken = outcomeB;
            fromReserve = xB;
        } else {
            fromToken = outcomeC;
            fromReserve = xC;
        }

        if (to == Outcome.A) {
            toToken = outcomeA;
            toReserve = xA;
        } else if (to == Outcome.B) {
            toToken = outcomeB;
            toReserve = xB;
        } else {
            toToken = outcomeC;
            toReserve = xC;
        }

        uint256 a = r - fromReserve;
        uint256 b = r - toReserve;

        // Ensure we stay in the region where the quadratic has a real solution.
        require(amountIn <= 2 * a, "amount too large");

        uint256 term = b * b + 2 * a * amountIn - amountIn * amountIn;
        uint256 root = sqrt(term);
        amountOut = root - b;

        require(amountOut > 0, "zero out");
        require(amountOut <= toReserve, "insufficient liquidity");

        // Move tokens.
        fromToken.transferFrom(msg.sender, address(this), amountIn);
        toToken.transfer(msg.sender, amountOut);

        // Update reserves.
        if (from == Outcome.A) {
            xA += amountIn;
        } else if (from == Outcome.B) {
            xB += amountIn;
        } else {
            xC += amountIn;
        }

        if (to == Outcome.A) {
            xA -= amountOut;
        } else if (to == Outcome.B) {
            xB -= amountOut;
        } else {
            xC -= amountOut;
        }

        emit Swap(from, to, amountIn, amountOut);
    }

    /*//////////////////////////////////////////////////////////////
                          MATH HELPERS
    //////////////////////////////////////////////////////////////*/

    function sqrt(uint256 y) internal pure returns (uint256 z) {
        if (y == 0) return 0;
        if (y <= 3) {
            return 1;
        }
        z = y;
        uint256 x = y / 2 + 1;
        while (x < z) {
            z = x;
            x = (y / x + x) / 2;
        }
    }
}