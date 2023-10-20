package main

import (
	"fmt"
	"latihan-package-access-modifier/calculator"
)

func main() {
    // Menggunakan fungsi Add dari package calculator
    result := calculator.Add(5, 5)
    fmt.Println("Hasil penjumlahan:", result)
}
