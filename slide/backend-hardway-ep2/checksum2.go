package main

import (
	"fmt"
)

func djb2(b []byte) uint64 {
	var hash uint64 = 5381
	for _, c := range b {
		hash = hash<<5 + hash + uint64(c)
	}
	return hash
}

func main() {
	data1 := "I do love you, don't I"
	data2 := "I don't love you, do I"
	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)

	djb21 := djb2([]byte(data1))
	fmt.Printf("djb2 data1: %x\n", djb21)

	djb22 := djb2([]byte(data2))
	fmt.Printf("djb2 data2: %x\n", djb22)
}
