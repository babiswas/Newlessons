package SlidingWindow

import (
	"math"
	"strings"
)

func FindMaxSubSimillar(str string, K int) int {
	maxfreq := func(mp map[string]int) int {
		temp := 0
		for _, value := range mp {
			if temp < value {
				temp = value
			}
		}
		return temp
	}
	start_index := 0
	end_index := 0
	mf := 0
	temp_lngth := math.MinInt
	mp := make(map[string]int)
	temp_str := strings.Split(str, "")
	for end_index < len(str) {
		newtemp := temp_str[end_index]
		_, ok := mp[newtemp]
		if !ok {
			mp[newtemp] = 1
		} else {
			mp[newtemp] += 1
		}

		winl := end_index - start_index + 1
		mf = maxfreq(mp)
		for (winl - mf) > K {
			newstr := temp_str[start_index]
			start_index += 1
			mp[newstr] -= 1
			if mp[newstr] == 0 {
				delete(mp, newstr)
			}
			winl = end_index - start_index + 1
			mf = maxfreq(mp)
		}
		winl = end_index - start_index + 1
		if temp_lngth < winl {
			temp_lngth = winl
		}
		end_index += 1
	}
	return temp_lngth
}
