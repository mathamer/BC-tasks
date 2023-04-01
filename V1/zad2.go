// Kreiraj beskonačnu petlju koja nasumično ispisuje jedan od znakova: , /, |, -

// očekivani izlaz u bilo kojem redosljedu
// \ Please Wait. Processing....
// — Please Wait. Processing....
// \ Please Wait. Processing....
// | Please Wait. Processing...

package main

import "fmt"

func main() {
	var chr = [4]string{"\\", "—", "\\", "|"}

	for {
		for i := 0; i <= 3; i++ {
			fmt.Println(chr[i], "Please Wait. Processing...")
		}
	}
}
