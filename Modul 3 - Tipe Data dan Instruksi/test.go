package main

import "fmt" 

func solve(base int, cek int) {
	bil := base
	digits := 1
	
	ans := false
	for bil > 0 {
		for base > 0 {
			if cek == base {
				ans = true
				break
			}
			base = base / 10
			if base > 0 {
				digits = digits * 10
			}
		}

		if ans {
			break
		}
		bil = bil % (digits)
		base = bil
		digits = 1
	}
	fmt.Println(ans)
}

func main() {
	// bil := make([]float64, 4)
	// max := -1.0

	// for i := 0; i < 4; i++ {
	// 	fmt.Scan(&bil[i])
	// 	if bil[i] > max {
	// 		max = bil[i]
	// 	}
	// }

	// for i := 0; i < 4; i++ {
	// 	fmt.Println("Bola", i, "terberat?", bil[i] == max)
	// }

	var a, b, c int

	fmt.Scan(&a, &b, &c)

	solve(a, b)
	solve(a, c)
	 
}