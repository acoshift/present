package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	start := time.Now()

	// START OMIT
	var (
		r1 int
		r2 int
	)

	var eg errgroup.Group
	eg.Go(func() error {
		var err error
		r1, err = work1()
		return err
	})
	eg.Go(func() error {
		var err error
		r2, err = work2()
		return err
	})
	err := eg.Wait()
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
