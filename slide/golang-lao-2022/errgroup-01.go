package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// START OMIT
	r1, err := work1()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	r2, err := work2()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// END OMIT

	fmt.Printf("r1: %d\nr2: %d\nduration: %s\n", r1, r2, time.Since(start))
}

func work1() (int, error) {
	fmt.Println("work1: start")
	time.Sleep(2 * time.Second)
	return 1, nil
}

func work2() (int, error) {
	fmt.Println("work2: start")
	time.Sleep(2 * time.Second)
	return 2, nil
}
