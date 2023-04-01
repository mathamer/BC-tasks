// Napiši funkciju koja kao rezultat vraća maksimalnu vrijednost i indeks maksimalne vrijednosti niza
// koji je predan kao parametar funkciji.

package main

import (
	"fmt"
)

func funkcija(arr []int) int {
	max := 0
	ind := 0
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

	var niz [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&niz[i])
	}

	fmt.Println(niz)

	fmt.Println("max: ", funkcija(niz[:]))
}
