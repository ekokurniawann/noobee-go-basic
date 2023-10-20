package main

import "fmt"

func car(name, color string) {
	fmt.Printf("Mobil %s Berwarna %s\n",name, color)
}

func main() {
	car("BMW", "Black")
}
