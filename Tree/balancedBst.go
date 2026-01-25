package BST

import (
	"fmt"
	"math"
)

type Balanced struct {
	val   int
	left  *Balanced
	right *Balanced
}

func maxm(lh, rh int) int {
	if lh > rh {
		return lh
	}
	return rh
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

func (b *Balanced) BalancedUtil() int {
	if b == nil {
		return 0
	}
	lh, rh := 0, 0
	if b.left != nil {
		lh = b.left.BalancedUtil()
	}
	if lh == -1 {
		return -1
	}
	if b.right != nil {
		rh = b.right.BalancedUtil()
	}
	if rh == -1 {
		return -1
	}
	if math.Abs(float64(lh-rh)) > 1 {
		return -1
	}
	return maxm(lh, rh) + 1
}

func (b *Balanced) isBalanced() bool {
	if b.BalancedUtil() == -1 {
		return false
	}
	return true
}

func FindIfBalanced() {
	node := Balanced{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("Verifying if the tree is balanced.")
	if node.isBalanced() {
		fmt.Println("The tree is balanced.")
	} else {
		fmt.Println("The tree is not balanced.")
	}
}

func FindIfBalancedTree() {
	node := Balanced{15, nil, nil}
	node.Insert(9)
	node.Insert(8)
	node.Insert(7)
	node.Insert(23)
	node.Insert(18)
	node.Insert(25)
	node.Insert(24)
	fmt.Println("Verifying if the tree is height balanced:")
	status := node.isBalanced()
	if status {
		fmt.Println("The tree is balanced b tree:")
	} else {
		fmt.Println("The tree isn't balanced btree.")
	}
}
