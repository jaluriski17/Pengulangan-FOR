package main

import "fmt"

func main() {
    // Contoh pengulangan for
    fmt.Println("Pengulangan menggunakan for:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Contoh pengulangan while
    fmt.Println("\nPengulangan menggunakan while:")
    j := 1
    for j <= 5 {
        fmt.Println(j)
        j++
    }

    // Contoh pengulangan menggunakan range pada slice
    fmt.Println("\nPengulangan menggunakan range:")
    buah := []string{"apel", "jeruk", "pisang"}
    for index, nama := range buah {
        fmt.Printf("Index %d: %s\n", index, nama)
    }
}
