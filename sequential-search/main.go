package main

import "log"

func main() {
	// setup angka yang dicari
	find := 2

	// setup list angka
	nums := []int{4, 2, 1, 2, 4, 8, 5, 10, 4, 9, 4}

	// pangil function sequential
	found, index := Sequential(find, nums...)
	log.Println(found, index)
}

/**
function ini berfungsi untuk melakukan algoritma sequential untuk pencarian data
*/

func Sequential(find int, nums ...int) (found bool, index int) {
	// lakukan looping terhadap seluruh isi array
	for i, num := range nums {
		// lakukan perbandingan, apakah ada yang
		// sama dengan data yang dicari atau engga

		if num == find {
			found = true
			index = i
			return
		}
	}
	// kalau gaada, index bakal diset as -1
	index = -1
	return
}
