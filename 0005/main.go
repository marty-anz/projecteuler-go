package main

import (
	"fmt"
)

func divisibleUpto20(p int) bool {
	for factor := 2; factor <= 20; factor++ {
		if p%factor != 0 {
			return false
		}
	}

	return true
}

func main() {
	var p int

	for p = 20; !divisibleUpto20(p); p++ {
	}

	fmt.Println(p)
}
