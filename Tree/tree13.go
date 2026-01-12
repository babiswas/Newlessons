package tree

import "fmt"

type BST struct {
	val   int
	left  *BST
	right *BST
}

func TreeNode(val int) *BST {
	return &BST{val, nil, nil}
}

func CreateTree(arr []int, start int, end int) *BST {
	if start > end {
		return nil
	}
	mid_index := (start + end) / 2
	node := TreeNode(arr[mid_index])
	node.left = CreateTree(arr, start, mid_index-1)
	node.right = CreateTree(arr, mid_index+1, end)
	return node
}

func (b *BST) Inorder() {
	if b == nil {
		return
	}
	if b.left != nil {
		b.left.Inorder()
	}
	fmt.Println(b.val)
	if b.right != nil {
		b.right.Inorder()
	}
}

func InorderTrav() {
	arr := []int{11, 22, 33, 44, 55, 66, 77}
	node := CreateTree(arr, 0, len(arr)-1)
	node.Inorder()
}
