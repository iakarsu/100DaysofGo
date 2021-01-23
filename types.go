package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	num1, num2 := 32.5, 232.2
	fmt.Println(add(num1, num2))

	a, b := "a", "b"
	fmt.Println(multiple(a, b))

	x := 15
	y := &x
	fmt.Println(y)
	fmt.Println(**y)

	*y = 12
	fmt.Println(x)
}
