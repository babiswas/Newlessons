package BST

import (
	"container/list"
	"fmt"
)

type Btree struct {
	val   int
	left  *Btree
	right *Btree
}

func (b *Btree) Insert(val int) {
	if b.val > val {
		if b.left != nil {
			b.left.Insert(val)
		} else {
			b.left = &Btree{val, nil, nil}
		}
	} else {
		if b.right != nil {
			b.right.Insert(val)
		} else {
			b.right = &Btree{val, nil, nil}
		}
	}
}

func (b *Btree) TraverseDLL(head *Btree) {
	var ptr *Btree
	ptr = head
	for ptr != nil {
		fmt.Println(ptr.val)
		ptr = ptr.right
	}
}

func (b *Btree) BtreeDLL() *Btree {
	stack := list.New()
	var head *Btree
	var prev *Btree
	head = nil
	prev = nil
	ptr := b
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else if stack.Len() != 0 {
			node := stack.Back()
			ptr = node.Value.(*Btree)
			stack.Remove(node)
			if head == nil {
				head = ptr
			} else {
				if ptr.left == prev {
					prev.right = ptr
				} else if prev.right == ptr {
					ptr.left = prev
				} else if (ptr.left != prev) && (prev.right != ptr) {
					prev.right = ptr
					ptr.left = prev
				}
			}
			prev = ptr
			ptr = ptr.right
		}
		if stack.Len() == 0 {
			break
		}
	}
	return head
}

func BtreeAsDLL() {
	node := Btree{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("Converting binary tree to DLL:")
	ptr := node.BtreeDLL()
	fmt.Println("Traversing the DLL:")
	node.TraverseDLL(ptr)
}
