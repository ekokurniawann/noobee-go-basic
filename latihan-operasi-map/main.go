package main

import "fmt"

func main() {
	// Membuat map fruits
	fruits := map[string]int{
		"Apple":  10,
		"Banana": 15,
		"Grape":  20,
		"Orange": 8,
	}
	fmt.Println("Sebelum ditambah/hapus:")
	fmt.Println(fruits)	
	fmt.Println()

	// Menambahkan data manggo dan strawberry ke map fruits
	fruits["Manggo"] = 12
	fruits["Strawberry"] =7
	fmt.Println("Setelah menambahkan Manggo dan Strawberry:")
	fmt.Println(fruits)
	fmt.Println()

	// Menghapus data orange dari map fruits
	delete(fruits, "Orange")
	fmt.Println(fruits)
	fmt.Println()

	// Mencetak jumlah data appel dari map fruits
	fmt.Println("Jumlah apel yang tersedia adalah:", fruits["Apple"])
	fmt.Println()

	// Mencetak daftar data buah-buahan dari map fruits beserta jumlahnya
	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	for key, value := range fruits {
		fmt.Println(key, ":", value)
	}
}
