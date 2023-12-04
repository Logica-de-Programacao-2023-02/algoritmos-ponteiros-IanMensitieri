package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarAoSaldo(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	contaExemplo := Conta{
		Saldo:   1000.0,
		Titular: "João",
	}

	fmt.Printf("Saldo antes: %.2f\n", contaExemplo.Saldo)

	adicionarAoSaldo(&contaExemplo, 500.0)

	fmt.Printf("Saldo depois: %.2f\n", contaExemplo.Saldo)
}
