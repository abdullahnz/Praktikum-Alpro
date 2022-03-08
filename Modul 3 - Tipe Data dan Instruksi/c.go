package main 

import (
	"fmt"
)

func main() {
	var n, kg, pon, ons float64

	for i := 0; i < 4; i++ {
		fmt.Scan(&n)

		kg  = n / 1000
		pon = n / 453.592
		ons = n / 28.3495
		
		fmt.Println(kg, pon, ons)
	}
}	