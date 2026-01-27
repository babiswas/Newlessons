package BST

import (
	"container/list"
	"fmt"
)

type Invert struct {
	val   int
	left  *Invert
	right *Invert
}

func (i *Invert) Insert(val int) {
	if i.val > val {
		if i.left != nil {
			i.left.Insert(val)
		} else {
			i.left = &Invert{val, nil, nil}
		}
	} else {
		if i.right != nil {
			i.right.Insert(val)
		} else {
			i.right = &Invert{val, nil, nil}
		}
	}
}

func (i *Invert) InvertTree() {
	if i == nil {
		return
	}
	i.left, i.right = i.right, i.left
	if i.left != nil {
		i.left.InvertTree()
	}
	if i.right != nil {
		i.right.InvertTree()
	}
}

func (i *Invert) LevelOrderTrav() {
	queue := list.New()
	queue.PushBack(i)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*Invert)
			queue.Remove(node)
			fmt.Print(ptr.val, " ")
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
		}
		fmt.Println("")
	}
}

func InvertTreeAlgo() {
	node := Invert{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("Invert the binary tree:")
	node.LevelOrderTrav()
	fmt.Println("=====================")
	node.InvertTree()
	node.LevelOrderTrav()
}
