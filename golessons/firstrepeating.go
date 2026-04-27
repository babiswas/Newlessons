package arr

import (
	"fmt"
	"math"
	"strings"
)

func FindFirstRepeat(str string) string {
	mp := make(map[string]bool)
	str_list := strings.Split(str, "")
	for _, val := range str_list {
		_, ok := mp[val]
		if !ok {
			mp[val] = true
		} else {
			return val
		}
	}
	return ""
}

func findfirstrepeat(str string) string {
	temp_str := strings.Split(str, "")
	start_index, repeat_index := math.MaxInt, math.MaxInt
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if temp_str[i] == temp_str[j] {
				if repeat_index > j {
					start_index = i
					repeat_index = j
				}
			}
		}
	}
	return temp_str[start_index]
}

func FindFirstRepeating() {
	temp_str := "geeksforgeeks"
	rpt_chr := FindFirstRepeat(temp_str)
	fmt.Println(rpt_chr)
	temp_str1 := "hello geeks"
	rpt_chr = FindFirstRepeat(temp_str1)
	fmt.Println(rpt_chr)

	fmt.Println("SEcond implementation:")
	newtemp := "geeksforgeeks"
	rpt_chr = findfirstrepeat(newtemp)
	fmt.Println(rpt_chr)
}
