// 1. Napiši program koji

// Kreira goroutinu koja generira nasumične brojeve nasumičan broj puta u rasponu 1 - 100
// Brojeve sprema u kanal ako su parni
// Nakon završetka zatvara kanal
// U glavoj goroutini ispisuje te brojeve iz kanala sve dok se kanal ne zatvori

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gorutina() {
	ch := make(chan int, rand.Intn(100))

	for i := 0; i < cap(ch); i++ {
		temp := rand.Intn(100)
		if temp%2 == 0 {
			ch <- temp
			fmt.Println(<-ch)
		}
	}
	close(ch)
}

func main() {

	go gorutina()
	time.Sleep(2 * time.Second)
}
