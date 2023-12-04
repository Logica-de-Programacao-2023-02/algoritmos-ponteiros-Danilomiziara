package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
}

func main() {
	livro1 := Livro{
		titulo: "penguins de madagascar",
		autor:  "Autor 1",
	}

	livro2 := Livro{
		titulo: "se beber não case",
		autor:  "Anônimo",
	}

	fmt.Println("Livro 1 antes:", livro1)
	fmt.Println("Livro 2 antes:", livro2)

	mudarTitulo(&livro1)
	mudarTitulo(&livro2)

	fmt.Println("Livro 1 depois:", livro1)
	fmt.Println("Livro 2 depois:", livro2)
}

func mudarTitulo(livro *Livro) {
	if livro.autor == "Anônimo" {
		livro.titulo = "Desconhecido"
	}
}
