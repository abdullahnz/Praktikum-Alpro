package main 

import (
	"fmt"
)

func main() {
	var a, b, c, d int 
	var result float64

	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b, &c, &d)
	
		result = (float64(b) - float64(d)) / (float64(a) - float64(c))
	
		fmt.Println(result)
	}
}	