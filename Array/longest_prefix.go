package arr

func LongestPrefix(arr []string) string {
	result := ""
	for i := range len(arr[0]) {
		for _, val := range arr {
			if string(val[i]) != string(arr[0][i]) {
				return result
			}
		}
		result += string(arr[0][i])
	}
	return result
}
