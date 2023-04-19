// Pratiti subscription samo za jednok korisnika

pragma solidity >=0.8.2 <0.9.0;

contract CellSubscription {
    uint monthlyCost;

    address private owner;

    // Pri kreiranju ugovora, samo jednom
    constructor(uint cost) {
        monthlyCost = cost;
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
    }

    function makePayment() public payable {
    } 

    receive() external payable {} 

    function isPaid() public view returns (bool) {
        return monthlyCost <= address(this).balance;
    }

    function withdraw() public {
        // Prenesi novac sa ugovora na onog tko je napravio zahtjev withdraw funkciju
        // Ako sam T-mobile tek onda transfer, inace nista

        if(isPaid() && msg.sender == owner){
            payable(msg.sender).transfer(address(this).balance);
        }
    }
}