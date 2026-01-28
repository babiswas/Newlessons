package BST

import (
	"container/list"
	"fmt"
)

type Merge struct {
	val   int
	left  *Merge
	right *Merge
}

func CreateTree1() *Merge {
	node := Merge{1, nil, nil}
	node.left = &Merge{3, nil, nil}
	node.left.right = &Merge{5, nil, nil}
	node.right = &Merge{2, nil, nil}
	return &node
}

func CreateTree2() *Merge {
	node := Merge{2, nil, nil}
	node.left = &Merge{1, nil, nil}
	node.left.left = &Merge{4, nil, nil}
	node.right = &Merge{3, nil, nil}
	node.right.left = &Merge{7, nil, nil}
	return &node
}

func (m *Merge) MergeTree(tree2 *Merge) *Merge {
	if m != nil && tree2 != nil {
		m.val = m.val + tree2.val
	} else if m == nil && tree2 != nil {
		m = &Merge{tree2.val, nil, nil}
	}

	if m.left != nil && tree2.left != nil {
		m.left.MergeTree(tree2.left)
	} else if m.left == nil && tree2.left != nil {
		m.left = tree2.left
	}

	if m.right != nil && tree2.right != nil {
		m.right.MergeTree(tree2.right)
	} else if m.right == nil && tree2.right != nil {
		m.right = tree2.right
	}

	return m

}

func (m *Merge) LevelOrderTrav() {
	queue := list.New()
	queue.PushBack(m)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			node := queue.Front()
			ptr := node.Value.(*Merge)
			fmt.Print(ptr.val, " ")
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
			queue.Remove(node)
		}
		fmt.Println("")
	}
}

func MergeTrees() {
	tree1 := CreateTree1()
	fmt.Println("Level order traversal of tree1:")
	tree1.LevelOrderTrav()
	tree2 := CreateTree2()
	fmt.Println("Level order traversal of tree2:")
	tree2.LevelOrderTrav()

	fmt.Println("Merging two trees:")
	tree1.MergeTree(tree2)

	fmt.Println("Level order traversal after merge:")
	tree1.LevelOrderTrav()
}
