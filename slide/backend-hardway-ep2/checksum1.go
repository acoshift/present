package main

import (
	"fmt"
	"math/big"
)

func sum(b []byte) []byte {
	s := new(big.Int)
	for _, p := range b {
		s.Add(s, new(big.Int).SetInt64(int64(p)))
	}
	return s.Bytes()
}

func main() {
	data1 := "I do love you, don't I"
	data2 := "I don't love you, do I"
	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)

	sum1 := sum([]byte(data1))
	fmt.Printf("sum data1: %x\n", sum1)

	sum2 := sum([]byte(data2))
	fmt.Printf("sum data2: %x\n", sum2)

	fmt.Println("hash collision!!!")
}
