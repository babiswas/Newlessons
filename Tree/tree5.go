package tree

import (
	"container/list"
	"fmt"
)

type Path struct {
	val   int
	left  *Path
	right *Path
}

func (p *Path) Insert(val int) {
	if p.val > val {
		if p.left != nil {
			p.left.Insert(val)
		} else {
			p.left = &Path{val, nil, nil}
		}
	} else {
		if p.right != nil {
			p.right.Insert(val)
		} else {
			p.right = &Path{val, nil, nil}
		}
	}
}

func (p *Path) Height() int {
	queue := list.New()
	ptr := p
	queue.PushBack(ptr)
	count := 0
	for queue.Len() > 0 {
		num := queue.Len()
		count += 1
		for num != 0 {
			elm := queue.Front()
			ptr := elm.Value.(*Path)
			queue.Remove(elm)
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
		}
	}
	return count
}

func HeightTree() {
	node := Path{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The height of the binary tree:")
	height := node.Height()
	fmt.Println(height)
}
