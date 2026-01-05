package tree

import "fmt"

type Lca struct {
	val   int
	left  *Lca
	right *Lca
}

func (l *Lca) Insert(val int) {
	if l.val > val {
		if l.left != nil {
			l.left.Insert(val)
		} else {
			l.left = &Lca{val, nil, nil}
		}
	} else {
		if l.right != nil {
			l.right.Insert(val)
		} else {
			l.right = &Lca{val, nil, nil}
		}
	}
}

func (l *Lca) FindNode(num int) *Lca {
	if l.val == num {
		return l
	}
	var n1 *Lca
	var n2 *Lca
	n1, n2 = nil, nil
	if l.left != nil {
		n1 = l.left.FindNode(num)
	}
	if l.right != nil {
		n2 = l.right.FindNode(num)
	}

	if n1 == nil && n2 == nil {
		return nil
	} else if n1 != nil {
		return n1
	} else if n2 != nil {
		return n2
	}
	return nil
}

func (l *Lca) lowestCommonAncestor(node1, node2 *Lca) *Lca {
	if (l == node1) || l == node2 || l == nil {
		return l
	}
	var ls *Lca
	var rs *Lca
	ls, rs = nil, nil
	if l.left != nil {
		ls = l.left.lowestCommonAncestor(node1, node2)
	}
	if l.right != nil {
		rs = l.right.lowestCommonAncestor(node1, node2)
	}
	if ls == nil {
		return rs
	} else if rs == nil {
		return ls
	} else {
		return l
	}
}

func LCA() {
	node := Lca{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(18)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	node.Insert(9)
	fmt.Println("The lowest common ancestor of the tree:")
	n1 := node.FindNode(9)
	n2 := node.FindNode(6)
	n3 := node.lowestCommonAncestor(n1, n2)
	if n3 != nil {
		fmt.Println(n3.val)
	} else {
		fmt.Println("The lowest common ancestor is nil")
	}
}
