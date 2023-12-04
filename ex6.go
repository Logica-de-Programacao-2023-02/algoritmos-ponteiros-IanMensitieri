package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func alterarTituloSeAnonimo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livroExemplo := Livro{
		Titulo: "O Livro Misterioso",
		Autor:  "Anônimo",
	}

	fmt.Printf("Antes: Título: %s, Autor: %s\n", livroExemplo.Titulo, livroExemplo.Autor)

	alterarTituloSeAnonimo(&livroExemplo)

	fmt.Printf("Depois: Título: %s, Autor: %s\n", livroExemplo.Titulo, livroExemplo.Autor)
}
