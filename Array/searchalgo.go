package Array

func BinarySearch(arr []int, key int) int {
	start_index := 0
	end_index := len(arr) - 1
	for start_index <= end_index {
		mid := (start_index + end_index) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			end_index = mid - 1
		} else if arr[mid] < key {
			start_index = mid + 1
		}
	}
	return -1
}

func BinarySearchReverse(arr []int, key int) int {
	start_index := 0
	end_index := len(arr) - 1
	for start_index <= end_index {
		mid := (start_index + end_index) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			start_index = mid + 1
		} else if arr[mid] < key {
			end_index = mid - 1
		}
	}
	return -1
}

func LastOccurance(arr []int, key int) int {
	start_index := 0
	last_index := len(arr) - 1
	result := -1
	for start_index <= last_index {
		mid := (start_index + last_index) / 2
		if arr[mid] == key {
			result = mid
			start_index = mid + 1
		} else if arr[mid] > key {
			last_index = mid - 1
		} else if arr[mid] < key {
			start_index = mid + 1
		}
	}
	return result
}

func OrderAgnostic(arr []int, K int) int {
	if arr[0] > arr[len(arr)-1] {
		return BinarySearchReverse(arr, K)
	} else {
		return BinarySearch(arr, K)
	}
}

func FirstOccurance(arr []int, K int) int {
	start_index := 0
	last_index := len(arr) - 1
	result := -1
	for start_index <= last_index {
		mid := (start_index + last_index) / 2
		if arr[mid] == K {
			result = mid
			last_index = mid - 1
		} else if arr[mid] > K {
			last_index = mid - 1
		} else if arr[mid] < K {
			start_index = mid + 1
		}
	}
	return result
}

func CountNumElements(arr []int, K int) int {
	index1 := FirstOccurance(arr, K)
	index2 := LastOccurance(arr, K)
	return index2 - index1 + 1
}
