package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/moonrhythm/dispatcher"
)

type Add struct {
	A, B   int
	Result int
}

func add(_ context.Context, m *Add) error {
	m.Result = m.A + m.B
	return nil
}

func main() {
	d := dispatcher.NewMux()
	d.Register(add)

	// START OMIT
	wrapper := dispatcher.HTTPHandlerWrapper{
		Dispatcher: d,
		Decoder: func(r *http.Request, v interface{}) error {
			return json.NewDecoder(r.Body).Decode(v)
		},
		Encoder: func(w http.ResponseWriter, r *http.Request, v interface{}) error {
			w.Header().Set("Content-Type", "application/json")
			return json.NewEncoder(w).Encode(v)
		},
		Result: "Result",
	}

	time.AfterFunc(1*time.Second, func() {
		resp, _ := http.Post("http://localhost:9091/add", "application/json", strings.NewReader(`
			{"a": 3, "b": 4}
		`))
		io.Copy(os.Stdout, resp.Body)
		os.Exit(0)
	})

	http.Handle("/add", wrapper.Handler(new(Add)))
	http.ListenAndServe(":9091", nil)
	// END OMIT
}
