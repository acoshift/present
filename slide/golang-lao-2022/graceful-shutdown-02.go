package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr: "[::1]:3333",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("server: responding in 5 seconds")
			time.Sleep(5 * time.Second)
			w.Write([]byte("ok"))
		}),
	}

	go func() {
		time.Sleep(200 * time.Millisecond)
		go sendRequest()
	}()

	// START OMIT
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop
	fmt.Println("server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("server shutdown error:", err)
		return
	}
	fmt.Println("server shutdown")
	// END OMIT
}

func sendRequest() {
	fmt.Println("request: send")
	resp, err := http.Get("http://[::1]:3333")
	if err != nil {
		fmt.Println("request:", err)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request:", err)
		return
	}
	fmt.Println("request:", string(b))
}
