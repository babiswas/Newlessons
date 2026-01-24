package BST

import (
	"container/list"
	"fmt"
)

type Postorder struct {
	val   int
	left  *Postorder
	right *Postorder
}

func (p *Postorder) Check() {
	if p.left != nil {
		p.left.Check()
	}
	fmt.Println(p.val)
	if p.right != nil {
		p.right.Check()
	}
}

func (p *Postorder) Insert(val int) {
	if p.val > val {
		if p.left != nil {
			p.left.Insert(val)
		} else {
			p.left = &Postorder{val, nil, nil}
		}
	} else {
		if p.right != nil {
			p.right.Insert(val)
		} else {
			p.right = &Postorder{val, nil, nil}
		}
	}
}

func (p *Postorder) Ptraversal() {
	stack := list.New()
	ptr := p
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else {
			if stack.Len() == 0 {
				break
			}
			if stack.Back().Value.(*Postorder).right == nil {
				elm := stack.Back()
				ptr = elm.Value.(*Postorder)
				fmt.Println(ptr.val)
				stack.Remove(elm)
				for stack.Back().Value.(*Postorder).right == ptr {
					elm = stack.Back()
					ptr = elm.Value.(*Postorder)
					fmt.Println(ptr.val)
					stack.Remove(elm)
					if stack.Len() == 0 {
						break
					}
				}
			}
			if stack.Len() != 0 {
				ptr = stack.Back().Value.(*Postorder).right
			} else {
				break
			}
		}
	}
}

func PostorderTrav() {
	node := Postorder{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(6)
	node.Insert(10)
	fmt.Println("The postorder traversal of the tree is:")
	node.Ptraversal()
}
