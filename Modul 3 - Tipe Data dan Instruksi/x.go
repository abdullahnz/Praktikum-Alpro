package main

import (
	"fmt"
)

func main() {
	var nilai int
	var npwp, covid bool

	var pajak int 

	fmt.Scan(&nilai, &npwp, &covid)

	if nilai >= 1e6 {
		pajak = 10 // ppn
		if !covid {
			pajak += 10 // bea
			if npwp {
				pajak += 10 // pph
			} else {
				pajak += 20 // pph
			}
		}
	}

	fmt.Printf("pajak = %d%%", pajak)
}