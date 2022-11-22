package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// START OMIT
	srv := http.Server{
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

	fmt.Println("server: start")
	log.Fatal(srv.ListenAndServe())
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
