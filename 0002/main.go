package main

import (
	"fmt"
)

func main() {
	f1, f2, sum := 1, 1, 0

	for f2 <= 4_000_000 {
		f1, f2 = f2, f1+f2

		if f2%2 == 0 {
			sum += f2
		}
	}

	fmt.Println(sum)
}
