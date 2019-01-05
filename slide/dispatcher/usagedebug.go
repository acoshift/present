package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/moonrhythm/dispatcher"
)

func add(_ context.Context, m *Add) error {
	m.Result = m.A + m.B
	return nil
}

type Add struct {
	A, B   int
	Result int
}

// START OMIT

type debugMux struct {
	dispatcher.Mux
}

func (d *debugMux) Dispatch(ctx context.Context, m dispatcher.Message) error {
	err := d.Mux.Dispatch(ctx, m)
	spew.Dump(m)
	if err != nil {
		spew.Dump(err)
	}
	return err
}

// END OMIT

func main() {
	d := debugMux{}
	d.Register(add)

	q := Add{A: 1, B: 2}
	err := d.Dispatch(context.TODO(), &q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", q.Result)
}
