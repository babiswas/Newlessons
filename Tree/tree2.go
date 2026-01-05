package tree

import (
	"container/list"
	"fmt"
)

type LevelOrder struct {
	val   int
	left  *LevelOrder
	right *LevelOrder
}

func (l *LevelOrder) Insert(val int) {
	if l.val > val {
		if l.left != nil {
			l.left.Insert(val)
		} else {
			l.left = &LevelOrder{val, nil, nil}
		}
	} else {
		if l.right != nil {
			l.right.Insert(val)
		} else {
			l.right = &LevelOrder{val, nil, nil}
		}
	}
}

func (l *LevelOrder) LevelOrderTrav() {
	fmt.Println("Level order traversal:")
	queue := list.New()
	queue.PushBack(l)
	for queue.Len() > 0 {
		elm := queue.Front()
		queue.Remove(elm)
		node := elm.Value.(*LevelOrder)
		fmt.Println(node.val)
		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
}

func TreeLevelOrder() {
	node := LevelOrder{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Level order traversal of the tree:")
	node.LevelOrderTrav()
}
