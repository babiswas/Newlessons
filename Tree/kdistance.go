package BST

import (
	"container/list"
	"fmt"
)

type Kdistance struct {
	val   int
	left  *Kdistance
	right *Kdistance
}

func (k *Kdistance) Insert(val int) {
	if k.val > val {
		if k.left != nil {
			k.left.Insert(val)
		} else {
			k.left = &Kdistance{val, nil, nil}
		}
	} else {
		if k.right != nil {
			k.right.Insert(val)
		} else {
			k.right = &Kdistance{val, nil, nil}
		}
	}
}

func (k *Kdistance) Convert() map[*Kdistance]*list.List {
	queue := list.New()
	queue.PushBack(k)
	adjList := make(map[*Kdistance]*list.List)
	for queue.Len() != 0 {
		node := queue.Front()
		ptr := node.Value.(*Kdistance)
		if ptr.left != nil {
			queue.PushBack(ptr.left)
			_, ok := adjList[ptr]
			if !ok {
				adjList[ptr] = list.New()
			}
			adjList[ptr].PushBack(ptr.left)
			_, ok = adjList[ptr.left]
			if !ok {
				adjList[ptr.left] = list.New()
			}
			adjList[ptr.left].PushBack(ptr)
		}
		if ptr.right != nil {
			queue.PushBack(ptr.right)
			_, ok := adjList[ptr]
			if !ok {
				adjList[ptr] = list.New()
			}
			adjList[ptr].PushBack(ptr.right)
			_, ok = adjList[ptr.right]
			if !ok {
				adjList[ptr.right] = list.New()
			}
			adjList[ptr.right].PushBack(ptr)
		}
		queue.Remove(node)
	}
	return adjList
}

func (k *Kdistance) FindKDistance(adjList map[*Kdistance]*list.List, node *Kdistance, K int) {
	visited := make(map[*Kdistance]bool)
	for key, _ := range adjList {
		visited[key] = false
	}
	queue := list.New()
	queue.PushBack(node)
	visited[node] = true
	level := 0
	for level != K {
		num := queue.Len()
		level += 1
		for num != 0 {
			node_ptr := queue.Front()
			temp_list := adjList[node_ptr.Value.(*Kdistance)]
			if level == K {
				fmt.Println(node_ptr.Value.(*Kdistance).val)
			}
			queue.Remove(node_ptr)
			for elm := temp_list.Front(); elm != nil; elm = elm.Next() {
				new_node := elm.Value.(*Kdistance)
				if !visited[new_node] {
					queue.PushBack(new_node)
					visited[new_node] = true
				}
			}
			num -= 1
		}
		if level == K {
			break
		}
	}
}

func KdistanceNodes(K int) {
	node := Kdistance{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	node.Insert(2)
	node.Insert(1)
	node.Insert(12)
	node.Insert(13)
	mp := node.Convert()
	node.FindKDistance(mp, node.left, 4)
}
