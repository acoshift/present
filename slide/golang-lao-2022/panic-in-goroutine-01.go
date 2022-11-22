package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	fmt.Println("main start")
	go func() {
		panic("oh no")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done")
	// END OMIT
}
