package main

import (
	"fmt"
)

const char = "0123456789abcdef"

func encode(src []byte) string {
	dst := make([]byte, len(src)*2)

	for i, p := range src {
		dst[i*2] = char[p>>4]
		dst[i*2+1] = char[p&0xf]
	}

	return string(dst)
}

func main() {
	data := "hello, encoding :D"

	encoded := encode([]byte(data))

	fmt.Printf("text:\n%s\n\n", data)
	fmt.Printf("encode:\n%s\n\n", encoded)
}
