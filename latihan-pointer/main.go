package main

import "fmt"

type Student struct {
    Name  string
    Class string
}

func (s *Student) SetMyName(name string) {
    s.Name = name
}

func (s Student) CallMyName() {
    fmt.Printf("Hello, My Name is %s\n", s.Name)
}

func main() {
    student := Student{
        Name:  "John",
        Class: "10A",
    }

    // Memanggil method SetMyName untuk mengubah nama
    student.SetMyName("Alice")

    // Memanggil method CallMyName untuk menampilkan pesan
    student.CallMyName()
}
