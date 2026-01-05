package tree

import (
	"container/list"
	"fmt"
)

type Diameter struct {
	val   int
	left  *Diameter
	right *Diameter
}

func (d *Diameter) Insert(val int) {
	if d.val > val {
		if d.left != nil {
			d.left.Insert(val)
		} else {
			d.left = &Diameter{val, nil, nil}
		}
	} else {
		if d.right != nil {
			d.right.Insert(val)
		} else {
			d.right = &Diameter{val, nil, nil}
		}
	}
}

func (d *Diameter) Height() int {
	queue := list.New()
	queue.PushBack(d)
	count := 0
	for queue.Len() > 0 {
		num := queue.Len()
		count += 1
		for num != 0 {
			elm := queue.Front()
			ptr := elm.Value.(*Diameter)
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

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (d *Diameter) Diameter() int {
	if d == nil {
		return 0
	}
	lh, rh := 0, 0
	if d.left != nil {
		lh = d.left.Height()
	}
	if d.right != nil {
		rh = d.right.Height()
	}
	max := findmax(findmax(lh, rh), lh+rh+1)
	dl, dr := 0, 0
	if d.left != nil {
		dl = d.left.Diameter()
	}
	if d.right != nil {
		dr = d.right.Diameter()
	}
	return findmax(findmax(max, dl), dr)
}

func (d *Diameter) finddiameter(maxm *int) int {
	if d == nil {
		return 0
	}
	ld, rd := 0, 0
	if d.left != nil {
		ld = d.left.finddiameter(maxm)
	}
	if d.right != nil {
		rd = d.right.finddiameter(maxm)
	}
	*maxm = findmax(*maxm, ld+rd)
	return 1 + findmax(ld, rd)
}

func (d *Diameter) TreeDiameter() int {
	maxm := 0
	d.finddiameter(&maxm)
	return maxm
}

func FindDiameter() {
	node := Diameter{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	node.Insert(5)
	node.Insert(4)
	node.Insert(14)
	node.Insert(9)
	fmt.Println("Diameter of the binary tree is:")
	fmt.Println(node.Diameter())
}

func NewTreeDiam() {
	node := Diameter{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	node.Insert(5)
	node.Insert(4)
	node.Insert(14)
	node.Insert(9)
	node.Insert(13)
	node.Insert(11)
	fmt.Println("Diameter of the b tree is:")
	fmt.Println(node.TreeDiameter())
}
