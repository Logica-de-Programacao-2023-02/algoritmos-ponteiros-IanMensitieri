package main

import "fmt"

func atualizarParaSomaNaturais(ptr *int, n int) {
    soma := 0
    for i := 1; i <= n; i++ {
        soma += i
    }
    *ptr = soma
}

func main() {
    valor := 0
    n := 5
    ptr := &valor

    atualizarParaSomaNaturais(ptr, n)

    fmt.Printf("O valor atualizado é: %d\n", valor)
}
