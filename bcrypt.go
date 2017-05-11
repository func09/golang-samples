package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password"
	salt := "salt"
	hashes := []string{}

	for i := 0; i < 10; i++ {
		hashedPassword := generateEncryptedPassword(password, salt)
		hashes = append(hashes, hashedPassword)
	}

	for _, hashedPassword := range hashes {
		fmt.Println(hashedPassword)
		fmt.Println(compareEncryptedPassword(hashedPassword, password, salt))
	}

}

func generateEncryptedPassword(password, salt string) (encryptedPassword string) {
	converted, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	encryptedPassword = string(converted)
	return
}

func compareEncryptedPassword(hashedPassword, password, salt string) (ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))

	if err == nil {
		ok = true
	}

	return
}
