package tree

import (
	"container/list"
	"fmt"
)

type Leftview struct {
	val   int
	left  *Leftview
	right *Leftview
}

func (l *Leftview) Insert(val int) {
	if l.val > val {
		if l.left != nil {
			l.left.Insert(val)
		} else {
			l.left = &Leftview{val, nil, nil}
		}
	} else {
		if l.right != nil {
			l.right.Insert(val)
		} else {
			l.right = &Leftview{val, nil, nil}
		}
	}
}

func (l *Leftview) Leftview() {
	queue := list.New()
	ptr := l
	queue.PushBack(ptr)
	for queue.Len() > 0 {
		elm := queue.Front()
		ptr = elm.Value.(*Leftview)
		fmt.Println(ptr.val)
		num := queue.Len()
		for num != 0 {
			elm = queue.Front()
			ptr = elm.Value.(*Leftview)
			queue.Remove(elm)
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
		}
	}
}

func (l *Leftview) RightView() {
	queue := list.New()
	ptr := l
	queue.PushBack(ptr)
	for queue.Len() > 0 {
		num := queue.Len()
		ptr := queue.Back().Value.(*Leftview)
		fmt.Println(ptr.val)
		for num != 0 {
			elm := queue.Front()
			ptr = elm.Value.(*Leftview)
			queue.Remove(elm)
			if ptr.left != nil {
				queue.PushBack(ptr.left)
			}
			if ptr.right != nil {
				queue.PushBack(ptr.right)
			}
			num -= 1
		}
	}
}

func LeftViewTree() {
	node := Leftview{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("The left view of the tree is:")
	node.Leftview()
}

func RightViewTree() {
	node := Leftview{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("The left view of the tree is:")
	node.RightView()
}
