package tree

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (n *Node) Insert(val int) {
	if n.val > val {
		if n.left == nil {
			n.left = &Node{val, nil, nil}
		} else {
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val, nil, nil}
		} else {
			n.right.Insert(val)
		}
	}
}

func (n *Node) InorderTraversal() {
	if n.left != nil {
		n.left.InorderTraversal()
	}
	fmt.Println(n.val)
	if n.right != nil {
		n.right.InorderTraversal()
	}
}

func (n *Node) PreorderTrav() {
	fmt.Println(n.val)
	if n.left != nil {
		n.left.PreorderTrav()
	}
	if n.right != nil {
		n.right.PreorderTrav()
	}
}

func (n *Node) PostorderTraversal() {
	if n.left != nil {
		n.left.PostorderTraversal()
	}
	if n.right != nil {
		n.right.PostorderTraversal()
	}
	fmt.Println(n.val)
}

func DisplayInorderTraversal() {
	node := Node{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Inorder traversal of the tree is:")
	node.InorderTraversal()
}

func DisplayPreorderTraversal() {
	node := Node{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Preorder traversal of the tree:")
	node.PreorderTrav()
}

func DisplayPostorderTraversal() {
	node := Node{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Postorder traversal of the tree:")
	node.PostorderTraversal()
}
