pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract MoonFlipVault {
    address public playerOwner;

    constructor(address _owner) {
        playerOwner = _owner;
    }

    receive() external payable {}

    function withdraw() public {
        require(
            playerOwner == msg.sender,
            "must be the owner to call this function."
        );
        (bool ok, ) = payable(playerOwner).call{value: address(this).balance}(
            ""
        );
        require(ok, "failed to send payment.");
    }
}

contract MoonFlip is Ownable, ReentrancyGuard {
    using Address for address;

    enum CoinSide {
        HEADS,
        TAILS
    }

    enum Game {
        DICE_ROLL,
        COIN_FLIP
    }

    event Outcome(Game game, bool won, address participant, uint256 value);

    // the fraction for 4.44%
    uint128 constant numerator = 111;
    uint128 constant denominator = 2500;

    uint256 private liquidityThreshold = 75 ether;

    bool public paused = false;

    uint256[] public betAmounts = [
        0.05 ether,
        0.1 ether,
        0.5 ether,
        1 ether,
        2.5 ether,
        5.0 ether
    ];

    struct Commit {
        // the block number of this commit.
        uint256 blockNumber;
        // the game nonce
        uint256 nonce;
        uint256 value; // value after fees.
        // the guess of the commit. Reason for guess to be an int vs bool is adapation for other game modes.
        // Move to uint8

        uint8 guess;
        bool exists;
    }

    // the administrator mappings
    mapping(address => bool) _administrators;

    // player address to game nonce to pending commits
    mapping(address => mapping(uint256 => Commit)) _pendingCommits;

    mapping(address => uint256) _gameNonces;

    address public houseAddress;

    modifier onlyAdmin() {
        require(_administrators[msg.sender]);
        _;
    }

    // administrative command to clear a pending transaction
    function clearPendingCommit(address addr, uint256 gameId)
        external
        onlyOwner
    {
        delete _pendingCommits[addr][gameId];
    }

    function addAdmin(address addr) external onlyOwner {
        _administrators[addr] = true;
    }

    function revokeAdmin(address addr) external onlyOwner {
        _administrators[addr] = false;
    }

    function setHouseAddress(address houseAddress_) external onlyOwner {
        require(houseAddress_ != address(0));
        houseAddress = houseAddress_;
    }

    function setBetRange(uint256[] calldata ranges) external onlyOwner {
        betAmounts = ranges;
    }

    function setLiquidityThreshold(uint256 liquidity) external onlyOwner {
        liquidityThreshold = liquidity;
    }

    function setPaused(bool state) external onlyOwner {
        paused = state;
    }

    receive() external payable {}

    function _withdraw(uint256 amount) private {
        (bool ok, ) = payable(owner()).call{value: amount}("");
        require(ok);
    }

    // This function is used to withdraw the amount in ether
    function withdraw(uint256 amount) external onlyOwner {
        _withdraw(amount * 1 ether);
    }

    function withdrawAll() external onlyOwner {
        // this function is used to
        _withdraw(address(this).balance);
    }

    modifier gameChecks() {
        // no smart contracts should be able to play the game...this isn't perfect for stopping bots.
        // but it's good enough.
        require(
            !msg.sender.isContract() && msg.sender == tx.origin,
            "no contracts..."
        );
        // we may need to pause the game for maintainance or debugging. Don't let the game
        require(!paused, "game is currently paused.");
        // don't allow for house payouts
        require(houseAddress != address(0));
        require(
            payable(this).balance >= msg.value * 4,
            "not enough liquidity for payout!"
        );
        require(
            msg.value >= checkForFee(betAmounts[0]) &&
                msg.value <= checkForFee(betAmounts[betAmounts.length - 1]),
            "bet is too high or too low!"
        );
        _;
    }

    event GuessSubmitted(Game game, Commit commit, address sender);

    function processGuess(
        address player,
        uint256 gameId,
        uint256 randomNumber
    ) external onlyAdmin {
        Commit memory commit = _pendingCommits[player][gameId];
        require(commit.exists, "commit does not exist");

        uint256 value = commit.value;
        uint8 gameChoice = uint8(
            randomNumber % 2 == 0 ? CoinSide.HEADS : CoinSide.TAILS
        );
        bool result = gameChoice == commit.guess;

        if (result) {
            value *= 2;
        }

        emit Outcome(Game.COIN_FLIP, result, player, value);

        delete _pendingCommits[player][gameId];

        if (result) {
            (bool ok, ) = payable(player).call{value: value}("");
            require(ok, "failed to send payout"); // execute the payout
        }
    }

    function flipCoin(CoinSide guess, uint256 betIndex)
        public
        payable
        gameChecks
        nonReentrant
    {
        require(betIndex >= 0 && betIndex < betAmounts.length);
        require(
            _pendingCommits[msg.sender][_gameNonces[msg.sender]].blockNumber !=
                block.number,
            "cannot play two games in the same block."
        );

        uint256 fee = getFee(betAmounts[betIndex]);
        uint256 wager = msg.value - fee;

        require(wager == betAmounts[betIndex], "sent too much .");

        uint256 nonce = _gameNonces[msg.sender];

        _pendingCommits[msg.sender][nonce] = Commit({
            blockNumber: block.number,
            guess: uint8(guess),
            nonce: nonce,
            value: wager,
            exists: true
        });

        emit GuessSubmitted(
            Game.COIN_FLIP,
            _pendingCommits[msg.sender][nonce],
            msg.sender
        );

        _gameNonces[msg.sender] = nonce + 1;

        // take the fee and send it to the house.

        // We don't want to keep all our funds in the contract. If we have a reasonable amount of liquidity in the contract
        // that will be replenished, then
        (bool ok, ) = payable(houseAddress).call{
            value: (
                address(this).balance >= liquidityThreshold ? msg.value : fee
            )
        }("");
        require(ok, "failed to send payment");
    }

    function checkForFee(uint256 num) internal pure returns (uint256) {
        uint256 value = num * numerator;
        value /= denominator;
        return num + value;
    }

    function getFee(uint256 num) internal pure returns (uint256) {
        uint256 value = num * numerator;
        value /= denominator;
        return value;
    }
}
