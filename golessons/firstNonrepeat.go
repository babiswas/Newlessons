package arr

import (
	"fmt"
	"strings"
)

func FirstNonRepeat(str string) string {
	temp_str := strings.Split(str, "")
	found := true
	for i := 0; i < len(temp_str); i++ {
		found = true
		for j := i + 1; j < len(temp_str); j++ {
			if temp_str[i] == temp_str[j] {
				found = false
			}
		}
		if found == true {
			return temp_str[i]
		}
	}
	return ""
}

func MapImplement(str string) string {
	temp_str := strings.Split(str, "")
	mp := make(map[string]int)
	for _, val := range temp_str {
		_, ok := mp[val]
		if !ok {
			mp[val] = 1
		} else {
			mp[val] += 1
		}
	}
	for _, val := range temp_str {
		if mp[val] == 1 {
			return val
		}
	}
	return ""
}

func FirstNotRepeat() {
	fmt.Println("Find first non repeating character:")
	temp_str := "geeksforgeeks"
	test := FirstNonRepeat(temp_str)
	fmt.Println(test)

	fmt.Println("Map implementstion of the non repeating char:")
	test = MapImplement(temp_str)
	fmt.Println(test)
}
