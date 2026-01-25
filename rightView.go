package BST

import (
	"container/list"
	"fmt"
)

type RightView struct {
	val   int
	left  *RightView
	right *RightView
}

func (r *RightView) Insert(val int) {
	if r.val > val {
		if r.left != nil {
			r.left.Insert(val)
		} else {
			r.left = &RightView{val, nil, nil}
		}
	} else {
		if r.right != nil {
			r.right.Insert(val)
		} else {
			r.right = &RightView{val, nil, nil}
		}
	}
}

func (r *RightView) RightView() {
	queue := list.New()
	queue.PushBack(r)
	for queue.Len() != 0 {
		num := queue.Len()
		fmt.Println(queue.Front().Value.(*RightView).val)
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*RightView)
			queue.Remove(node)
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
		}
	}
}

func RightViewTree() {
	node := RightView{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The right view of the binary tree is:")
	node.RightView()
}
