package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func main() {
	livro := Livro{Titulo: "the witcher", Autor: "Andrzej Sapkowski", Preco: 52.0}
	fmt.Printf("Antes: Preço = %.2f\n", livro.Preco)
	Desconto(&livro, 10)
	fmt.Printf("Depois: Preço = %.2f\n", livro.Preco)
}
func Desconto(l *Livro, desconto float64) {
	l.Preco *= 1 - desconto/100
}
