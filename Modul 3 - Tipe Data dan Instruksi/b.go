package main 

import (
	"fmt"
)

func main() {
	var a, b, c string
	
	fmt.Scan(&a, &b, &c)

	ans := (a == b) || (a == c) || (b == c)
	
	fmt.Println(ans) 
}	