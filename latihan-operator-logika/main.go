package main

import "fmt"

func main() {
	gender := "male"
	age := 17

	if gender == "female" {
		if age >= 21 {
			fmt.Println("Pelamar dapat dikerjakan")
		} else {
			fmt.Println("Pelamar tidak dapat dikerjakan")
		}
	} else {
		if age > 18 {
			fmt.Println("Pelamar dapat dikerjakan")
		} else {
			fmt.Println("Pelamar tidak dapat dikerjakan")
		}
	}
}
