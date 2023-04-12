// Kreiraj lanac povezanih struktura Node

// Za referencu (pokazivač) na prethodnu strukturu koristite njezin hash
// Hash() metoda strukture Node vraća hash cijele Node strukture

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Data     string
	PrevHash Hash
}

type Hash [32]byte

func (n *Node) Hash() Hash {
	return ValuesHash([]byte(n.Data), n.PrevHash[:])
}

func ValuesHash(a ...[]byte) (retval Hash) {
	buffer := new(bytes.Buffer)
	for _, val := range a {
		buffer.Write(val)
	}
	retval = sha256.Sum256(buffer.Bytes())
	return
}

func (n *Node) addElement(data string) *Node {
	newNode := Node{Data: data, PrevHash: n.Hash()}
	return &newNode
}

func main() {
	a := Node{Data: "<GENESIS>"}
	b := a.addElement("a")
	d := b.addElement("b")

	a1 := Node{Data: "<GENESIS>"}
	b1 := a1.addElement("a")
	d1 := b1.addElement("b")

	// hash mora biti jednak za oba lanca
	fmt.Printf("Prvi blockchain: %s \n", d.Hash())
	fmt.Printf("Drugi blockchain: %s \n", d1.Hash())
}
