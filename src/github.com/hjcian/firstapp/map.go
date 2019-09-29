package main

import (
	"fmt"
)

func main() {
	dict := map[string]int{
		"a": 123,
		"b": 456,
	}
	// fmt.Printf("%v \n ", dict)
	// m := map[[]int]int{} // invalid map key type []int

	// slice can not be key, array can
	// m := map[[3]int]int{}
	// fmt.Printf("%v \n ", m)

	// check key in map
	// v, ok := dict["c"]
	// fmt.Printf("%v %v \n ", v, ok) // 0 false
	// v, ok := dict["b"]
	// fmt.Printf("%v %v \n ", v, ok) // 456 true

	copied := dict
	dict["a"] = 1111            // OR copied["a"] = 1111
	fmt.Printf("%v \n", dict)   // map[a:1111 b:456]
	fmt.Printf("%v \n", copied) // map[a:1111 b:456]

	// create map by make
	a := make(map[string]int)
	a = map[string]int{
		"c": 123,
		"d": 456,
	}
	fmt.Printf("%v \n", a) // map[a:1111 b:456]
}
