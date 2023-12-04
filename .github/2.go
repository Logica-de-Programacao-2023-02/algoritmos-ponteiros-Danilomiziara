package main

import (
	"fmt"
)

func verificarParImpar(num *int) {
	if *num%2 == 0 {
		*num = 0
	} else {
		*num = 1
	}
}

func main() {
	numero := 7

	fmt.Println("Número original:", numero)

	verificarParImpar(&numero)

	fmt.Println("Número atualizado:", numero)
}
