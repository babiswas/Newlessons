package SlidingWindow

import "math"

func BestTimeSellBuy(arr []int) int {
	left, right := 0, 0
	max_profit := math.MinInt
	for right < len(arr) {
		if arr[left] > arr[right] {
			left = right
			right += 1
			continue
		}
		profit := arr[right] - arr[left]
		if profit > max_profit {
			max_profit = profit
		}
		right += 1
	}
	return max_profit
}
