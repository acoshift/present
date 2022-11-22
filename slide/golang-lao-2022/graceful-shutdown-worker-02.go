package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// START 3 OMIT
func main() {
	w := newWork()
	go w.Do()

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop
	fmt.Println("worker is shutting down...")
	w.Stop()
	fmt.Println("worker shutdown")
}

// END 3 OMIT

// START 1 OMIT
type work struct {
	done chan struct{}
	stop chan struct{}
}

func newWork() *work {
	return &work{
		done: make(chan struct{}),
	}
}

// END 1 OMIT

// START 2 OMIT
func (w *work) Do() {
	var i int
	for {
		// do something
		fmt.Println("work: doing", i)
		time.Sleep(2 * time.Second)
		fmt.Println("work: done", i)
		i++

		select {
		default:
		case <-w.done:
			fmt.Println("work: stop")
			close(w.stop)
			return
		}
	}
}

func (w *work) Stop() {
	w.stop = make(chan struct{})
	close(w.done)
	<-w.stop
}

// END 2 OMIT
