package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	a := 10
	ponteiro := &a
	b := ponteiro
	*b = 20

	value1 := 10
	value2 := 20
	total := soma(&value1, &value2)
	println(total)
	println(value1)
}
