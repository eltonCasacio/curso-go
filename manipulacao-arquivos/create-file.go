package main

import (
	"fmt"
	"os"
)

func CreateFile() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// tamanho, err := f.Write([]byte("aiai"))
	tamanho, err := f.WriteString("Meu arquivo texto criado em GO")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho do arquivo criado: %d\n", tamanho)
}
