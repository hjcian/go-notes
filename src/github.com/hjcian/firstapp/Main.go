package main

import (
	"fmt"
	// "strconv"
	// "math"
)

const (
	isA = 1 << iota
	isB
	isC
	isD
)

func main() {
	var roles byte = isA | isC
	fmt.Printf("%b \n", roles)
	const a = 8
	fmt.Printf("%v, %T \n", a, a)
}
