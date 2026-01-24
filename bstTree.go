package BST

import (
	"container/list"
	"fmt"
)

type NewTree struct {
	val   int
	left  *NewTree
	right *NewTree
}

func (n *NewTree) Insert(val int) {
	if n.val > val {
		if n.left != nil {
			n.left.Insert(val)
		} else {
			n.left = &NewTree{val, nil, nil}
		}
	} else {
		if n.right != nil {
			n.right.Insert(val)
		} else {
			n.right = &NewTree{val, nil, nil}
		}
	}
}

func (n *NewTree) levelOrder() {
	queue := list.New()
	queue.PushBack(n)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*NewTree)
			fmt.Print(ptr.val)
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			queue.Remove(node)
			fmt.Print(" ")
			num -= 1
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func CreateTree() {
	node := NewTree{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The level order traversal of the queue is:")
	node.levelOrder()
}
