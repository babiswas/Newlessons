package algo

import (
	"container/list"
	"math"
)

func IsBstPreorder(arr []int) bool {
	stack := list.New()
	root := math.MinInt
	for i := 0; i < len(arr); i++ {
		for stack.Len() != 0 && arr[i] > stack.Back().Value.(int) {
			root = stack.Back().Value.(int)
			stack.Remove(stack.Back())
		}
		if arr[i] < root {
			return false
		}
		stack.PushBack(arr[i])
	}
	return true
}
