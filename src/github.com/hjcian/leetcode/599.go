package main

import (
	"fmt"
	"runtime"
)

func findRestaurant(list1 []string, list2 []string) []string {
	dict := make(map[string]int)
	for i, v := range list2 {
		dict[v] = i
	}
	minIdxSum := len(list1) + len(list2)
	results := make([]string, 0)
	for i, v := range list1 {
		j, ok := dict[v]
		if ok && i+j <= minIdxSum {
			results = append(results, v)
			minIdxSum = i + j
		}
	}
	return results
}

type suit struct {
	Andy  []string
	Doris []string
}

func main() {
	fmt.Println(runtime.Version())
	// a := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	// b := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	// result := findRestaurant(a, b)
	// fmt.Printf("%v \n", result)

	testSuites := []suit{
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		},
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"KFC", "Shogun", "Burger King"},
		},
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
		},
	}

	for _, suit := range testSuites {
		result := findRestaurant(suit.Andy, suit.Doris)
		fmt.Printf("%v \n", result)
	}
}
