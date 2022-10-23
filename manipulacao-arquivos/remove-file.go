package main

import "os"

func main() {
	f, err := os.Create("remove.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("Este arquivo será removido"))
	err = os.Remove("remove.txt")
	if err != nil {
		panic(err)
	}

}
