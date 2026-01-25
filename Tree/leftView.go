package BST

import (
	"container/list"
	"fmt"
)

type LeftView struct {
	val   int
	left  *LeftView
	right *LeftView
}

func (l *LeftView) Insert(val int) {
	if l.val > val {
		if l.left != nil {
			l.left.Insert(val)
		} else {
			l.left = &LeftView{val, nil, nil}
		}
	} else {
		if l.right != nil {
			l.right.Insert(val)
		} else {
			l.right = &LeftView{val, nil, nil}
		}
	}
}

func (l *LeftView) Leftview() {
	queue := list.New()
	queue.PushBack(l)
	for queue.Len() != 0 {
		num := queue.Len()
		fmt.Println(queue.Back().Value.(*LeftView).val)
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*LeftView)
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			queue.Remove(node)
			num -= 1
		}
	}
}

func LeftViewBST() {
	node := LeftView{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Left view of the binary tree:")
	node.Leftview()
}
