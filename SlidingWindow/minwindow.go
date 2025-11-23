package SlidingWindow

import (
	"fmt"
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

func MinWindowSubstr(str string, substr string) int {
	temp_lngth := math.MaxInt
	temp_str := strings.Split(str, "")
	index1 := 0
	index2 := 0
	mp := make(map[string]int)
	tp := make(map[string]int)
	start_index := 0
	end_index := 0
	for _, val := range substr {
		temp_str := string(val)
		mp[temp_str] += 1
	}
	for _, val := range substr {
		temp_str := string(val)
		tp[temp_str] = 0
	}
	for end_index < len(str) {
		new_str := temp_str[end_index]
		_, ok := tp[new_str]
		if ok {
			tp[new_str] += 1
		}
		if isMatching(mp, tp) {
			for isMatching(mp, tp) {
				if temp_lngth > (end_index - start_index + 1) {
					temp_lngth = end_index - start_index + 1
					index1 = start_index
					index2 = end_index
				}
				newtemp := temp_str[start_index]
				start_index += 1
				_, ok = tp[newtemp]
				if ok {
					tp[newtemp] -= 1
				}
			}
		}
		end_index += 1
	}
	fmt.Println(strings.Join(temp_str[index1:index2+1], ""))
	return temp_lngth
}
