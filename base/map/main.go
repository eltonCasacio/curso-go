package main

import "fmt"

func main() {
	salarios := map[string]float64{"Elton": 1200, "Roberto": 2500, "Daniel": 8000}
	fmt.Println("Salários")
	for key, value := range salarios {
		fmt.Printf("\t%s: R$%.2f\n", key, value)
	}

	fmt.Printf("\nSalário do Elton: %.2f\n\n", salarios["Elton"])

	delete(salarios, "Daniel")
	for name, salario := range salarios {
		fmt.Printf("\t%s: R$%.2f\n", name, salario)
	}

	salarios["Casacio"] = 12000
	fmt.Printf("Funcionario Casacio adicionado %v\n", salarios)

}
