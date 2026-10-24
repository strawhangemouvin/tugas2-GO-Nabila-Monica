package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Println("Sebelum", a, b)

	a, b = b, a

	fmt.Println("Sesudah", a, b)
}
