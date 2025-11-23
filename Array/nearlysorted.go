package Array

func NearlySorted(arr []int, K int) int {
	start_index := 0
	end_index := 0
	for start_index <= end_index {
		mid := (start_index + end_index) / 2
		if arr[mid] == K {
			return mid
		} else if arr[mid+1] == K {
			return mid + 1
		} else if arr[mid-1] == K {
			return mid - 1
		}

		if arr[mid] > K {
			end_index = mid - 2
		} else if arr[mid] < K {
			start_index = mid + 2
		}
	}
	return -1
}
