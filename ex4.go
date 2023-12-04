package main

import "fmt"

func atualizarParaSomaUltimosDigitos(ptr *int) {
    valor := *ptr

    ultimoDigito := valor % 10
    valor /= 10
    penultimoDigito := valor % 10
    soma := ultimoDigito + penultimoDigito

    *ptr = soma
}

func main() {
    numero := 1234
    ptr := &numero

    atualizarParaSomaUltimosDigitos(ptr)

    fmt.Printf("O valor atualizado é: %d\n", numero)
}
