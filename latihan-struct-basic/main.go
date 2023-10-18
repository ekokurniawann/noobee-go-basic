package main

import "fmt"

type Book struct {
	Judul string
	Penulis string
}


func main() {
	// Menginisialisasi struct secara literal
	book := Book{
		Judul:   "Pemrograman Golang",
		Penulis: "John Doe",
	}

	// Menginisialisasi Struct Saat Membuat Variabel
	var book1 Book 
	book1.Judul = "Algoritma dan Struktur Data"
	book1.Penulis = "Jane smith"

	var book2 Book 
	book2.Judul = "Machine Learning for Beginners"
	book2.Penulis = "David Johnson"


	// Mengakses dan mencetak nilai dari struct 
	fmt.Println("Informasi tentang Book 1:")
	fmt.Println("Judul: ", book.Judul)
	fmt.Println("Penulis: ", book.Penulis)
	fmt.Println()

	fmt.Println("Informasi tentang Book 2:")
	fmt.Println("Judul: ", book1.Judul)
	fmt.Println("Penulis: ", book1.Penulis)
	fmt.Println()

	fmt.Println("Informasi tentang Book 3:")
	fmt.Println("Judul: ", book2.Judul)
	fmt.Println("Penulis: ", book2.Penulis)
	fmt.Println()

	// Cara Kedua
	// Pertama buat slice dari struct Book
	books := []Book{
		{Judul: "Pemrograman golang", Penulis: "John Doe"},
		{Judul: "Algoritma dan Struktur Data", Penulis: "Jane Smith"},
		{Judul: "Machine Learning for Beginners", Penulis: "David Johnson"},
	}
	
	for i, book := range books {
		fmt.Println("Informasi tentang Book", i+1)
		fmt.Println("Judul: ", book.Judul)
		fmt.Println("Penulis: ", book.Penulis)
		fmt.Println()
	}

}
