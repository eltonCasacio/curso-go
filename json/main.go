package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 123,
		Saldo:  243700,
	}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	var conta2 Conta
	jsonPuro := []byte(`{"n":432,"s":0}`)
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Conta 2: %v\n", conta2)
}
