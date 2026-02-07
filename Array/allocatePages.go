package algo

import "math"

func AllocatePages(arr []int, m int) int {
	if len(arr) < m {
		return -1
	}

	find_students := func(arr []int, mid int) int {
		count := 0
		total := 0
		for i := 0; i < len(arr); i++ {
			total += arr[i]
			if total > mid {
				count += 1
				total = arr[i]
			}
		}
		return count + 1
	}
	sum := func(arr []int) int {
		add := 0
		for i := 0; i < len(arr); i++ {
			add += arr[i]
		}
		return add
	}
	max := func(arr []int) int {
		maxm := math.MinInt
		for i := 0; i < len(arr); i++ {
			if maxm < arr[i] {
				maxm = arr[i]
			}
		}
		return maxm
	}
	low := max(arr)
	high := sum(arr)
	for low < high {
		mid := low + (high-low)/2
		num_students := find_students(arr, mid)
		if num_students == m {
			high = mid
		} else if num_students > m {
			low = mid + 1
		} else if num_students < m {
			high = mid - 1
		}
	}
	return high
}
