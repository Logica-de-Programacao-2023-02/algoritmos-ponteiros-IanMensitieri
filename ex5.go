package main

import (
	"fmt"
	"math"
)

func atualizarParaMediaComPi(ptr *float64) {
	valor := *ptr

	media := (valor + math.Pi) / 2

	*ptr = media
}

func main() {
	numero := 3.14
	ptr := &numero

	atualizarParaMediaComPi(ptr)

	fmt.Printf("O valor atualizado é: %.2f\n", numero)
}
