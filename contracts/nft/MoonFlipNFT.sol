pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Strings.sol";



import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";



contract MoonFlipNFT is ERC721Enumerable, Ownable, ReentrancyGuard {
    using Strings for uint256;
    using Address for address;

    enum MintState {
        PUBLIC,
        WHITELIST,
        CLOSED
    }

    uint16 public constant MAX_SUPPLY = 555;

    uint16 public constant UNIQUE_ART = 2;


    // the max supply - set after
    uint16 public maxSupply = MAX_SUPPLY;


    // keeps track of the current total supply.
    uint256 public _totalSupply;

    // the base uri for the token
    string public baseURI;

    // the state of whether minting is allowed.
    MintState private mintingState = MintState.CLOSED;

    mapping(address => bool) _whitelisted;
    mapping(address => bool) _claimed;

    constructor() ERC721("MoonFlip NFTs", "MOONFLIP") {
    }
    

    function addWhitelistAddresses(address[] calldata addr) external onlyOwner {
        for (uint256 i = 0; i < addr.length; i++) {
            _whitelisted[addr[i]] = true;
        }
    }

    function setBaseURI(string calldata _uri) external onlyOwner {
        baseURI = _uri;
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override
        returns (string memory)
    {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        // double ternary op because i like to watch the wolrd burn
        return (
                    bytes(baseURI).length > 0
                        ? string(
                            abi.encodePacked(
                                baseURI,
                                tokenId % UNIQUE_ART,
                                ".json"
                            )
                        )
                        : ""
                );
    }

    modifier mintOpen() {
        require(mintingState != MintState.CLOSED, "minting not active.");
        _;
    }

    modifier whitelistOnly() {
        require(mintingState == MintState.WHITELIST);
        _;
    }

    function ownerMint(uint256 amount) external onlyOwner {
        // ensure minting the amount entered would not mint over the max supply
        require(
            _totalSupply + amount < uint256(maxSupply),
            "mint amount would be out of range."
        );

        for (uint256 i = _totalSupply; i < _totalSupply + amount; ++i) {
            _safeMint(_msgSender(), i + 1);
        }

        _totalSupply += amount;
    }
    
    function claim() external mintOpen nonReentrant {
        require(!_msgSender().isContract(), "smart contracts not allowed!");

        if (mintingState == MintState.WHITELIST) {
            require(_whitelisted[_msgSender()], "not whitelisted!");
        }
        require(!_claimed[_msgSender()], "already claimed!"); // can only claim an NFT once
        require(_totalSupply + 1 < uint256(maxSupply), "minted out!"); // cannot mint if all 555 are already minted.

        _safeMint(_msgSender(), _totalSupply + 1);

        _totalSupply += 1;
        _claimed[_msgSender()] = true;
    }

    function totalSupply() override public view returns(uint256) {
        return _totalSupply;
    }

    // toggles the state of mint
    function setMintingState(MintState state) external onlyOwner {
        mintingState = state;
    }
}