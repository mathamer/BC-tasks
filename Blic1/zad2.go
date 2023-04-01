// Generiraj niz 10 nasumičnih brojeva
// te pozovi i ispiši rezultate funkcije za pronalazak max vrijednosti i popratnog indeksa.

package main

import (
	"fmt"
	"math/rand"
)

func funkc(arr []int) int {
	max := 0
	ind := 0

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			ind = i
		}
	}
	fmt.Println("index: ", ind)
	return max
}

func main() {

	var niz [9]int

	for i := 0; i <= rand.Intn(9); i++ {
		niz[i] = rand.Intn(9)
	}

	fmt.Println("max: ", funkc(niz[:]))
}
