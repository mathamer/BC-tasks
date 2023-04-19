//     Napraviti ugovor za vlastiti coin ugovor
// Funkcionalnosti:
// balanceOf(korisnik) => vraća količinu novaca za korisnika
// transfer(kome, koliko) => funkcija prebacuje s jednog računa na drugi
// deposit() => funkcija za pošiljatelja bilježi koliko je poslao novaca

    pragma solidity >=0.8.2 <0.9.0;

    contract MatijaCoin {
        uint Mcoin;

        constructor(uint cost) {
            Mcoin = cost;
        }

        mapping (address => uint) private balances;

        function balanceOf(address user) public view returns (uint) {
            return balances[user];
        }

        function transfer(address to, uint amount) public {
            require(balances[msg.sender] >= amount, "Not enough funds");
            balances[msg.sender] -= amount;
            balances[to] += amount;
        }

        function deposit() public payable {
            balances[msg.sender] += msg.value;
        }
    }

    