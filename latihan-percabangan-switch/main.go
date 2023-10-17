// Latihan
/**
Buatlah program Go untuk melakukan pengecekan sebuah variable apakah huruf tersebut adalah huruf vokal atau konsonan menggunakan percabangan switch.
*/

package main

import (
	"fmt"
)

func main() {
	// input user manual
	var huruf string
	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	switch huruf {
	case "a", "A", "i", "I", "u", "U", "e", "E", "o", "O":
        fmt.Printf("Huruf %v adalah huruf vokal", huruf)
    case "b", "B", "c", "C", "d", "D", "f", "F", "g", "G", "h", "H", "j", "J", "k", "K", "l", "L", "m", "M", "n", "N", "p", "P", "q", "Q", "r", "R", "s", "S", "t", "T", "v", "V", "w", "W", "x", "X", "y", "Y", "z", "Z":
        fmt.Printf("Huruf %v adalah huruf konsonan", huruf)
	default:
		fmt.Printf("Inputan tidak valid")
	}
}
