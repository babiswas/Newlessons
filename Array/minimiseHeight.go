package algo

import (
	"sort"
)

func minm(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxm(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinimiseHeight(arr []int, K int) int {
	sort.Ints(arr)
	min_diff := arr[len(arr)-1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-K < 0 {
			continue
		}
		max_num := maxm(arr[i-1]+K, arr[len(arr)-1]-K)
		min_num := min(arr[i]-K, arr[0]+K)
		min_diff = min(min_diff, max_num-min_num)
	}
	return min_diff
}
