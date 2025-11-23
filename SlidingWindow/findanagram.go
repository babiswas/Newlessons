package SlidingWindow

import "strings"

func FindAnagram(str string, substr string) []int {
	index_list := make([]int, 0)
	temp_str := strings.Split(str, "")
	start_index := 0
	end_index := 0
	mp := make(map[string]int)
	for _, val := range strings.Split(substr, "") {
		mp[val] += 1
	}
	ismatching := func(mp, tp map[string]int) bool {
		for key, val := range mp {
			_, ok := tp[key]
			if !ok {
				return false
			}
			if tp[key] != val {
				return false
			}
		}
		return true
	}

	tp := make(map[string]int)

	for _, val := range strings.Split(substr, "") {
		tp[val] = 0
	}

	for end_index < len(str) {
		newstr := temp_str[end_index]
		_, ok := tp[newstr]
		if ok {
			tp[newstr] += 1
		}
		if end_index-start_index+1 == len(substr) {
			if ismatching(mp, tp) {
				for ismatching(mp, tp) {
					index_list = append(index_list, start_index)
					newtemp := temp_str[start_index]
					start_index += 1
					_, ok := tp[newtemp]
					if ok {
						tp[newtemp] -= 1
					}
				}
			} else {
				newtemp := temp_str[start_index]
				start_index += 1
				_, ok := tp[newtemp]
				if ok {
					tp[newtemp] -= 1
				}
			}
		}
		end_index += 1
	}
	return index_list
}
