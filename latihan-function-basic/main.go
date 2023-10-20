package main

import "fmt"

type Car struct {
	data map[string]string
}


// Cara pertama dengan menggunakan dua fungsi
func prosesCar(car map[string]string) string {
	name := car["name"]
	color := car["color"]
	result := fmt.Sprintf("Mobil %s Berwarna %s\n", name, color)
	return result
}

func printCar(message string) {
	fmt.Println(message)
}

// Cara kedua menggunakan satu fungsi
func prosesCar1() string {
	car := make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"
	deskripsi := fmt.Sprintf("Mobil %s Berwarna %s\n", car["name"], car["color"])
	return deskripsi
}

// Cara ketiga menggunakan anonymous function
var printCar2 = func() string {
	car := make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"
	deskripsi := fmt.Sprintf("Mobil %s Berwarna %s\n", car["name"], car["color"])
	return deskripsi
}

// Cara ke empat dengan membuat method
func (c *Car) printCar3() string {
	name := c.data["name"]
	color := c.data["color"]
	deskripsi := fmt.Sprintf("Mobil %s Berwarna %s\n", name, color)
	return deskripsi
	
}

func main() {
	// Implementasi struct untuk cara pertama
	car := map[string]string{
		"name": "BMW",
		"color": "Black",
	}

	// Mencetak Cara pertama
	message := prosesCar(car)
	printCar(message)

	// Mencetak Cara kedua
	message1 := prosesCar1()
	fmt.Println(message1)

	// Mencetak Cara ketiga
	message2 := printCar2()
	fmt.Println(message2)

	// Cara ke empat 
	// Implementasi struct untuk cara ke empat
	car1 := Car{data: map[string]string{"name": "BMW", "color": "Black"}}
	
	message3 :=car1.printCar3()
	fmt.Println(message3)
}
