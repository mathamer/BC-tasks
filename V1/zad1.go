// Sumiraj parne brojeve između 1 i 1000
// očekivani izlaz: 2 + 4 + 6 + 8 + 10 + … + 1000 = ...

package main

import "fmt"

func main() {
	var suma = 0

	for i := 2; i <= 1000; i += 2 {
		suma += i
		if i < 1000 {
			fmt.Print(i, " + ")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println(" = ", suma)
}
