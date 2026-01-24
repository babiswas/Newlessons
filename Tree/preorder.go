package BST

import (
	"container/list"
	"fmt"
)

type Preorder struct {
	val   int
	left  *Preorder
	right *Preorder
}

func (p *Preorder) Insert(val int) {
	if p.val > val {
		if p.left != nil {
			p.left.Insert(val)
		} else {
			p.left = &Preorder{val, nil, nil}
		}
	} else {
		if p.right != nil {
			p.right.Insert(val)
		} else {
			p.right = &Preorder{val, nil, nil}
		}
	}
}

func (p *Preorder) PreorderTrav() {
	stack := list.New()
	stack.PushBack(p)
	for stack.Len() != 0 {
		node := stack.Back()
		ptr := node.Value.(*Preorder)
		fmt.Println(ptr.val)
		stack.Remove(node)
		if ptr.right != nil {
			stack.PushBack(ptr.right)
		}
		if ptr.left != nil {
			stack.PushBack(ptr.left)
		}
	}
}

func PreorderTraversal() {
	node := Preorder{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("The preorder traversal of the tree is:")
	node.PreorderTrav()
}
