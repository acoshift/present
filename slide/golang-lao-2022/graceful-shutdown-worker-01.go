package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go doWork()
	select {}
}

func doWork() {
	var i int
	for {
		// do something
		fmt.Println("doing work", i)
		time.Sleep(2 * time.Second)
		fmt.Println("done work", i)
		i++
	}
}

// END OMIT
