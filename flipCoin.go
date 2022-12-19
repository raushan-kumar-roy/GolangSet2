package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	coin := []string{
		"heads",
		"tails",
	}

	rand.Seed(time.Now().UnixNano())

	side := coin[rand.Intn(len(coin))]

	fmt.Println("Flipped the coin and you get : ", side)
}
