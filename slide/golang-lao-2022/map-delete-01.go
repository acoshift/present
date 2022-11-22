package main

import (
	"fmt"
	"runtime"
)

func main() {
	// START OMIT
	m := map[int]int{}
	addData(m)

	for k := range m {
		delete(m, k)
	}

	printMemStats()
	runtime.GC()
	printMemStats()
	_ = m[0]
	// END OMIT
}

func map1() {
	m := map[int]int{}
	addData(m)
	printMemStats()
}

func addData(m map[int]int) {
	for i := 0; i < 500000; i++ {
		m[i] = i
	}
}

func printMemStats() {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	fmt.Printf("Alloc = %v MiB\n", s.Alloc/1024/1024)
}
