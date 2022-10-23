package main

import "os"

func main() {
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	println(string(arquivo))
}
