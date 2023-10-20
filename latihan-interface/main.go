package main

import "fmt"

// Membuat interface
type Arithmetic interface {
    Add() float64
    Subtract() float64 // Tambahkan method Subtract
    Multiply() float64
    Divide() (float64, error)
}

// Memuat struct 
type Num struct {
    a, b float64
}

// Membuat constructor function
func NewNum(a, b float64) *Num {
    return &Num{a, b}
}

func (n Num) Add() float64 {
    return n.a + n.b
}

func (n Num) Subtract() float64 {
    return n.a - n.b
}

func (n Num) Multiply() float64 {
    return n.a * n.b
}

func (n Num) Divide() (float64, error) {
    if n.b == 0 {
        return 0, fmt.Errorf("Tidak bisa membagi dengan nol")
    }
    return n.a / n.b, nil
}

func main() {
    num := NewNum(4.0, 2.0)
    
    // Menggunakan fungsi Substract
	hasilSubstract := num.Subtract()
	fmt.Printf("Hasil pengurangan: %v\n", hasilSubstract)

	// Menggunakan fungsi Divide
    // hasilDivide, err:= num.Divide()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
    // fmt.Printf("Hasil pembagian: %v\n", hasilDivide)
}
