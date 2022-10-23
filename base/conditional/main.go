package main

func main() {
	a := 10
	b := 20
	c := 0

	if a > b && c > a {
		println("a > b && c > a")
	} else if a < b && c < b {
		println("a < b && c < b")
	} else {
		println("else")
	}

	switch c {
	case 0:
		println("zero")
	default:
		println("d")
	}
}
