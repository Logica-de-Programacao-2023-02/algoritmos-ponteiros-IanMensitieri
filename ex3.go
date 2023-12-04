package main

import "fmt"

func inverterString(ptr *string) {
    original := *ptr
    reversed := ""
    for i := len(original) - 1; i >= 0; i-- {
        reversed += string(original[i])
    }
    *ptr = reversed
}

func main() {
    texto := "Go é incrível!"
    ptr := &texto

    inverterString(ptr)

    fmt.Printf("A string invertida é: %s\n", texto)
}
