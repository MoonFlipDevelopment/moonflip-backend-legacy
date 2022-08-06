pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

interface IERC20 {
    function balanceOf(address _addr) external returns (uint256);

    function transfer(address _to, uint256 _value)
        external
        returns (bool);
}

contract MoonFlipVault {
    error NotEnoughSignatures();

    mapping(bytes32 => bool) completedProposals;
    mapping(address => bool) signers;

    address owner;

    // the current number of signers
    uint256 numOfSigners;

    // this is the current nonce of withdrawal proposals
    uint256 nonce;

    constructor() {
        owner = msg.sender;
    }

    receive() external payable {}

    function addSigner(address _signer) external onlyOwner {
        signers[_signer] = true;
        numOfSigners++;
    }

    function removeSigner(address _signer) external onlyOwner {
        signers[_signer] = false;
        numOfSigners--;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "not the owner");
        _;
    }

    function hashProposalToSign(
        address token,
        address to,
        uint256 amount
    ) public view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(address(this), token, to, amount, nonce)
            );
    }

    function hashNativeProposalToSign(address to, uint256 amount)
        public
        view
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked(address(this), address(0), to, amount, nonce)
            );
    }

    function withdrawToken(
        address token,
        address to,
        uint256 amount,
        bytes[] calldata signatures
    ) external onlyOwner {
        bytes32 withdrawalHash = hashProposalToSign(address(token), to, amount);

        require(
            !completedProposals[withdrawalHash],
            "proposal already completed!"
        );
        require(
            IERC20(token).balanceOf(address(this)) >= amount,
            "number too large to withdraw"
        );
        require(numOfSigners > 1, "not enough signers");
        require(
            signatures.length >= numOfSigners - 1,
            "must have at least one minus the total number of signers"
        );

        uint256 validSignatures = 0;
        for (uint256 i; i < signatures.length; i++) {
            address addr = ECDSA.recover(
                ECDSA.toEthSignedMessageHash(withdrawalHash),
                signatures[i]
            );

            if (signers[addr]) {
                validSignatures++;
            }
        }

        // this is an (total - 1) / ( total ), example: 2/3, 3/4.
        if (validSignatures >= numOfSigners - 1) {
            nonce++; // if the proposal fails, i still want nonce to be incremented. this can prevent locks.
            completedProposals[withdrawalHash] = true;

            bool ok = IERC20(token).transfer(to, amount);
            require(ok, "failed to send");
        } else {
            revert NotEnoughSignatures();
        }
    }

    /**
     Withdraws to a given address, assuming all the signers signed the proposal message
    */
    function withdraw(
        address to,
        uint256 amount,
        bytes[] calldata signatures
    ) external onlyOwner {
        bytes32 withdrawalHash = hashProposalToSign(address(0), to, amount);

        require(
            !completedProposals[withdrawalHash],
            "proposal already completed!"
        );
        require(
            address(this).balance >= amount,
            "number too large to withdraw"
        );
        require(numOfSigners > 1, "not enough signers");
        require(
            signatures.length >= numOfSigners - 1,
            "must have at least one minus the total number of signers"
        );

        uint256 validSignatures = 0;
        for (uint256 i; i < signatures.length; i++) {
            address addr = ECDSA.recover(
                ECDSA.toEthSignedMessageHash(withdrawalHash),
                signatures[i]
            );

            if (signers[addr]) {
                validSignatures++;
            }
        }

        // this is an (total - 1) / ( total ), example: 2/3, 3/4.
        if (validSignatures >= numOfSigners - 1) {
            nonce++; // if the proposal fails, i still want nonce to be incremented. this can prevent locks.
            completedProposals[withdrawalHash] = true;

            (bool ok, ) = payable(to).call{value: amount}("");
            require(ok, "failed to send");
        } else {
            revert NotEnoughSignatures();
        }
    }
}
