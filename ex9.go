package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := 0.10
	livro.Preco = livro.Preco - (livro.Preco * desconto)
}

func main() {
	livroExemplo := Livro{
		Titulo: "Go Programming",
		Autor:  "John Doe",
		Preco:  50.0,
	}

	fmt.Printf("Preço antes do desconto: R$%.2f\n", livroExemplo.Preco)

	aplicarDesconto(&livroExemplo)

	fmt.Printf("Preço depois do desconto: R$%.2f\n", livroExemplo.Preco)
}
