package SlidingWindow

import (
	"fmt"
	"math"
	"strings"
)

func LongestRepeat(str string) int {
	index1 := 0
	index2 := 0
	temp_str := strings.Split(str, "")
	start_index := 0
	end_index := 0
	temp_len := math.MinInt
	mp := make(map[string]int)
	for end_index < len(str) {
		temp_string := temp_str[end_index]
		_, ok := mp[temp_str[end_index]]
		if !ok {
			mp[temp_string] = 1
			if temp_len < len(mp) {
				temp_len = len(mp)
				index1 = start_index
				index2 = end_index
			}
		} else {
			for mp[temp_string] != 0 {
				new_string := temp_str[start_index]
				start_index += 1
				mp[new_string] -= 1
				if new_string == temp_string {
					delete(mp, new_string)
					break
				}
				if mp[new_string] == 0 {
					delete(mp, new_string)
				}
			}
			mp[temp_string] += 1
			if temp_len < len(mp) {
				temp_len = len(mp)
				index1 = start_index
				index2 = end_index
			}
		}
		end_index += 1
	}
	temp := strings.Join(temp_str[index1:index2+1], "")
	fmt.Println(temp)
	return temp_len
}
