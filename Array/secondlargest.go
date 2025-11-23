package Array

import (
	"sort"
)

func FindSecondLargest(arr []int) int {
	if len(arr) < 2 {
		return arr[0]
	}
	first_largest := arr[0]
	second_largest := arr[1]

	if first_largest < second_largest {
		first_largest, second_largest = second_largest, first_largest
	}

	index := 0
	for index < len(arr) {
		if arr[index] > first_largest {
			first_largest, second_largest = arr[index], first_largest
		} else {
			if arr[index] > second_largest {
				second_largest = arr[index]
			}
		}
		index += 1
	}
	sort.Ints(arr)
	return second_largest
}
