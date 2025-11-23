package Util

import "math"

func FindMinmelement(arr []int) int {
	start_index := 0
	last_index := len(arr) - 1
	result := math.MaxInt
	for start_index < last_index {
		mid := (start_index + last_index) / 2
		if arr[mid] > arr[start_index] {
			if result > arr[start_index] {
				result = arr[start_index]
			}
			start_index = mid + 1
		} else if arr[mid] < arr[last_index] {
			if result > arr[mid] {
				result = arr[mid]
			}
			last_index = mid - 1
		}
	}
	return result
}
