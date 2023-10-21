package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Fibonacci(n int) string {
	// Penanganan kasus khusus jika n adalah 0 atau 1
	if n <= 0 {
		return ""
	} else if n == 1 {
		return "1"
	}

	// Inisialisasi dua bilangan pertama dalam deret Fibonacci
	a, b := 1, 1
	fibonacciSeq := []int{a, b}

	// Membangun deret Fibonacci hingga mencapai panjang yang diinginkan
	for i := 2; i < n; i++ {
		next := a + b
		fibonacciSeq = append(fibonacciSeq, next)
		a, b = b, next
	}

	// Mengonversi deret Fibonacci ke bentuk string
	fibStr := make([]string, len(fibonacciSeq))
	for i, num := range fibonacciSeq {
		fibStr[i] = strconv.Itoa(num)
	}

	// Menggabungkan semua angka dengan koma sebagai pemisah
	return strings.Join(fibStr, ",")
}

func main() {
	n := 10
	fibStr := Fibonacci(n)
	fmt.Println(fibStr)
}
