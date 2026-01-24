package BST

import "fmt"

type ArrayBST struct {
	val   int
	left  *ArrayBST
	right *ArrayBST
}

func CreateBTree(arr []int, start, end int) *ArrayBST {
	if start > end {
		return nil
	}
	midIndex := (start + end) / 2
	node := ArrayBST{arr[midIndex], nil, nil}
	node.left = CreateBTree(arr, start, midIndex-1)
	node.right = CreateBTree(arr, midIndex+1, end)
	return &node
}

func inorder_trav(node *ArrayBST) {
	if node == nil {
		return
	}
	if node.left != nil {
		inorder_trav(node.left)
	}
	fmt.Println(node.val)
	if node.right != nil {
		inorder_trav(node.right)
	}
}

func BinaryTree() {
	arr := []int{10, 17, 5, 16, 8, 51, 6, 3, 7}
	fmt.Println("Creating a binary tree:")
	node := CreateBTree(arr, 0, len(arr)-1)
	fmt.Println("Inorder traversal of the tree:")
	inorder_trav(node)
}
