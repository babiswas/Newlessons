package algo

import (
	"math"
	"strings"
)

func isMatching(mp, tp map[string]int) bool {
	for key, val := range mp {
		if tp[key] < val {
			return false
		}
	}
	return true
}

func MinWindowSubstr(str string, substr string) (int, int, int) {
	temp_str := strings.Split(substr, "")
	start_index := 0
	end_index := 0
	index1 := 0
	index2 := 0
	min_window := math.MaxInt
	mp := make(map[string]int)
	tp := make(map[string]int)
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
		if isMatching(mp, tp) {
			for isMatching(mp, tp) {
				if end_index-start_index+1 < min_window {
					min_window = end_index - start_index + 1
					index1 = start_index
					index2 = end_index
				}
				temp := string(str[start_index])
				start_index += 1
				_, ok = tp[temp]
				if ok {
					tp[temp] -= 1
				}
			}
		}
		end_index += 1
	}
	return min_window, index1, index2
}
