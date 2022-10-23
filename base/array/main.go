package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for i := 0; i < len(meuArray); i++ {
		fmt.Println(meuArray[i])
	}

	meuArray2 := []int{2, 4, 6, 8}
	fmt.Println(meuArray2)

}
