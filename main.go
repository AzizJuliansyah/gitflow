package main

import "fmt"

func sayHello() {
    fmt.Println("Halo Dunia!")
}
func greet(nama string) {
    fmt.Println("Halo", nama)
}
func tambah(a int, b int) int {
    return a + b
}
func bagi(a int, b int) (int, int) {
    return a / b, a % b
}



func main() {
    sayHello()
    greet("Aziz")

    hasil := tambah(3, 5)
    fmt.Println("Hasil tambah:", hasil)

    hasilBagi, sisa := bagi(10, 3)
    fmt.Println("Hasil bagi:", hasilBagi)
    fmt.Println("Sisa bagi:", sisa)
}
