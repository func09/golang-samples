package main

import (
	"math/rand"
	"time"
)

var randRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// ランダム文字列生成
func main() {
	for i := 0; i < 10; i++ {
		println(generateRandomStrings(40))
	}
}

func generateRandomStrings(size int) string {
	var str string

	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(randRunes))
		str += string(randRunes[n])
	}

	return str
}
