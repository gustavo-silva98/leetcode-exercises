package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	var y string
	for i := len(s) - 1; i >= 0; i-- {
		y += string(s[i])
	}
	return s == y
}

func main() {
	num := 123
	isPalindrome(num)
}
