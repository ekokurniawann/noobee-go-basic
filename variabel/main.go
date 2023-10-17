package main

import "fmt"

var x int = 10

func main() {
	// Membuat variabel secara explisit
	var myName string = "NooBeeID"

	var yourName string 
	yourName = "Eko"

	var nickName = "NooBee"

	myAge := 22 

	fmt.Println(myName)
	fmt.Println(yourName)
	fmt.Println(nickName)
	fmt.Println(myAge)

	// Variabel di luar fungsi
	fmt.Println(x)

	// Deklarasi variabel tanpa nilai awal
	var a string
	var b int 
	var c bool

	fmt.Println(a)
	fmt.Println(b)	
	fmt.Println(c)

	// Deklarasi multiple variabel
	var d, e, f, g  = 1, 2, 3, "Hello"

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Deklarasi multiple variabel dalam blok
	var (
		firstName string = "NooBee"
		lastName string = "ID"
		height int = 170
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(height)

}
