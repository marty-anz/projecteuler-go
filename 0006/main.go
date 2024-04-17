package main

import (
	"fmt"
)

func sumOfSquaresUpto(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func triangleNumber(n int) int {
	return n * (n + 1) / 2
}

func main() {

	t := triangleNumber(100)

	fmt.Println(t*t - sumOfSquaresUpto(100))
}
