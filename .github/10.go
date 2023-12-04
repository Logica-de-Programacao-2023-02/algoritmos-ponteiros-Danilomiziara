package main

import "fmt"

func main() {
	var primos []int
	tamanho := 50

	fillPrimo(&primos, tamanho)

	fmt.Println("Primeiros", tamanho, "números primos:", primos)
}

func Primo(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func fillPrimo(slice *[]int, size int) {
	count := 0
	num := 2

	for count < size {
		if Primo(num) {
			*slice = append(*slice, num)
			count++
		}
		num++
	}
}
