package BST

import (
	"container/list"
	"fmt"
)

type BurnTree struct {
	val   int
	left  *BurnTree
	right *BurnTree
}

func (b *BurnTree) Insert(val int) {
	if b.val > val {
		if b.left != nil {
			b.left.Insert(val)
		} else {
			b.left = &BurnTree{val, nil, nil}
		}
	} else {
		if b.right != nil {
			b.right.Insert(val)
		} else {
			b.right = &BurnTree{val, nil, nil}
		}
	}
}

func (b *BurnTree) ConvertTreeToGraph() map[*BurnTree]*list.List {
	queue := list.New()
	mp := make(map[*BurnTree]*list.List)
	queue.PushBack(b)
	for queue.Len() != 0 {
		node := queue.Front()
		ptr := node.Value.(*BurnTree)
		if ptr.left != nil {
			queue.PushBack(ptr.left)
			_, ok := mp[ptr.left]
			if !ok {
				mp[ptr.left] = list.New()
			}
			mp[ptr.left].PushBack(ptr)
			_, ok = mp[ptr]
			if !ok {
				mp[ptr] = list.New()
			}
			mp[ptr].PushBack(ptr.left)
		}
		if ptr.right != nil {
			queue.PushBack(ptr.right)
			_, ok := mp[ptr.right]
			if !ok {
				mp[ptr.right] = list.New()
			}
			mp[ptr.right].PushBack(ptr)
			_, ok = mp[ptr]
			if !ok {
				mp[ptr] = list.New()
			}
			mp[ptr].PushBack(ptr.right)
		}
		queue.Remove(node)

	}
	return mp
}

func (b *BurnTree) TimeTakenToBurnTree(node *BurnTree, mp map[*BurnTree]*list.List) int {
	visited := make(map[*BurnTree]bool)
	for key, _ := range mp {
		visited[key] = false
	}
	queue := list.New()
	queue.PushBack(node)
	visited[node] = true
	level := 0
	for queue.Len() != 0 {
		num := queue.Len()
		level += 1
		for num != 0 {
			node := queue.Front()
			temp_list := mp[node.Value.(*BurnTree)]
			queue.Remove(node)
			for elm := temp_list.Front(); elm != nil; elm = elm.Next() {
				new_node := elm.Value.(*BurnTree)
				if !visited[new_node] {
					queue.PushBack(new_node)
					visited[new_node] = true
				}
			}
			num -= 1
		}
	}
	return level
}

func BurnBtree() {
	node := BurnTree{15, nil, nil}
	node.Insert(22)
	node.Insert(26)
	node.Insert(20)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The time taken to burn complete tree:")
	mp := node.ConvertTreeToGraph()
	for key, obj := range mp {
		fmt.Print(key.val)
		fmt.Print("------>")
		tmp_lst := obj
		for elm := tmp_lst.Front(); elm != nil; elm = elm.Next() {
			fmt.Print(elm.Value.(*BurnTree).val)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	time_taken := node.TimeTakenToBurnTree(node.left, mp)
	fmt.Println(time_taken)
}
