package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "superman"
	fmt.Println("password:", password)
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	fmt.Println("hashed:", string(hashed))
}
