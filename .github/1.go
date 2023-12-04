package main

import "fmt"

func main() {
	x := 0
	fmt.Printf("Antes: x = %d\n", x)
	somaNaturais(&x, 100)
	fmt.Printf("Depois: x = %d\n", x)
}

func somaNaturais(n *int, valor int) {
	soma := 0
	for i := 1; i <= valor; i++ {
		soma += i
	}
	*n = soma
}
