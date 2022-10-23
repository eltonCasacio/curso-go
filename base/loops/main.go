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

	for k, v := range meuArray {
		fmt.Println(k, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//loop infinit
	// for {
	// 	fmt.Println(i)
	// }

}
