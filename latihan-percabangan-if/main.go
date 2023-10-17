// Latihan
/**
Buatlah program Go untuk melakukan pengecekan apakah suatu variable bernilai ganjil atau genap menggunakan percabangan if.
*/

package main

import "fmt"

func main() {
	// input manual user
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	if num % 2 != 0 {
		fmt.Printf("%d Merupakan bilangan ganjil\n", num)
	} else {
		fmt.Printf("%d Merupakan bilangan genap\n", num)
	}

}	
