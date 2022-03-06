# Praktikum Algoritma Pemrograman

Repository ini akan berisi materi maupun cheatsheet tentang syntax basic golang yang digunakan dalam materi perkuliahan Algoritma Pemrograman di Telkom University.

## Golang Installation

Langkah-langkah lengkap bisa dibaca di dokumentasi golang.

### Linux

TODO

### Windows

TODO

### Mac

TODO

## Go Basic Commands

Command lengkap dari golang dapat dilihat dengan mengeksekusi command `go`. Berikut informasi lengkapnya.

```sh
$ go
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

Lalu, untuk melihat option atau cara penggunaan dari suatu command dapat menambahkan parameter `help` dengan format `go help <command>`. Contoh sebagai berikut.

```sh
$ go help run
$ go help build
$ go help version
```

### Golang File Execution

Untuk meneksekusi file golang dengan cara sederhana, dapat dilakukan dengan command `run`. Sedikit penjelasan tentang command ini, command ini melakukan build pada temporary memory atau RAM.

```sh
$ go run <namafile.go>
```

Contoh

```sh
$ go run swap.go
$ go run sum.go
```

### Build Executable File

Executable file adalah sebuah binary file yang tidak memerlukan interpreter untuk mengeksekusi file itu sendiri. Untuk membuat sebuah file executable kita perlu mengeksekusi command `build`. Command `build` sendiri memerlukan sebuah parameter nama file golang dan akan menghasilkan sebuah file executable jika tidak terjadi error.

```sh
$ go build <namafile.go>
```

Misalkan kita mempunyai file `swap.go` dan akan melakukan kompilasi pada file tersebut. Maka setelah menjalankan build akan menghasilkan binary file baru.

```sh
$ ls -l
total 1
-rw-r--r-- 1 nobody 197609     235 Mar  5 18:51 swap.go

$ go build swap.go

$ ls -l
total 1885
-rwxr-xr-x 1 nobody 197609 1927168 Mar  5 18:55 swap.exe*
-rw-r--r-- 1 nobody 197609     235 Mar  5 18:51 swap.go

$ ./swap
<program running>
```

Note, command `ls` adalah command di `Unix` dan `Unix-like` yang digunakan untuk melihat isi direktori. Untuk di `Windows` dapat menggunakan command `dir`.

### Whats different between run and build ?

Pada dasarnya cara kerjanya sama, perbedaannya adalah `run` akan mengkompilasi (build) file golang pada temporary directory *atau* RAM lalu mengeksekusinya. Sedangkan `build` akan mengkompilasi pada direktori saat ini dan akan menghasilkan file binary.

Penggunaan `run` sangat disarankan (penulis) apabila program hanya merupakan program sederhana khususnya pada perkuliahan di mata kuliah ini, karena dinilai lebih mempersingkat waktu pada saat testing program.

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

<!-- HOW STRICT IS GOLANG? -->