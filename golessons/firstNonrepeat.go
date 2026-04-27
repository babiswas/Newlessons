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

func FirstNotRepeat() {
	fmt.Println("Find first non repeating character:")
	temp_str := "geeksforgeeks"
	test := FirstNonRepeat(temp_str)
	fmt.Println(test)
}
