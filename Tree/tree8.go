package tree

import (
	"container/list"
	"fmt"
	"math"
)

type Balanced struct {
	val   int
	left  *Balanced
	right *Balanced
}

func (b *Balanced) Insert(val int) {
	if b.val > val {
		if b.left != nil {
			b.left.Insert(val)
		} else {
			b.left = &Balanced{val, nil, nil}
		}
	} else {
		if b.right != nil {
			b.right.Insert(val)
		} else {
			b.right = &Balanced{val, nil, nil}
		}
	}
}

func (b *Balanced) Height() int {
	queue := list.New()
	ptr := b
	queue.PushBack(ptr)
	count := 0
	for queue.Len() > 0 {
		num := queue.Len()
		count += 1
		for num != 0 {
			elm := queue.Front()
			ptr = elm.Value.(*Balanced)
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

func maxm(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *Balanced) isBalanced() bool {
	if b == nil {
		return true
	}
	h1, h2 := 0, 0
	if b.left != nil {
		h1 = b.left.Height()
	}

	if b.right != nil {
		h2 = b.right.Height()
	}

	fmt.Println(h1, h2)
	if math.Abs(float64(h1-h2)) > float64(1) {
		return false
	}
	l1, l2 := true, true
	if b.left != nil {
		l1 = b.left.isBalanced()
	}

	if b.right != nil {
		l2 = b.right.isBalanced()
	}

	if !l1 || !l2 {
		return false
	}

	return true
}

func (b *Balanced) balancedUtil() int {
	if b == nil {
		return 0
	}
	lh, rh := 0, 0
	if b.left != nil {
		lh = b.left.balancedUtil()
	}
	if lh == -1 {
		return -1
	}
	if b.right != nil {
		rh = b.right.balancedUtil()
	}
	if rh == -1 {
		return -1
	}
	if math.Abs(float64(lh-rh)) > float64(1) {
		return -1
	}
	return maxm(lh, rh) + 1
}

func (b *Balanced) ISBalanced() bool {
	status := b.balancedUtil()
	if status == -1 {
		return false
	}
	return true
}

func IsBalancedTree() {
	fmt.Println("Constructing the tree:")
	node := Balanced{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("Is the tree balanced.")
	status := node.isBalanced()
	fmt.Println(status)
}

func BalancedTree() {
	node := Balanced{15, nil, nil}
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	fmt.Println("Is the binary tree balanced.")
	status := node.ISBalanced()
	fmt.Println(status)
}
