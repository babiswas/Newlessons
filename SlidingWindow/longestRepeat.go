package algo

import (
	"fmt"
	"math"
	"strings"
)

func LongestRepeatReplace(str string, K int) int {
	temp_str := strings.Split(str, "")
	max_frequency := func(mp map[string]int) int {
		temp := 0
		for _, val := range mp {
			if temp < val {
				temp = val

			}
		}
		return temp
	}
	mp := make(map[string]int)
	start_index := 0
	end_index := 0
	temp_length := math.MinInt
	for end_index < len(temp_str) {
		newstr := temp_str[end_index]
		_, ok := mp[newstr]
		if !ok {
			mp[newstr] = 1
		}
		if ok {
			mp[newstr] += 1
		}
		max_freq := max_frequency(mp)
		fmt.Println(max_freq)
		window_lngth := end_index - start_index + 1
		for (window_lngth - max_freq) > K {
			temp := temp_str[start_index]
			start_index += 1
			_, ok = mp[temp]
			if ok {
				mp[temp] -= 1
				if mp[temp] == 0 {
					delete(mp, temp)
				}
			}
			window_lngth = end_index - start_index + 1
			max_freq = max_frequency(mp)
		}
		window_lngth = end_index - start_index + 1
		if temp_length < window_lngth {
			temp_length = window_lngth
		}
		end_index += 1
	}
	return temp_length
}
