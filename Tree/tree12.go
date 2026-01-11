package tree

import "fmt"

type Sum struct {
	val   int
	left  *Sum
	right *Sum
}

func (s *Sum) Insert(val int) {
	if s.val > val {
		if s.left != nil {
			s.left.Insert(val)
		} else {
			s.left = &Sum{val, nil, nil}
		}
	} else {
		if s.right != nil {
			s.right.Insert(val)
		} else {
			s.right = &Sum{val, nil, nil}
		}
	}
}

func (s *Sum) SumToLeaf(num int) int {
	if s == nil {
		return 0
	}
	num = num*10 + s.val
	if s.left == nil && s.right == nil {
		return num
	}
	if s.left != nil && s.right == nil {
		return s.left.SumToLeaf(num)
	}
	if s.left == nil && s.right != nil {
		return s.right.SumToLeaf(num)
	}
	return s.left.SumToLeaf(num) + s.right.SumToLeaf(num)
}

func SumRootLeaf() {
	node := Sum{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The sum from root.leaf is:")
	num := 0
	fmt.Println(node.SumToLeaf(num))
}
