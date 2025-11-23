package Array

func SearchRotated(arr []int, start, end, key int) int {
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > arr[start] {
			if key >= arr[start] && key <= arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if key <= arr[end] && key >= arr[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
