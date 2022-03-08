package main 

import (
	"fmt"
)

func main() {
	var tahun int

	ans := true
	for i := 0; i < 4; i++ { 
		fmt.Scan(&tahun)
		
		if !(tahun % 400 == 0 || tahun % 4 == 0 && tahun % 100 != 0) {
			ans = false
		}
	}

	fmt.Println(ans)
}	