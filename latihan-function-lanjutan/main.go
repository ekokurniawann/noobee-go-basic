package main

import "fmt"

func hitungRataRata(nilaiSiswa ...float32) float32 {
    total := float32(0)
	// jumlahNilai := 0

    for _, nilai := range nilaiSiswa {
        total += nilai
		// jumlahNilai++
    }

	// rataRata := total / float32(jumlahNilai)
    // return rataRata, jumlahNilai

    return total / float32(len(nilaiSiswa))
}

func main() {
    nilaiSiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}

	// rataRata, jumlahNilai := hitungRataRata(nilaiSiswa...)
    rataRata := hitungRataRata(nilaiSiswa...)
	// rataRata := hitungRataRata(nilaiSiswa[0], nilaiSiswa[1], nilaiSiswa[2], nilaiSiswa[3], nilaiSiswa[4])


	// fmt.Printf("Jumlah nilai siswa dalam kelas: %d\n", jumlahNilai)
    fmt.Printf("Jumlah nilai siswa dalam kelas: %d\n", len(nilaiSiswa))
    fmt.Printf("Nilai rata-rata siswa dalam kelas: %.2f\n", rataRata)
}
