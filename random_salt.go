package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		println(generateSalt(10))
	}
}

func generateSalt(size uint) (salt string) {
	rand.Seed(time.Now().UnixNano())
	data := md5.Sum([]byte(strconv.Itoa(rand.Int())))
	str := fmt.Sprintf("%x", data)
	salt = str[:size]
	return
}
