package Util

import "math"

func CalculatePages(arr []int, K int) int {
	max := func(arr []int) int {
		temp := math.MinInt
		for i := 0; i < len(arr); i++ {
			if temp > arr[i] {
				temp = arr[i]
			}
		}
		return temp
	}
	allocate := func(arr []int, mid int) int {
		num := 0
		sum := 0
		count := 0
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
			count += 1
			if sum > mid {
				num += 1
				sum = mid
			}
			if count == len(arr) {
				num += 1
			}
		}
		return num
	}
	sum := func(arr []int) int {
		temp := 0
		for i := 0; i < len(arr); i++ {
			temp += arr[i]
		}
		return temp
	}
	high := sum(arr)
	low := max(arr)
	num := 0
	for low < high {
		mid := (low + high) / 2
		num = allocate(arr, mid)
		if num < K {
			high = mid - 1
		} else if num > K {
			low = mid + 1
		} else if num == K {
			high = mid
		}
	}
	return high
}
