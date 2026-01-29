package BST

import (
	"container/list"
	"fmt"
)

type InorderPreorder struct {
	val   int
	left  *InorderPreorder
	right *InorderPreorder
}

func IndexInorder(inorder []int) map[int]int {
	mp_inorder := make(map[int]int)
	for index, val := range inorder {
		mp_inorder[val] = index
	}
	return mp_inorder
}

func CreateNewTree(inorder, preorder []int, mp_inorder map[int]int, istart, iend int, pstart, pend int) *InorderPreorder {
	if (istart > iend) || (pstart > pend) {
		return nil
	} else {
		root := InorderPreorder{}
		root.val = preorder[pstart]
		inorder_index := mp_inorder[root.val]
		nums := inorder_index - istart
		root.left = CreateNewTree(inorder, preorder, mp_inorder, istart, inorder_index-1, pstart+1, pstart+nums)
		root.right = CreateNewTree(inorder, preorder, mp_inorder, inorder_index+1, iend, pstart+nums+1, pend)
		return &root
	}
}

func LevelOrderTrav(root *InorderPreorder) {
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*InorderPreorder)
			queue.Remove(node)
			fmt.Print(ptr.val, " ")
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
		}
		fmt.Println("")
	}
}

func CreateTreeNew() {
	inorder_arr := []int{40, 20, 50, 10, 60, 30}
	preorder_arr := []int{10, 20, 40, 50, 30, 60}
	mp_inorder := IndexInorder(inorder_arr)
	newtree := CreateNewTree(inorder_arr, preorder_arr, mp_inorder, 0, len(inorder_arr)-1, 0, len(preorder_arr)-1)
	fmt.Println("Performing level order traversal:")
	LevelOrderTrav(newtree)
}
