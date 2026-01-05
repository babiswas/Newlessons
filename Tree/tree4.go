package tree

import (
	"container/list"
	"fmt"
)

type ItrativeInorder struct {
	val   int
	left  *ItrativeInorder
	right *ItrativeInorder
}

func (i *ItrativeInorder) Insert(val int) {
	if i.val > val {
		if i.left != nil {
			i.left.Insert(val)
		} else {
			i.left = &ItrativeInorder{val, nil, nil}
		}
	} else {
		if i.right != nil {
			i.right.Insert(val)
		} else {
			i.right = &ItrativeInorder{val, nil, nil}
		}
	}
}

func (i *ItrativeInorder) Inorder() {
	stack := list.New()
	ptr := i
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else if stack.Len() > 0 {
			elm := stack.Back()
			ptr = elm.Value.(*ItrativeInorder)
			fmt.Println(ptr.val)
			stack.Remove(elm)
			ptr = ptr.right
		} else {
			break
		}
	}
}

func (i *ItrativeInorder) Ipostorder() {
	stack := list.New()
	ptr := i
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else {
			if stack.Len() == 0 {
				break
			}
			if stack.Back().Value.(*ItrativeInorder).right == nil {
				elm := stack.Back()
				ptr = elm.Value.(*ItrativeInorder)
				fmt.Println(ptr.val)
				stack.Remove(elm)
				for stack.Back().Value.(*ItrativeInorder).right == ptr {
					elm = stack.Back()
					ptr = elm.Value.(*ItrativeInorder)
					fmt.Println(ptr.val)
					stack.Remove(elm)
					if stack.Len() == 0 {
						break
					}
				}
			}
			if stack.Len() != 0 {
				ptr = stack.Back().Value.(*ItrativeInorder).right
			} else {
				break
			}
		}
	}
}

func IterInorder() {
	node := ItrativeInorder{15, nil, nil}
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	fmt.Println("Iterative inorder traversal:")
	node.Inorder()
}

func (i *ItrativeInorder) IPreorder() {
	stack := list.New()
	ptr := i
	stack.PushBack(ptr)
	for stack.Len() > 0 {
		elm := stack.Back()
		node := elm.Value.(*ItrativeInorder)
		stack.Remove(elm)
		fmt.Println(node.val)
		if node.right != nil {
			stack.PushBack(node.right)
		}
		if node.left != nil {
			stack.PushBack(node.left)
		}
	}
}

func IterativePreorder() {
	node := ItrativeInorder{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Iterative preorder traversal.")
	node.IPreorder()
}

func PostorderIterative() {
	fmt.Println("Post order traversal of the tree:")
	node := ItrativeInorder{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	node.Ipostorder()
}
