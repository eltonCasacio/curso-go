package main

import "fmt"

const a = "Hello Human"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Hello"
	e float64 = 1.2
	f ID      = 1
)

func main() {

	whatType := 12.343453453453454
	fmt.Printf("whatType é %T\n", whatType)

	fmt.Printf("O tipo de A é %T\n", a)
	fmt.Printf("O tipo de B é %T\n", b)
	fmt.Printf("O tipo de C é %T\n", c)
	fmt.Printf("O tipo de D é %T\n", d)
	fmt.Printf("O tipo de E é %T\n", e)
	fmt.Printf("O tipo de F é %T\n", f)
}
