package algo

import "math"

func LongestSubstrRepeat(str string) (int, int, int) {
	mp := make(map[string]int)
	temp_len := math.MinInt
	index1 := 0
	index2 := 0
	start_index := 0
	end_index := 0
	for end_index < len(str) {
		newstr := string(str[end_index])
		_, ok := mp[newstr]
		if !ok {
			mp[newstr] = 1
		} else {
			for mp[newstr] > 0 {
				temp := string(str[start_index])
				start_index += 1
				mp[temp] -= 1
				if mp[temp] == 0 {
					delete(mp, temp)
				}
			}
			mp[newstr] = 1
		}
		if temp_len < (end_index - start_index + 1) {
			temp_len = end_index - start_index + 1
			index1 = start_index
			index2 = end_index
		}
		end_index += 1
	}
	return temp_len, index1, index2
}
