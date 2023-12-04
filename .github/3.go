package main

import "fmt"

func main() {
	s := "The witcher 3"
	fmt.Printf("Antes: %s\n", s)
	reverse(&s)
	fmt.Printf("Depois: %s\n", s)
}

func reverse(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}
