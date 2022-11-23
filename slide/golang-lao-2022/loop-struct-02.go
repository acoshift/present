package main

import "fmt"

// START OMIT
type X struct {
	ID int
}

func main() {
	xs := []X{
		{1},
		{2},
		{3},
	}
	var rs []*X
	for _, x := range xs {
		x := x // <============
		rs = append(rs, &x)
	}

	for _, r := range rs {
		fmt.Println(r.ID)
	}
}

// END OMIT
