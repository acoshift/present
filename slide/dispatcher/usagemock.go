package main

import (
	"context"
	"fmt"
	"log"

	"github.com/moonrhythm/dispatcher"
)

func add(_ context.Context, m *Add) error {
	m.Result = m.A + m.B
	return nil
}

// START OMIT
type Add struct {
	A, B   int
	Result int
}

func main() {
	d := dispatcher.NewMux()
	d.Register(func(_ context.Context, m *Add) error {
		return fmt.Errorf("can not add")
	})

	q := Add{A: 1, B: 2}
	err := d.Dispatch(context.TODO(), &q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", q.Result)
}

// END OMIT
