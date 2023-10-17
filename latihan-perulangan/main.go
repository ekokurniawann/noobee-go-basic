package main

import "fmt"

func main() {
    rows := 5

    for i := rows; i >= 1; i-- {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
	fmt.Println()

	for i := 1; i <= rows; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }

	for i := 1; i <= 50; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Print("NooBee, ")
        } else if i%3 == 0 {
            fmt.Print("Noo, ")
        } else if i%5 == 0 {
            fmt.Print("Bee, ")
        } else {
            fmt.Print(i, ", ")
        }
    }
}
