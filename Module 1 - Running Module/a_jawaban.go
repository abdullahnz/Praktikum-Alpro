package main

import (
	"fmt"
)

func main() {
	var a, b, hasil_penjumlahan int
	var s string
	fmt.Scan(&s, &a, &b)
	hasil_penjumlahan = a + b
	fmt.Println("Kata =", s)
	fmt.Println("Jumlah =", hasil_penjumlahan)
}
