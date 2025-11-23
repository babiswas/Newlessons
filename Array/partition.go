package Array

func Partition(arr []int, start, end int) (int, []int) {
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
	return wall, arr
}
