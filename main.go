package main

import "fmt"

func sayHello() {
    fmt.Println("Halo Dunia!")
}
func greet(nama string) {
    fmt.Println("Halo", nama)
}

func main() {
    sayHello()
    greet("Aziz")
}
