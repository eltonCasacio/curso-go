package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("\nCapacidade Inicial do slice: %v\n", cap(s))
	fmt.Printf("####################################################################\n\n")

	a := s[2:]
	b := s[:2]
	c := s[1:3]
	d := s[:0]
	e := s[0:]

	fmt.Printf("A) - %v, %v Ignora os dois primeiros valores: %v\n", len(a), cap(a), a)
	fmt.Printf("B) - %v, %v A partir da posicao zero, pega 2 elementos: %v\n", len(b), cap(b), b)
	fmt.Printf("C) - %v, %v Pega os intems da posicao 1 até posicao 2: %v\n", len(c), cap(c), c)
	fmt.Printf("D) - %v, %v Ignora todos os items e retorna array vazio: %v\n", len(d), cap(d), d)
	fmt.Printf("E) - %v, %v Pega todos os items: %v\n", len(e), cap(e), e)
	fmt.Printf("\n\n####################################################################\n\n")

	fmt.Printf("Append - Se o slice estiver na capacidade maxima e dermos um append, a capacidade do slice é dobrada\n\n")
	fmt.Printf("Tamanho(%v) - Capacidade(%v) inicial do slice\n", len(s), cap(s))
	s = append(s, 10)
	fmt.Printf("Tamanho(%v) - Capacidade(%v) atual do slice\n", len(s), cap(s))

	s = append(s, 11)
	s = append(s, 12)
	fmt.Printf("Tamanho(%v) - Capacidade(%v) atual do slice\n", len(s), cap(s))

}
