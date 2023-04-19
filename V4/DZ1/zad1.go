package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"crypto/elliptic"
)

const AddressLength = 4

type Hash [32]byte
type Address [AddressLength]byte

type Transaction struct {
	From      Address
	To        Address
	Amount    int
	Signature []byte
}

func NewTransaction(from Address, to Address, amount int) (t *Transaction) {
	t = &Transaction{}
	t.From = from
	t.To = to
	t.Amount = amount
	return
}

func ValuesHash(a ...[]byte) (retval Hash) {
	buffer := new(bytes.Buffer)
	for _, val := range a {
		buffer.Write(val)
	}
	retval = sha256.Sum256(buffer.Bytes())
	return
}

func (t *Transaction) TxHash() Hash {
	return ValuesHash(t.From[:], t.To[:], []byte(fmt.Sprintf("%d", t.Amount))[:])
}

func (t *Transaction) Sign(key ecdsa.PrivateKey, keyp ecdsa.PublicKey) {
	txHash := t.TxHash()
	r, s, err := ecdsa.Sign(rand.Reader, &key, txHash[:])
	if err != nil {
		panic(err)
	}
	signature := append(r.Bytes()[:32], s.Bytes()[:32]...)

	// r2, s2, err2 := ecdsa.Sign(rand.Reader, &key, signature)
	// if err2 != nil {
	// 	panic(err2)
	// }
	// signature = append(signature, r2.Bytes()[:32]...)
	// signature = append(signature, s2.Bytes()[:32]...)

	t.Signature = signature
}

func (t *Transaction) Verify(pubKey ecdsa.PublicKey) bool {
	txHash := t.TxHash()
	if len(t.Signature) != 64 {
		return false
	}
	r := new(big.Int).SetBytes(t.Signature[:32])
	s := new(big.Int).SetBytes(t.Signature[32:])

	return ecdsa.Verify(&pubKey, txHash[:], r, s)
}

func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}

func main() {
	tx := NewTransaction(GenerateNewAddress(), GenerateNewAddress(), 20)

	pubKeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKey := privateKey.PublicKey

	tx.Sign(*privateKey, pubKey)

	fmt.Println("Valid? ", tx.Verify(pubKey))
}
