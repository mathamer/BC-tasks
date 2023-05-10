// TemperatureControl contract
// Napiši pametan ugovor koji omogućuje mijenjanje varijable temperature unutar pametnog ugovora.

// Temperaturu može mjenjati samo vlasnik ugovora ili neke od autoriziranih adresa (može ih biti više).

// Kod kreiranja ugovora, popis autoriziranih adresa je prazan, međutim (samo) vlasnik ugovora može dodati nove adrese putem odgovarajuće funkcije (za koju samo vlasnik ima ovlast pozvati).

pragma solidity >=0.8.2 <0.9.0;

contract TemperatureControl {
    address private owner;
    mapping (address => bool) private authorized;
    uint private temperature;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    modifier onlyAuthorized() {
        require(authorized[msg.sender], "Only authorized can call this function");
        _;
    }

    constructor() {
        owner = msg.sender;
        authorized[owner] = true;
    }

    function addAuthorized(address addr) public onlyOwner {
        authorized[addr] = true;
    }

    function changeTemperature(uint temp) public onlyAuthorized {
        temperature = temp;
    }
}