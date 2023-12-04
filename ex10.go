package main

import (
	"fmt"
	"math"
)

func ehPrimo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func preencherComPrimos(slice *[]int, tamanho int) {
	numero := 2
	for len(*slice) < tamanho {
		if ehPrimo(numero) {
			*slice = append(*slice, numero)
		}
		numero++
	}
}

func main() {
	tamanho := 5
	var numerosPrimos []int

	preencherComPrimos(&numerosPrimos, tamanho)

	fmt.Printf("Os %d primeiros números primos são: %v\n", tamanho, numerosPrimos)
}
