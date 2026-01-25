package BST

import (
	"container/list"
	"fmt"
)

type Height struct {
	val   int
	left  *Height
	right *Height
}

func (h *Height) Insert(val int) {
	if h.val > val {
		if h.left != nil {
			h.left.Insert(val)
		} else {
			h.left = &Height{val, nil, nil}
		}
	} else {
		if h.right != nil {
			h.right.Insert(val)
		} else {
			h.right = &Height{val, nil, nil}
		}
	}
}

func (h *Height) HeightBST() int {
	queue := list.New()
	queue.PushBack(h)
	height := 0
	for queue.Len() != 0 {
		num := queue.Len()
		height += 1
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*Height)
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
	return height
}

func HeightOfBST() {
	node := Height{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("Height of the binary serach tree:")
	height := node.HeightBST()
	fmt.Println(height)
}
