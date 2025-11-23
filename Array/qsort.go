package Array

func partition(arr []int, start, end int) int {
	wall := start - 1
	pivot := arr[end]
	for i := start; i <= end; i++ {
		if arr[i] < pivot {
			wall += 1
			arr[wall], arr[i] = arr[i], arr[wall]
		}
	}
	wall += 1
	arr[wall], arr[end] = arr[end], arr[wall]
	return wall
}

func Qsort(arr []int, start, end int) {
	if start <= end {
		index := partition(arr, start, end)
		Qsort(arr, start, index-1)
		Qsort(arr, index+1, end)
	}
}
