package main

import "fmt"

func kurang(a int, b int) int {
    return a - b
}

func main() {
    fmt.Println("5 - 3 =", kurang(5, 3))
}
