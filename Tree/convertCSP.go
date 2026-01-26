package BST

import (
	"container/list"
	"fmt"
)

type ChildrenS struct {
	val   int
	left  *ChildrenS
	right *ChildrenS
}

func (c *ChildrenS) Insert(val int) {
	if c.val > val {
		if c.left != nil {
			c.left.Insert(val)
		} else {
			c.left = &ChildrenS{val, nil, nil}
		}
	} else {
		if c.right != nil {
			c.right.Insert(val)
		} else {
			c.right = &ChildrenS{val, nil, nil}
		}
	}
}

func (c *ChildrenS) LevelOrder() {
	queue := list.New()
	queue.PushBack(c)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*ChildrenS)
			fmt.Print(ptr.val, " ")
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			queue.Remove(node)
			num -= 1
		}
		fmt.Println("")
	}
}

func (c *ChildrenS) ConvertCS() {
	if c == nil {
		return
	}
	var child int
	if c.left != nil {
		child += c.left.val
	}
	if c.right != nil {
		child += c.right.val
	}
	if child >= c.val {
		c.val = child
	} else {
		if c.left != nil {
			c.left.val = c.val
		}
		if c.right != nil {
			c.right.val = c.val
		}
	}
	if c.left != nil {
		c.left.ConvertCS()
	}
	if c.right != nil {
		c.right.ConvertCS()
	}
	var total int
	total = 0
	if c.left != nil {
		total += c.left.val
	}
	if c.right != nil {
		total += c.right.val
	}
	if c.left != nil || c.right != nil {
		c.val = total
	}
}

func ConvertTree() {
	node := ChildrenS{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("Convert tree to follow children sum property.")
	node.LevelOrder()
	fmt.Println("===============")
	node.ConvertCS()
	node.LevelOrder()
}
