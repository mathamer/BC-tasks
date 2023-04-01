// Pretvori i sortiraj niz stringova
// kreiraj niz cijelih brojeva koji su u obliku stringova
// pretvori niz stringova u niz integera
// sortiraj
// ispiši

// ulaz: {"2", "3", "-1", "34", "53"}
// izlaz: {-1, 2, 3, 34, 53}

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var niz = [5]string{"2", "3", "-1", "34", "53"}

	var intNiz [5]int

	fmt.Println(niz)

	for i := 0; i < len(niz); i++ {

		// string to int
		temp, err := strconv.Atoi(niz[i])
		if err != nil {
			// ... handle error
			panic(err)
		}
		intNiz[i] = temp
	}

	sort.Ints(intNiz[:])

	fmt.Println(intNiz)

}
