package main

import (
	"fmt"
	"math"
)

func mediaPi(x *float64) {
	*x = (*x + math.Pi) / 2
}

func main() {
	x := 3.0
	fmt.Printf("Antes: x = %.2f\n", x)
	mediaPi(&x)
	fmt.Printf("Depois: x = %.2f\n", x)
}
