package main

import (
	"fmt"
)

func scan() (int, int, int) {
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}

	return 0, 0, 0
}

func main() {
	a, b, c := scan()

	fmt.Println(a * b * c)
}
