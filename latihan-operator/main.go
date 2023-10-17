// Latihan
package main

import "fmt"

func main() {
	// Sekarang
	umurSaya, umurKakak := 20, 30

	fmt.Printf("Umur saya sekarang adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak sekarang adalah : %d\n", umurKakak)
	fmt.Println()

	// Tahun depan
	umurSaya +=1
	umurKakak +=1

	fmt.Printf("Umur saya tahun depan adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak tahun depan adalah : %d\n", umurKakak)
	fmt.Println()

	// 5 tahun lagi
	umurSaya +=5
	umurKakak +=5

	fmt.Printf("Umur saya 5 tahun lagi adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak 5 tahun lagi adalah : %d\n", umurKakak)
}
