package Util

func FindPeakElement(arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if (mid < len(arr)-1) && (mid > 0) {
			if (arr[mid] > arr[mid-1]) && (arr[mid] > arr[mid+1]) {
				return mid
			} else {
				if arr[mid] < arr[mid+1] {
					low = mid + 1
				} else if arr[mid] < arr[mid-1] {
					high = mid - 1
				}
			}
		} else if mid == 0 {
			if arr[mid] > arr[mid+1] {
				return mid
			} else {
				low = mid + 1
			}
		} else if mid == len(arr)-1 {
			if arr[mid] > arr[mid-1] {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
