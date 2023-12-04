package main

import "fmt"

func main() {
	var numero int = 10
	ponteiro := &numero
	fmt.Println("Valor antes da modificação:", numero)
	mudarValor(ponteiro)
	fmt.Println("Valor depois da modificação:", numero)
}
func mudarValor(ponteiro *int) {
	*ponteiro = 42
}
