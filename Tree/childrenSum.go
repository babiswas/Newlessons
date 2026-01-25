package BST

import "fmt"

type ChildrenSum struct {
	val   int
	left  *ChildrenSum
	right *ChildrenSum
}

func (c *ChildrenSum) isChildrenSum() bool {
	if (c == nil) || ((c.left == nil) && (c.right == nil)) {
		return true
	}
	lc, rc := 0, 0
	if c.left != nil {
		lc = c.left.val
	}
	if c.right != nil {
		rc = c.right.val
	}
	if lc+rc != c.val {
		return false
	}
	status1, status2 := true, true
	if c.left != nil {
		status1 = c.left.isChildrenSum()
	}
	if c.right != nil {
		status2 = c.right.isChildrenSum()
	}
	return status1 && status2
}

func VerifyChildrenSum() {
	node := ChildrenSum{10, nil, nil}
	node.left = &ChildrenSum{8, nil, nil}
	node.left.left = &ChildrenSum{3, nil, nil}
	node.left.right = &ChildrenSum{5, nil, nil}
	node.right = &ChildrenSum{2, nil, nil}
	node.right.right = &ChildrenSum{2, nil, nil}
	status := node.isChildrenSum()
	fmt.Println("Children sum property status:", status)
}
