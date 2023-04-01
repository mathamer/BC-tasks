package main
  
import (
    "fmt"
    "math/rand"
)
  
func main() {

	var niz [9]int

	for i := 1; i <= rand.Intn(9); i++ {
		niz[i] = rand.Intn(9)
	}
	fmt.Println(niz)
}