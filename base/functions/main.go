package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := soma(0, 10)
	if err != nil {
		fmt.Println("Digite algum valor maior que zero")
		return
	}

	//Closure - anonimous function
	total := func() int {
		return somaVariadic(2, 2, 2)
	}()

	fmt.Println(total)
}

func soma(a int, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, errors.New("Invalid value")
	}
	return a + b, nil
}

func somaVariadic(numeros ...int) int {
	var sum int
	for _, numero := range numeros {
		sum = sum + numero
	}

	return sum
}
