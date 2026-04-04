package arr

func Concatenate(arr []int, n int) []int {
	arr_len := len(arr)
	new_arr := make([]int, arr_len*n)
	counter := 0
	for index := range new_arr {
		new_arr[index] = arr[counter]
		counter += 1
		if counter >= len(arr) {
			counter = 0
		}
	}
	return new_arr
}
