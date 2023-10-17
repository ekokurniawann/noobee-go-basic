package main

import "fmt"

func main() {
	//array 
	animals := [...]string{"Cat", "Dog", "Pinguin", "Chicken", "Snake"}

	mammals := animals[0:2]
	notMammals := animals[2:5]
	havelegs := animals[0:4]

	fmt.Println(mammals)
	fmt.Println(notMammals)
	fmt.Println(havelegs)
	fmt.Println()

	animals[1] = "Cow"
	animals[2] = "Bird"

	fmt.Println(mammals)
	fmt.Println(notMammals)
	fmt.Println(havelegs)
}
