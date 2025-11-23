package Util

import (
	"fmt"
	"math"
)

func MinTimeMItem(arr []int, m, n int) int {
	fmt.Println("Finding the minimum time required to produce m items:")
	minm_time := math.MaxInt
	number_items := 0
	index := 0
	time := 0
	for {
		for index < n {
			number_items += time / arr[index]
			index += 1
		}
		if number_items >= m {
			if minm_time > time {
				minm_time = time
			}
			break
		}
		index = 0
		number_items = 0
		time += 1
	}
	return minm_time
}

func MinmTime(arr []int, m, n int) int {
	minm_time := 1
	total := 0
	max_time := m * arr[n-1]
	result := math.MaxInt
	for minm_time < max_time {
		mid := (minm_time + max_time) / 2
		for i := 0; i < n; i++ {
			total += mid / arr[i]
		}
		if total < m {
			minm_time = mid + 1
			total = 0
		} else {
			max_time = mid
			if result > mid {
				result = mid
			}
			total = 0
		}
	}
	return result
}
