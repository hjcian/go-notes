package main

import (
	"fmt"
)

func main() {
	// naive cteation
	// grade1 := 97
	// grade2 := 90
	// grade3 := 84
	// fmt.Printf("Grade: %v, %v, %v", grade1, grade2, grade3)

	// define an array
	// grades := [3]int{97, 85, 93}
	// grades := [...]int{97, 85, 93}     // no need to assign length
	// fmt.Printf("Grade: %v \n", grades) // Grade: [97 85 93]

	// define an array of string
	// var students [3]string
	// fmt.Printf("Student: %v \n", students) // Student: [  ]
	// students[0] = "melody"
	// students[1] = "max"
	// students[2] = "chuchutaro"
	// fmt.Printf("Student: %v \n", students)            // Student: [melody max chuchutaro]
	// fmt.Printf("Student #1: %v \n", students[1])      // Student #1: max
	// fmt.Printf("# of Students: %v \n", len(students)) // # of Students: 3

	// define a matrix
	// var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// var matrix [3][3]int
	// matrix[0] = [3]int{1, 0, 0}
	// matrix[1] = [3]int{0, 1, 0}
	// matrix[2] = [3]int{0, 0, 1}
	// fmt.Printf("%v\n", matrix) // [[1 0 0] [0 1 0] [0 0 1]]\

	// copy an array
	// ori := [3]int{1, 0, 0} // Or: var ori [3]int = [3]int{1, 0, 0}
	// copied := ori
	// copied[1] = 5
	// fmt.Printf("ori: %v \n", ori)       // ori: [1 0 0]
	// fmt.Printf("copied: %v \n", copied) // copied: [1 5 0]

	// copy by pointer, take reference of ori
	// ori2 := [3]int{1, 0, 0} // Or: var ori [3]int = [3]int{1, 0, 0}
	// copied2 := &ori2
	// copied2[1] = 5
	// fmt.Printf("ori2: %v \n", ori2)       // ori2: [1 5 0]
	// fmt.Printf("copied2: %v \n", copied2) // copied2: &[1 5 0]

	// Slices are like references to arrays
	// a := [...]int{1, 2, 3, 4, 5}
	// fmt.Printf("len: %v \n", len(a))
	// fmt.Printf("capacity: %v \n", cap(a))
	// b := a[:3]
	// c := a[2:5]
	// // d := a[:10] // panic: runtime error: slice bounds out of range [:10] with capacity 5
	// c[0] = 52
	// fmt.Printf("%v %v %v\n", a, b, c) // [1 2 52 4 5] [1 2 52] [52 4 5]
	// a[2] = 62
	// fmt.Printf("%v %v %v\n", a, b, c) // [1 2 62 4 5] [1 2 62] [62 4 5]

	// how append vary the capacity
	// a := make([]int, 0, 0)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [] len: 0, cap: 0
	// a = append(a, 1)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1] len: 1, cap: 1
	// a = append(a, 1)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1 1] len: 2, cap: 2
	// a = append(a, 1)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1 1 1] len: 3, cap: 4
	// a = append(a, 2, 2)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1 1 1 2 2] len: 5, cap: 8

	// if we initialized array by assign a expected capacity, will more efficient to program
	// a := make([]int, 0, 8)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [] len: 0, cap: 8
	// a = append(a, 1)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1] len: 1, cap: 8
	// a = append(a, 1)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1 1] len: 2, cap: 8
	// a = append(a, 1)
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1 1 1] len: 3, cap: 8
	// a = append(a, []int{2, 2}...)                           // create by more generic assignment
	// fmt.Printf("%v len: %v, cap: %v \n", a, len(a), cap(a)) // [1 1 1 2 2] len: 5, cap: 8

	a := []int{}
	// a := []int{1, 2, 3, 4, 5}
	// c := []int{9, 8, 7}
	fmt.Printf("%v, %v, %v, %T \n", a, len(a), cap(a), a)
	// fmt.Printf("%v %v\n", a[:2], a[3:])
	// b := append(a[:2], c[:]...)
	// fmt.Printf("%v \n", b)
	// fmt.Printf("%v \n", a)

}
