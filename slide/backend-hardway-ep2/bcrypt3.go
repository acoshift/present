package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	start := time.Now()
	runSHA256()
	fmt.Println("SHA256", time.Since(start))

	start = time.Now()
	runBcrypt()
	fmt.Println("Bcrypt", time.Since(start))
}

func runSHA256() {
	sha256.Sum256([]byte("admin"))
}

func runBcrypt() {
	bcrypt.GenerateFromPassword([]byte("admin"), 10)
}
