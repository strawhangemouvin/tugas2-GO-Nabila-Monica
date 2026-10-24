package main

import "fmt"

func main() {
	panjang := 8
	lebar := 5

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	isLebihDari30 := luas > 30

	fmt.Printf("Panjang : %d\n", panjang)
	fmt.Printf("Lebar   : %d\n", lebar)
	fmt.Printf("Luas    : %d\n", luas)
	fmt.Printf("Keliling: %d\n", keliling)
	fmt.Printf("Luas > 30?: %t\n", isLebihDari30)
}
