package main

import (
	"fmt"
	"github.com/hjcian/ds/stack"
)

func main() {
	a := stack.NewItemStack()
	fmt.Println(123)
	a.Push(123)
	a.Push(456)
	a.Push(789)
	fmt.Println("init: ", a)
	b := a.Pull()
	fmt.Println("pull: ", b)
	c := a.Pull()
	fmt.Println("pull: ", c)
	d := a.Pull()
	fmt.Println("pull: ", d)
	e := a.Pull()
	fmt.Println("pull: ", e)
	fmt.Println("left: ", a)
}
