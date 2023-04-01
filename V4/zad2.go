// 2. Napiši program koji

// Kreira 100 goroutina
// Svaka od goroutina "spava" nasumičan broj sekundi od 1-10
// Na kraju spisuje "goroutina broj <i> je završila"
// Program završava nakon što sve goroutine završe i ispisuje "Kraj programa"

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gorutina(i int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	fmt.Println("goroutina broj", i, "je završila")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go gorutina(i, &wg)
	}
	wg.Wait()
	fmt.Println("Kraj programa")
}
