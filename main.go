package main

import (
	"fmt"
)

func main() {
	a := 1243
	fmt.Println(&a)
	b := &a
	fmt.Println(*b)
}
