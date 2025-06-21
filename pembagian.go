package main

import "fmt"

func bagi(a float64, b float64) float64 {
    if b == 0 {
        fmt.Println("Error: pembagian dengan nol")
        return 0
    }
    return a / b
}

func main() {
    fmt.Println("9 / 3 =", bagi(9, 3))
}
