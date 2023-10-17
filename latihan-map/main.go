package main

import "fmt"

func main() {
	var contact map[string]int 
	contact = map[string]int{}

	contact["Alice"] = 1234567890
	contact["Bob"] = 9876543210

	fmt.Printf("Semua kontak %v\n", contact)
	fmt.Printf("Nomor telepon Alice: %v\n", contact["Alice"])
	fmt.Println()

	contact["Charlie"] = 5555555555
	fmt.Printf("Setelah menambah kontak Charlie: %v\n", contact)		
	fmt.Println()

	delete(contact, "Bob")
	fmt.Printf("Setelah menghapus kontak Bob: %v\n", contact)
	fmt.Println()

	fmt.Println("Data kontak:")
	for key, val := range contact {
		fmt.Println(key, ":", val)
	}

}
