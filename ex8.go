package main

import "fmt"

func modificarValor(ponteiro *int) {
    *ponteiro = 42
}

func main() {
    variavel := 10
    ponteiro := &variavel

    fmt.Printf("Valor antes da modificação: %d\n", variavel)

    modificarValor(ponteiro)

    fmt.Printf("Valor depois da modificação: %d\n", variavel)
}
