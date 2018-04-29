package main

import (
	"fmt"

	binance "github.com/bloc4ain/go-binance"
)

func main() {
	// Ping
	if err := binance.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Ping: OK")
	}

	// GetServerTime
	if t, err := binance.GetServerTime(); err != nil {
		panic(err)
	} else {
		fmt.Println("Server time:", t)
	}
}
