package main

import "fmt"

func atualizarParaParOuImpar(ptr *int) {
    if *ptr%2 == 0 {
        *ptr = 0
    } else {
        *ptr = 1
    }
}

func main() {
    valor := 7
    ptr := &valor

    atualizarParaParOuImpar(ptr)

    fmt.Printf("O valor atualizado é: %d\n", valor)
}
