package arr

import (
	"container/list"
	"fmt"
)

var pal_list = list.New()

func is_palindrome(str string) bool {
	start := 0
	end := len(str) - 1
	for start <= end {
		if str[start] == str[end] {
			start += 1
			end -= 1
		} else {
			return false
		}
	}
	return true
}

func PalindromePartition(str string, start, end int, temp []string) {
	if start > end {
		pal_list.PushBack(temp)
		return
	}
	for i := start; i <= end; i++ {
		temp_str := str[start : i+1]
		if is_palindrome(temp_str) {
			temp = append(temp, temp_str)
			PalindromePartition(str, i+1, end, temp)
			temp = temp[:len(temp)-1]
		}
	}
}

func PalPart() {
	fmt.Println("Performing palindromic partition:")
	test := "aaba"
	temp := make([]string, 0)
	PalindromePartition(test, 0, len(test)-1, temp)
	for el := pal_list.Front(); el != nil; el = el.Next() {
		lst := el.Value.([]string)
		fmt.Println(lst)
	}
}
