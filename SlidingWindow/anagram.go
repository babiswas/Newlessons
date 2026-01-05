package algo

import "strings"

func Anagram(str string, substr string) []int {
	temp_str := strings.Split(substr, "")
	index_list := make([]int, 0)
	start_index := 0
	end_index := 0
	mp := make(map[string]int)
	tp := make(map[string]int)

	is_matching := func(mp, tp map[string]int) bool {
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
	for _, val := range temp_str {
		_, ok := mp[val]
		if !ok {
			mp[val] = 1
		} else {
			mp[val] += 1
		}
	}
	for _, val := range temp_str {
		tp[val] = 0
	}
	for end_index < len(str) {
		newstr := string(str[end_index])
		_, ok := tp[newstr]
		if ok {
			tp[newstr] += 1
		}
		if end_index-start_index+1 == len(substr) {
			if is_matching(mp, tp) {
				index_list = append(index_list, start_index)
			}
			temp := string(str[start_index])
			start_index += 1
			_, ok := tp[temp]
			if ok {
				tp[temp] -= 1
			}
		}
		end_index += 1
	}
	return index_list
}
