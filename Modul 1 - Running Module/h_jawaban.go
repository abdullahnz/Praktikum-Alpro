package main

import (
  "fmt"
)

func hitungVolume(r, t float64) float64 {
  pi := 3.14
  return r * r * pi * t
}

func main() {
  var r, t float64
  fmt.Scan(&r, &t)
  fmt.Println(hitungVolume(r, t))
}
