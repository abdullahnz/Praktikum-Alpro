package main

import (
  "fmt"
)

func main() {
  var string string
  fmt.Scan(&string)
  for string != "selesai" {
    fmt.Println(string)
    fmt.Scan(&string)
  }
}
