package algo

import (
	"math"
	"sort"
)

func MaxFrequency(arr []int, K int) int {
	sort.Ints(arr)
	start_index := 0
	end_index := 0
	result := math.MinInt
	total := 0
	for end_index < len(arr) {
		total += arr[end_index]
		for arr[end_index]*(end_index-start_index+1) > total+K {
			num := arr[start_index]
			total -= num
			start_index += 1
		}
		if result < (end_index - start_index + 1) {
			result = (end_index - start_index + 1)
		}
		end_index += 1
	}
	return result
}
