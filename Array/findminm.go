package Array

import "math"

func Minm(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func FindMinm(arr []int) int {
	low := 0
	high := len(arr) - 1
	ans := math.MaxInt
	for low <= high {
		mid := (low + high) / 2
		if arr[low] <= arr[mid] {
			ans = Minm(ans, arr[low])
			low = mid + 1
		} else {
			ans = Minm(ans, arr[mid])
			high = mid - 1
		}
	}
	return ans
}
