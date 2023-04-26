// Simple bank contract
// Kreiraj smart contract za jednostavnu banku

    // svaki od klijenata prenosi određeni broj svojih novaca na ugovor
    // ugovor interno prati tko je koliko sredstava prenjeo na ugovor

// Za svakog od klijenata banke prati koliko ima novaca (Hint: koristiti mapping)

// Sadrži sljedeće funkcije:
    // deposit() - uvećava se balance pošiljatelja za prenesenu vrijednost (Hint: msg.value)
    // withdraw(uint withdrawAmount) - provjerava da li pošiljatelj transakcije ima dovoljno novaca/balance
        // ako nema baca grešku
        // ako ima smanjuje njegov balance za withdrawAmount i prenosi novac na račun pošiljatelja transakcije
    // balance() - vraća balance pošiljatelja transakcije
    // depositsBalance() - vraća količlinu novaca koji se nalazi na ugovoru (balance ugovora)

    pragma solidity >=0.8.2 <0.9.0;

    contract Banka {
        uint novci;

        constructor(uint cost) {
            novci = cost;
        }

        mapping (address => uint) private balances;

        function deposit() public payable {
            balances[msg.sender] += msg.value;
        }

        function withdraw(uint withdrawAmount) public payable {
            require(balances[msg.sender] >= withdrawAmount, "Not enough funds");
            balances[msg.sender] -= withdrawAmount;
            payable(msg.sender).transfer(withdrawAmount);
        }

        function balance() public view returns (uint) {
            return balances[msg.sender];
        }

        function depositsBalance() public view returns (uint) {
            return address(this).balance;
        }
    }

    