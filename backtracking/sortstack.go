package arr

import "fmt"

func PushStack(arr *[]int, num int) {
	if len(*arr) == 0 {
		*arr = append(*arr, num)
		return
	} else if num >= (*arr)[len(*arr)-1] {
		*arr = append(*arr, num)
		return
	}
	temp := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	PushStack(arr, num)
	*arr = append(*arr, temp)
}

func SortStack(arr *[]int, index int) {
	if index == -1 {
		return
	}
	temp := (*arr)[index]
	*arr = (*arr)[:len(*arr)-1]
	SortStack(arr, index-1)
	PushStack(arr, temp)
}

func StackSorting() {
	fmt.Println("Stack Sorting")
	stack := make([]int, 0)
	arr := []int{5, 4, 2, 9, 6}
	fmt.Println("Sorting the stack:")
	for _, val := range arr {
		stack = append(stack, val)
	}
	SortStack(&arr, len(arr)-1)
	fmt.Println("exit", arr)
}
