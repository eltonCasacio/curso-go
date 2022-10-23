package main

import "fmt"

func main() {
	var minhaVar interface{}
	minhaVar = "Elton"
	println(minhaVar.(string))

	value, ok := minhaVar.(int)
	fmt.Printf("O valor é %v e o ok é %v\n", value, ok)

}
