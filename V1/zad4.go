// Napiši funkciju koja čita korisnički unos u komandoj liniji

// funkcija vraća podatak koji je korisnik upisao i njegovu duljinu (broj karaktera)
// napisati poziv te funkcije i ispis vrijednosti funkcije

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var dobitni = rand.Intn(9)
	var num = 0
	fmt.Scan(&num)

	for i := 1; i <= 5; i++ {
		dobitni = rand.Intn(9)
		if dobitni == num && i == 1 {
			fmt.Println("Pobjednik isprve! WOW!!!")
			break
		} else if dobitni == num {
			fmt.Println("Bravo, pogodio si broj")
			break
		} else if i == 5 {
			fmt.Println("Pokušaj ponovno :(")
		}
	}
}
