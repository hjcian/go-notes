// reference: https://flaviocopes.com/golang-data-structure-linked-list/
package linkedlist

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	content Item
	next    *Node
}

type LinkedList struct {
	head *Node
	size int
	// lock sync.RWMutex
}

func (ll *LinkedList) Append(t Item) {
	node := &Node{t, nil}
	if ll.head == nil {
		ll.head = node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = node
	}
	ll.size++
}

func (ll *LinkedList) Insert(pos int, t Item) error {
	node := Node{t, nil}
	if pos < 0 || pos > ll.Size() {
		return fmt.Errorf("Index out of bounds")
	}
	if pos == 0 {
		node.next = ll.head
		ll.head = &node
		ll.size++
		return nil
	}
	curr := ll.head
	counter := 1
	for curr != nil {
		if pos == counter {
			node.next = curr.next
			curr.next = &node
			break
		}
		counter++
		curr = curr.next
	}
	ll.size++
	return nil
}

func (ll *LinkedList) RemoveAt(pos int) (*Item, error) {
	if pos < 0 || pos >= ll.Size() {
		return nil, fmt.Errorf("Index out of bounds")
	}
	if pos == 0 {
		ret := ll.head
		ll.head = ret.next
		ll.size--
		return &ret.content, nil
	}
	curr := ll.head
	counter := 0
	for curr != nil {
		if counter+1 == pos {
			ret := curr.next
			curr.next = curr.next.next
			ll.size--
			return &ret.content, nil
		}
		counter++
		curr = curr.next
	}
	return nil, nil
}

func (ll *LinkedList) IndexOf(t Item) int {
	curr := ll.head
	counter := 0
	for curr != nil {
		if curr.content == t {
			return counter
		}
		counter++
		curr = curr.next
	}
	return -1
}

func (ll *LinkedList) IsEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) String() {
	// ll.lock.RLock()
	// defer ll.lock.RUnlock()
	node := ll.head
	j := 0
	fmt.Print(ll.Size(), " [ ")
	for {
		if node == nil {
			break
		}
		j++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Print("] ")
	fmt.Println()
}
