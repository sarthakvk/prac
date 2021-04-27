package main

import (
	"fmt"
)

type T struct {
	X string
}

func (t *T)String() string {
	return t.X
}

func main() {
	var a fmt.Stringer
	a = &T{"sarthak"}
	fmt.Println(a)
}