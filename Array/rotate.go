package Array

func RotateArray(arr []int, K int) {
	reverse := func(arr []int, start, end int) {
		start_index := start
		end_index := end
		for start_index <= end_index {
			arr[start_index], arr[end_index] = arr[end_index], arr[start_index]
			start_index += 1
			end_index -= 1
		}
	}
	reverse(arr, 0, K-1)
	reverse(arr, K, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
}
