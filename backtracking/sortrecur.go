package arr

import "fmt"

func Insert(arr []int, num int, index int) {
	if index == 0 {
		arr[index] = num
		return
	} else if arr[index-1] <= num {
		arr[index] = num
		return
	}
	temp := arr[index-1]
	Insert(arr, num, index-1)
	arr[index] = temp
}

func Sort(arr []int, index int) {
	if index == -1 {
		return
	}
	temp := arr[index]
	Sort(arr, index-1)
	Insert(arr, temp, index)
}

func Sorting() {
	fmt.Println("Sorting numbers using recursion.")
	mrr := []int{5, 2, 6, 8, 1, 3}
	Sort(mrr, len(mrr)-1)
	fmt.Println(mrr)
}
