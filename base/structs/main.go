package main

import (
	"errors"
	"fmt"
)

type Endereco struct {
	Logradouro  string
	Numero      string
	Complemento string
	Bairro      string
}

type Client struct {
	Name     string
	Age      int
	Ativo    bool
	Endereco // Composicao
}

func (c *Client) Deactivate() {
	c.Ativo = false
}

func (c *Client) Activate(client *Client) error {
	if client.Name == "" {
		return errors.New("Name is required")
	}

	c.Ativo = true

	return nil
}

func main() {
	client := Client{
		Name:  "Elton",
		Age:   35,
		Ativo: true,
	}

	fmt.Println(client)

	client.Deactivate()
	fmt.Println(client)

	client.Activate(&client)
	fmt.Println(client)

}
