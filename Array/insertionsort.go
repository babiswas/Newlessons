package Array

func InsertionSort(arr []int, start, end int) {
	for i := start + 1; i <= end; i++ {
		temp := arr[i]
		hole := i
		for (arr[hole-1] > temp) && (hole > 0) {
			arr[hole] = arr[hole-1]
			arr[hole-1] = temp
			hole = hole - 1
			if hole <= 0 {
				break
			}
		}
		arr[hole] = temp
	}
}
