package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pathFile := "arquivo.txt"
	f, err := os.Create(pathFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write([]byte("Hello Human"))
	if err != nil {
		panic(err)
	}

	fileBytes, err := os.ReadFile(pathFile)
	if err != nil {
		panic(err)
	}

	teste := []string{string(fileBytes), "teste"}
	res := strings.Join(teste, " ")
	fmt.Println(res)
}
