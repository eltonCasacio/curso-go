package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // definindo que o contexto terá 3 segundos para executar...caso contrario cancela a execucao
	defer cancel()
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // se atingiar o tempo limite do contexto
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(2 * time.Second): // se executar dentro do tempo limite do contexto
		fmt.Println("Hotel booked")
	}
}
