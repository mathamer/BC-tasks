// Ispiši dnevne zahtjeve web servera
// podaci sadrže 8-satne totale za svaki dan
// Svi podaci su zapisani u jednom slice-u, ali se svaki 3 uzastopni podaci odnose na jedan dan
// ispiši totale za svaki dan, a zatim i sumu svih zahjeva
// koristiti dnevne slice-ove (kreirati slice za svakih 3 podataka i onda raditi operacije nad njime)

//    ulaz: reqs := []int{
//    	500, 600, 250, // 1. dan
//    	900, 800, 600, // 2. dan
//    	150, 654, 235, // 3. dan
//    	121, 876, 285, // 4. dan
//      }

package main

import (
	"fmt"
)

func suma(arr []int) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

func main() {

	var niz [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&niz[i])
	}

	fmt.Println(niz)

	fmt.Println(suma(niz[:]))
}
