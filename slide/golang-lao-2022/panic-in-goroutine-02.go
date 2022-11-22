package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	fmt.Println("main start")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic:", err)
			}
		}()
		panic("oh no")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done")
	// END OMIT
}
