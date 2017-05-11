package main

import (
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		println(randomInt(10))
	}
}

func randomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
