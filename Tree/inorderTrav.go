package BST

import (
	"container/list"
	"fmt"
)

type Inorder struct {
	val   int
	left  *Inorder
	right *Inorder
}

func (i *Inorder) Insert(val int) {
	if i.val > val {
		if i.left != nil {
			i.left.Insert(val)
		} else {
			i.left = &Inorder{val, nil, nil}
		}
	} else {
		if i.right != nil {
			i.right.Insert(val)
		} else {
			i.right = &Inorder{val, nil, nil}
		}
	}
}

func (i *Inorder) InorderTrav() {
	stack := list.New()
	ptr := i
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else if stack.Len() != 0 {
			node := stack.Back()
			ptr = node.Value.(*Inorder)
			stack.Remove(node)
			fmt.Println(ptr.val)
			ptr = ptr.right
		} else {
			break
		}
	}
}

func InorderTraversal() {
	node := Inorder{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The inorder traversal of the tree is:")
	node.InorderTrav()
}
