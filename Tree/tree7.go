package tree

import (
	"container/list"
	"fmt"
	"sort"
)

type TopView struct {
	val   int
	left  *TopView
	right *TopView
}

type Level struct {
	level int
	ptr   *TopView
}

func (t *TopView) Insert(val int) {
	if t.val > val {
		if t.left != nil {
			t.left.Insert(val)
		} else {
			t.left = &TopView{val, nil, nil}
		}
	} else {
		if t.right != nil {
			t.right.Insert(val)
		} else {
			t.right = &TopView{val, nil, nil}
		}
	}
}

func (t *TopView) VerticalTrav() {
	mp := make(map[int][]int)
	queue := list.New()
	level_obj := Level{0, t}
	queue.PushBack(level_obj)
	node := &TopView{}
	index := 0
	for queue.Len() > 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			level_obj = elm.Value.(Level)
			node = level_obj.ptr
			index = level_obj.level
			_, ok := mp[index]
			if !ok {
				mp[index] = make([]int, 0)
				mp[index] = append(mp[index], node.val)
			} else {
				mp[index] = append(mp[index], node.val)
			}
			queue.Remove(elm)
			if node.left != nil {
				queue.PushBack(Level{index - 1, node.left})
			}
			if node.right != nil {
				queue.PushBack(Level{index + 1, node.right})
			}
			num -= 1
		}
	}
	fmt.Println(mp)
	index_list := make([]int, 0)
	for ix, _ := range mp {
		index_list = append(index_list, ix)
	}
	sort.Ints(index_list)
	fmt.Println("Top view of the tree is:")
	for _, nums := range index_list {
		fmt.Println(mp[nums][0])
	}
	fmt.Println("Bottom view of the tree is:")
	for _, nums := range index_list {
		fmt.Println(mp[nums][len(mp[nums])-1])
	}
}

func TopAndBottom() {
	node := TopView{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("Top view and bootom view of the tree:")
	node.VerticalTrav()
}
