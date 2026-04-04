package arr

import "fmt"

func FindTargetElement(arr []int, sum int) bool {
	status := false
	mp := make(map[int]int)
	for index, val := range arr {
		temp := sum - val
		_, ok := mp[temp]
		if ok {
			status = true
			fmt.Println(temp, val)
			return status
		}
		mp[val] = index
	}
	return status
}
