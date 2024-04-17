package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			return false
		}

		start += 1
		end -= 1
	}

	return true
}

func main() {
	max := -1

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			p := i * j
			if isPalindrome(fmt.Sprintf("%d", p)) {
				if p > max {
					max = p
				}
			}
		}
	}

	fmt.Println(max)
}
