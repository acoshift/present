package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data1 := "I do love you, don't I"
	data2 := "I don't love you, do I"
	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)

	sha11 := sha1.Sum([]byte(data1))
	fmt.Printf("sha1 data1: %x\n", sha11)

	sha12 := sha1.Sum([]byte(data2))
	fmt.Printf("sha1 data2: %x\n", sha12)
}
