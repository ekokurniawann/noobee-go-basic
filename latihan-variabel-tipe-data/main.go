// Latihan
/*
Buatlah program yang menghitung total harga belanja dengan menghitung harga produk A (10.000), produk B (15.000), dan produk C (7.000) yang dibeli oleh pengguna. Simpan total harga dalam sebuah variabel.
*/

package main

import "fmt"

func main() {
	// cara pertama input dari user
	var a int
	fmt.Print("Masukkan prodcutA: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Masukkan prodcutB: ")
	fmt.Scan(&b)

	var c int
	fmt.Print("Masukkan prodcutC: ")
	fmt.Scan(&c)

	totalPrice := a + b + c
	fmt.Printf("Total harga belanja: %d\n",totalPrice)
	fmt.Println()
	

	// cara kedua
	var productA, productB, productC int =  10000, 15000, 7000

	totalPriceOne := productA + productB + productC
	fmt.Printf("Total harga belanja: %d\n",totalPriceOne)
	fmt.Println()


	// Cara ketiga
	var products = map[string]int{
		"A":10000, 
		"B":15000, 
		"C":7000,
	}

	totalPriceTwo := products["A"] + products["B"] + products["C"]
	fmt.Printf("Total harga belanja: %d\n",totalPriceTwo)

}
