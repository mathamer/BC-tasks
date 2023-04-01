// Kreiraj loto: program generira nasumične brojeve i provjerava da li je to vrijednost koju je korisnik upisao

// pretpostavljena korisnička vrijednost: 6 (može je i korisnik unjeti)
// maksimalni broj pokušaja: 5 (broj iteracija i nasumičnih izvlačenja broja)
// Ispis:
// ako korisnik pogodi broj iz prvog pokušaja ispisuje: "Pobjednik isprve! WOW!!!"
// Inače, ako pogodi ispisuje: "Bravo, pogodio si broj"
// ako korisnik ne pogodi u maksimalnom broj pokušaja ispisuje se poruka: "Pokušaj ponovno :("

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
