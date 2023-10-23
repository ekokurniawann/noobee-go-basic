package main

import (
	"fmt"
	"os"
)

// fungsi cetak
func cetak(text string) {
	fmt.Println(text)
}

func main() {
	// Defer
	// defer cetak("text 1")
	// defer cetak("text 2")
	// cetak("text 3")
	// defer cetak("text 4")
	// cetak("text 5")

	// Contoh lainnya
	num1 := 5

	if num1 % 2 > 0 {
		cetak("Ini adalah bilangan ganjil")
		// defer cetak("akan dieksekusi setelah proses di atas selesai")
		func() {
			defer cetak("ini akan dieksekusi ketika proses di atas selesai")
		}()
	}

	cetak("Oh tentu tidak, Defer akan di eksekusi setelah ini")

	// Exit
	names := []string{"Noob", "Bee", "Go", "NodeJS"}

	for _, name := range names {
		if name == "Go" {
			os.Exit(1)
		}
		cetak(name)
	}
}
