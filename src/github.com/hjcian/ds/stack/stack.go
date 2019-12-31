package stack

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

// ItemStack - is a Last-In-First-Out (LIFO) data structure of any item
type ItemStack struct {
	items []Item
}

// NewItemStack - a structure factory
func NewItemStack() *ItemStack {
	return &ItemStack{items: make([]Item, 0)}
}

// Push - add item to top of pile
func (s *ItemStack) Push(i Item) {
	s.items = append(s.items, i)
}

// Pull - pull item from top of pile
func (s *ItemStack) Pull() Item {
	if len(s.items) == 0 {
		return nil
	}
	ret := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return ret
}
