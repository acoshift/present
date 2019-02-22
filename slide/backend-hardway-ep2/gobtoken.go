package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

// START OMIT
type Token struct {
	UserID string
	Nonce  [6]byte
}

func main() {
	t := Token{UserID: "123"}
	token, _ := generateToken(t, []byte("supersecret"))
	fmt.Println(token)
}

// END OMIT

func generateToken(t Token, key []byte) (string, error) {
	rand.Read(t.Nonce[:])

	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(t)
	if err != nil {
		return "", err
	}

	sig := signHMAC(buf.Bytes(), key)
	token := base64.RawURLEncoding.EncodeToString(buf.Bytes())
	token += "."
	token += base64.RawURLEncoding.EncodeToString(sig)
	return token, nil
}

func signHMAC(data []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	return h.Sum(data)
}
