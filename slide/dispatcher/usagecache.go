package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

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

func (*Add) Cache() {}

type Cacheable interface {
	Cache()
}

type cacheMux struct {
	dispatcher.Mux
	cache map[string][]byte
}

func encode(m interface{}) []byte {
	buf := bytes.Buffer{}
	gob.NewEncoder(&buf).Encode(m)
	return buf.Bytes()
}

func decode(p []byte, m interface{}) {
	gob.NewDecoder(bytes.NewReader(p)).Decode(m)
}

// START OMIT

func (d *cacheMux) Dispatch(ctx context.Context, m dispatcher.Message) error {
	if d.cache == nil {
		d.cache = make(map[string][]byte)
	}
	_, cacheable := m.(Cacheable)
	var key string
	if cacheable {
		key = string(encode(m))
		p := d.cache[key]
		if p != nil {
			fmt.Println("cache hit", m)
			decode(p, m)
			return nil
		}
		fmt.Println("cache miss", m)
	}
	err := d.Mux.Dispatch(ctx, m)
	if err != nil {
		return err
	}
	if cacheable {
		d.cache[key] = encode(m)
	}
	return nil
}

// END OMIT

func main() {
	d := cacheMux{}
	d.Register(add)

	q := Add{A: 1, B: 2}
	d.Dispatch(context.TODO(), &q)
	fmt.Println("Result:", q.Result)

	q = Add{A: 2, B: 2}
	d.Dispatch(context.TODO(), &q)
	fmt.Println("Result:", q.Result)

	q = Add{A: 1, B: 2}
	d.Dispatch(context.TODO(), &q)
	fmt.Println("Result:", q.Result)
}
