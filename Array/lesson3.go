package arr

import "strings"

func IfAnagram(s, t string) bool {
	temp_str := strings.Split(s, "")
	temp_str1 := strings.Split(t, "")

	mp1 := make(map[string]int)
	mp2 := make(map[string]int)

	for _, val := range temp_str {
		_, ok := mp1[val]
		if !ok {
			mp1[val] = 1
		} else {
			mp1[val] += 1
		}
	}

	for _, val := range temp_str1 {
		_, ok := mp2[val]
		if !ok {
			mp2[val] = 1
		} else {
			mp2[val] += 1
		}
	}

	is_matching := func(mp1, mp2 map[string]int) bool {
		status := false
		for key, val := range mp1 {
			_, ok := mp2[key]
			if !ok {
				status = false
				return status
			}
			if mp2[key] < val || mp2[key] > val {
				status = false
				return status
			}
		}
		status = true
		return status
	}

	return is_matching(mp1, mp2)
}
