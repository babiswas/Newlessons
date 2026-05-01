package arr

import "fmt"

type Stack []int

func (s *Stack) is_empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) push_data(data int) {
	if len(*s) == 5 {
		fmt.Println("Stack overflow.")
		return
	}
	*s = append(*s, data)
}

func CreateStack() *Stack {
	s := make(Stack, 0)
	return &s
}

func (s *Stack) pop_data() int {
	if s.is_empty() {
		fmt.Println("Stack is empty.")
		return -1
	}
	data := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return data
}

func (s *Stack) traverse_stack() {
	for i := len(*s) - 1; i >= 0; i-- {
		fmt.Println((*s)[i])
	}
}

func StackData() {
	fmt.Println("Displaying the stack data structure.")
	stack := CreateStack()
	stack.push_data(21)
	stack.push_data(22)
	stack.push_data(25)
	stack.push_data(32)
	stack.push_data(36)
	fmt.Println("Traversing Stack:")
	stack.traverse_stack()
	fmt.Println("Displaying stack overflow.")
	stack.push_data(25)
	fmt.Println("Emptying the stack.")
	for {
		fmt.Println("Popped data:")
		num := stack.pop_data()
		if num == -1 {
			break
		}
	}
}
