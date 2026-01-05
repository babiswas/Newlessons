package tree

import (
	"container/list"
	"fmt"
)

type LevelOrder2 struct {
	val   int
	left  *LevelOrder2
	right *LevelOrder2
}

func (l *LevelOrder2) Insert(val int) {
	if l.val > val {
		if l.left != nil {
			l.left.Insert(val)
		} else {
			l.left = &LevelOrder2{val, nil, nil}
		}
	} else {
		if l.right != nil {
			l.right.Insert(val)
		} else {
			l.right = &LevelOrder2{val, nil, nil}
		}
	}
}

func (l *LevelOrder2) LevelOrder() {
	queue := list.New()
	queue.PushBack(l)
	for queue.Len() > 0 {
		nums := queue.Len()
		for nums != 0 {
			elm := queue.Front()
			node := elm.Value.(*LevelOrder2)
			fmt.Print(node.val)
			fmt.Print(" ")
			queue.Remove(elm)
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
			nums -= 1
		}
		fmt.Println("")
	}
}

func LevelOrderTrav() {
	node := LevelOrder2{15, nil, nil}
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	fmt.Println("Level order traversal of the tree:")
	node.LevelOrder()
}
