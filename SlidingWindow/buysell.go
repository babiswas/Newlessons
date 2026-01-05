package algo

import "math"

func BuySell(arr []int) (int, int, int) {
	start := 0
	end := 0
	buy_price := 0
	sell_price := 0
	maxprofit := math.MinInt
	for end < len(arr) {
		if arr[end] < arr[start] {
			start = end
			end += 1
			continue
		}
		profit := arr[end] - arr[start]
		if profit > maxprofit {
			buy_price = arr[start]
			sell_price = arr[end]
			maxprofit = profit
		}
		end += 1
	}
	return maxprofit, buy_price, sell_price
}
