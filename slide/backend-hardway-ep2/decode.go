package main

import (
	"fmt"
)

const char = "0123456789abcdef"

// ascii => byte
func getByte(s byte) byte {
	for i, c := range char {
		if s == byte(c) {
			return byte(i)
		}
	}
	return 0
}

func decode(src string) []byte {
	dst := make([]byte, len(src)/2)

	for i := range dst {
		s0 := getByte(src[i*2])
		s1 := getByte(src[i*2+1])
		dst[i] = s0<<4 + s1
	}

	return dst
}

func main() {
	encoded := "68656c6c6f2c20656e636f64696e67203a44"

	fmt.Printf("encode:\n%s\n\n", encoded)
	fmt.Printf("decode:\n%s\n", decode(encoded))
}
