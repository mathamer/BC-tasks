// https://hackmd.io/@rosaj/Bk-wqdy4n
// 1. Marketplace
// In this assignment, you will be creating a Solidity smart contract for an online marketplace where users can buy and sell items using cryptocurrency. 
// Your contract will include functions for buying and selling items, as well as modifiers to ensure that only authorized users can access certain functions. 
// You will also need to define a structure for the items being sold.

// Requirements:

// Define a structure for the items being sold. Each item should have the following properties:
// name (string): the name of the item
// description (string): a description of the item
// price (uint): the price of the item in wei
// seller (address): the address of the seller
// Define a mapping to keep track of the items being sold. The keys for the mapping should be the item IDs (uint), and the values should be the items themselves (struct).
// Write a function for adding a new item to the marketplace. This function should take in the name, description, and price of the item, and add it to the mapping of items. The seller should be set as the address of the person who called the function.
// Write a function for buying an item from the marketplace. This function should take in the ID of the item being purchased, and transfer the appropriate amount of cryptocurrency from the buyer to the seller. The item should then be removed from the mapping of items.
// Write a modifier to ensure that only the seller of an item can remove it from the marketplace.
// Write a modifier to ensure that only authorized buyers can purchase items from the marketplace.
// Write a function for removing an item from the marketplace. This function should only be callable by the seller of the item, and should remove the item from the mapping of items.
// Write a function for retrieving the details of an item. This function should take in the ID of the item, and return the name, description, price, and seller address.
// Write a function for retrieving the number of items currently for sale on the marketplace.

pragma solidity >=0.8.2 <0.9.0;

contract Marketplace {
    struct Item {
        string name;
        string description;
        uint price;
        address seller;
    }

    mapping (uint => Item) private items;

    uint private itemsCount;

    modifier onlySeller(uint id) {
        require(msg.sender == items[id].seller, "Only seller can remove item");
        _;
    }

    modifier onlyBuyer(uint id) {
        require(msg.sender != items[id].seller, "Only buyer can buy item");
        _;
    }

    modifier enoughFunds(uint id) {
        require(msg.value >= items[id].price, "Not enough funds");
        _;
    }

    function addItem(string memory name, string memory description, uint price) public {
        items[itemsCount] = Item(name, description, price, msg.sender);
        itemsCount++;
    }

    function buyItem(uint id) public payable onlyBuyer(id) enoughFunds(id) {
        payable(items[id].seller).transfer(items[id].price);
        delete items[id];
    }

    function removeItem(uint id) public onlySeller(id) {
        delete items[id];
    }

    function getItem(uint id) public view returns (string memory, string memory, uint, address) {
        return (items[id].name, items[id].description, items[id].price, items[id].seller);
    }

    function getItemsCount() public view returns (uint) {
        return itemsCount;
    }
}