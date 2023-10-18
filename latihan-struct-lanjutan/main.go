package main

import "fmt"

// Membuat struct orang tua
type Orangtua struct {
	Nama string
	Umur int
}

// Membuat struct anak
type Anak struct {
	Nama     string
	Umur     int
	Kelas	string
	Orangtua Orangtua //Embedding
}

func main() {
	// Membuat slice dari struct orangtua
	dataOrangTua := []Orangtua{
		{Nama: "Budi", Umur: 40},
		{Nama: "Citra", Umur: 39},
		{Nama: "Eva", Umur: 35},
		{Nama: "Rudi", Umur: 40},
		{Nama: "Dewi", Umur: 38},
		{Nama: "Hendro", Umur: 42},
	}

	// Membuat slice dari struct anak
	dataSiswa := []Anak{
		{Nama: "Ali", Umur: 15, Kelas: "9A", Orangtua: dataOrangTua[0]},
		{Nama: "David", Umur: 14, Kelas: "8B", Orangtua: dataOrangTua[1]},
		{Nama: "Fina", Umur: 16, Kelas: "10C", Orangtua: dataOrangTua[2]},
		{Nama: "Eva", Umur: 12, Kelas: "6B", Orangtua: dataOrangTua[3]},
		{Nama: "Faisal", Umur: 11, Kelas: "5A", Orangtua: dataOrangTua[4]},
		{Nama: "Grace", Umur: 10, Kelas: "4C", Orangtua: dataOrangTua[5]},
	}
	
	// Menampilkan informasi siswa
	for i, v := range dataSiswa {
		fmt.Printf("Informasi dataSiswa %d\n", i+1)
		fmt.Printf("Nama: %s, Umur: %d, Kelas: %s\n", v.Nama, v.Umur, v.Kelas)
		fmt.Printf("Orang Tua: %s, Umur: %d\n", v.Orangtua.Nama, v.Orangtua.Umur)
		fmt.Println()
	}

	// Menampilkan daftar siswa
	fmt.Println("Daftar Siswa:")
	for i, v := range dataSiswa[3:] {
		fmt.Printf("Nama Siswa: %s, Umur: %d, Kelas: %s\n", v.Nama, v.Umur, v.Kelas)
		fmt.Printf("Orang Tua: %s, Umur: %d\n", dataOrangTua[i+3].Nama, dataOrangTua[i+3].Umur)
		fmt.Println()
	}
}
