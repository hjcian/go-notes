package main

import (
	"fmt"
	"github.com/hjcian/ds/queue"
)

func main() {
	a := queue.NewSomethingQueue()
	fmt.Println(123)
	a.Push(123)
	a.Push(456)
	a.Push(789)
	fmt.Println(a)
}
