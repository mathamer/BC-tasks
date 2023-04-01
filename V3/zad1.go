// Game store

// Deklariraj strukture
// item: id (int), name (string), price (int)
// game: anonimno polje item, genre (string)
// Kreiraj game slice s podacima:

// id	name	price	Genre
// 1	god of war	50	action adventure
// 2	x-com 2	30	strategy
// 3	minecraft	20	sandbox
// Iteriraj kroz slice i ispiši sve podatke struktura

package main

import (
	"fmt"
)

type game struct {
	item
	genre string
}

type item struct {
	id    int
	name  string
	price int
}

func main() {
	g := game{
		genre: "action adventure",
		item: item{
			id:    1,
			name:  "god of war",
			price: 50,
		},
	}
	g1 := game{
		genre: "strategy",
		item: item{
			id:    2,
			name:  "x-com 2",
			price: 30,
		},
	}
	g2 := game{
		genre: "sandbox",
		item: item{
			id:    3,
			name:  "minecraft",
			price: 20,
		},
	}

	games := []game{g, g1, g2}

	for _, i := range games {
		fmt.Println(i.id, i.name, i.price, i.genre)
	}
}
