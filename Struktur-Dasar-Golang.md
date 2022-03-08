## Golang Basic Structure

Program "Hello, World!" sederhana.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

Penjelasan mengenai program diatas,

- `package` adalah keyword go yang mendefinisikan bundle kode mana yang dimiliki file ini. Hanya ada satu `package` per folder, dan setiap `.go` file harus mendeklarasikan nama `package` yang sama di bagian atas filenya. Dalam contoh ini, kode diatas dimiliki oleh `main`.
- `import` adalah keyword go yang memberitahu kepada go compiler `package` mana saja yang akan digunakan (include) pada file ini. Dalam kasus ini, kita mengimport `fmt` package yang merupakan standart library dari golang.
- `func main()` adalah fungsi utama yang akan dieksekusi oleh program ketika program berjalan.
- `fmt.Println` adalah sebuah fungsi yang terdapat pada `fmt` package yang akan mencetak suatu nilai ke screen ketika program berjalan.
