package main

import "fmt"

type Number interface {
	int | float64
}

func soma[T Number](nums map[string]T) T {
	var sum T
	for _, value := range nums {
		sum += value
	}
	return sum
}

func comparar[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	var nums = map[string]float64{
		"one": 1.0,
		"two": 2.5,
	}
	fmt.Printf("%.2f\n", soma(nums))

	println(comparar("E", "E"))
	println(comparar(1, 1))
}
