// Kreiraj strukture s klasnog diagrama
// Kreiraj funkciju `GetGenesisBlock` koja uvijek vraća "prazan" blok
// Kreiraj funkciju `CreateBlock` koja služi kao konstruktor `Block`-a i prima parametare `number` i `previousBlock`
// funkcija u novo kreirani blok upisuje trenutno vrijeme, Number varijablu, te postavlja referencu na prethodni blok
// inicijalizira Transaction slice
// U `main` funkciji kreirajte 2 bloka pomoću funkcije `CreateBlock`
// Prilikom kreiranja prvog bloka, za `prevBlock` uzima se block koji vraća `GetGenesisBlock`, a za kreiranje drugog bloka uzima se referenca na prvog

package main

import "fmt"

type Transaction struct {
	from   string
	to     string
	amount int
}

type Block struct {
	Number int
	Time   string
	txs    []Transaction
}

func GetGenesisBlock() *Block {
	return &Block{
		Number: 0,
		Time:   "now",
		txs:    []Transaction{},
	}
}

func CreateBlock(number int, prevBlock *Block) *Block {
	return &Block{
		Number: number,
		Time:   "now",
		txs:    []Transaction{},
	}
}

func (b *Block) AddTxs(t Transaction) {
	b.txs = append(b.txs, t)
}

func (b *Block) String() string {
	x := fmt.Sprintf("Block #%d:\n", b.Number)
	for _, t := range b.txs {
		x += fmt.Sprintf("  From: %s, To: %s, Amount: %d\n", t.from, t.to, t.amount)
	}
	return x
}

func main() {
	t1 := Transaction{"user", "user2", 5000}
	t2 := Transaction{"user3", "user", 1000}

	block := CreateBlock(1, GetGenesisBlock())
	block.AddTxs(t1)
	block.AddTxs(t2)
	block2 := CreateBlock(2, block)

	fmt.Println("Transactions:")
	fmt.Print(block.String())
	fmt.Print(block2.String())
}
