package BST

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
	if l.left != nil {
		n1 = l.left.FindNode(num)
	}
	if l.right != nil {
		n2 = l.right.FindNode(num)
	}
	if n1 != nil {
		return n1
	} else if n2 != nil {
		return n2
	} else {
		return nil
	}

}

func (l *Lca) LowestCommonAncestor(node1, node2 *Lca) *Lca {
	if l == node1 || l == node2 || l == nil {
		return l
	}
	var n1 *Lca
	var n2 *Lca
	n1, n2 = nil, nil
	if l.left != nil {
		n1 = l.left.LowestCommonAncestor(node1, node2)
	}
	if l.right != nil {
		n2 = l.right.LowestCommonAncestor(node1, node2)
	}
	if n1 != nil && n2 != nil {
		return l
	} else if n2 != nil && n1 == nil {
		return n2
	} else {
		return n1
	}
}

func FindLCA() {
	node := Lca{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	first_node := node.FindNode(6)
	second_node := node.FindNode(10)
	test := node.LowestCommonAncestor(first_node, second_node)
	fmt.Println("The lowest common ancestor of the tree:")
	fmt.Println(test.val)
}
