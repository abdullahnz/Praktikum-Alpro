package main

import (
  "fmt"
)

func main() {
  var n, t1, t2, t3, jumlah int
  fmt.Scan(&n)
  jumlah = 0
  for i := 0; i < n; i++ {
    fmt.Scan(&t1, &t2, &t3)
    if t1 + t2 + t3 >= 2 {
      jumlah = jumlah + 1
    }
  }
  fmt.Println(jumlah)
}
