package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	counter := make(map[rune]int)
	for i, v := range s {
		fmt.Println(i, v)
		counter[v]++
	}
	fmt.Println(counter)
	sum := 0
	flag := true
	for _, v := range counter {
		if v%2 == 0 {
			sum += v
		} else {
			sum += v - 1
			if flag {
				flag = false
				sum++
			}
		}
	}
	return sum
}

func main() {

	inputs := []string{
		"abccccdd",
	}
	for _, v := range inputs {
		ret := longestPalindrome(v)
		fmt.Printf("%v, %v \n", v, ret)
	}
}
