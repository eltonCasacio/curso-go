package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
