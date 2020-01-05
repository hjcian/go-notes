package main

import (
	"container/list"
	"fmt"
)

type Item interface{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (bst *Tree) Insert(key int) {
	n := &TreeNode{key, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(node, newNode *TreeNode) {
	if newNode.Val < node.Val {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func (bst *Tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *TreeNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Val)
		stringify(n.Right, level)
	}
}

func buildTree(testCase *[]Item) *Tree {
	var bst Tree
	for _, item := range *testCase {
		key, ok := item.(int)
		if ok {
			bst.Insert(key)
		}
	}
	bst.String()
	return &bst
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	included := make([]*TreeNode, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		qnode := queue.Front()
		node := qnode.Value.(*TreeNode) // type assertion
		if node != nil {
			if node.Val >= L && node.Val <= R {
				included = append(included, node)
			}
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		}
		queue.Remove(qnode)
	}
	sum := 0
	for _, node := range included {
		sum += node.Val
	}
	return sum
}

func main() {
	case1 := []Item{10, 5, 15, 3, 7, nil, 18}
	bst := buildTree(&case1)
	caseResult := rangeSumBST(bst.root, 7, 15)
	if caseResult != 32 {
		fmt.Printf("expected 32, return %v", caseResult)
	}

	case2 := []Item{10, 5, 15, 3, 7, 13, 18, 1, nil, 6}
	bst = buildTree(&case2)
	caseResult = rangeSumBST(bst.root, 6, 10)
	if caseResult != 23 {
		fmt.Printf("expected 23, return %v", caseResult)
	}
}
