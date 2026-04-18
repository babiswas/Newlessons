package arr

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(str string) string {
	str_arr := strings.Split(str, "")
	start := 0
	end := len(str) - 1
	for start < end {
		str_arr[start], str_arr[end] = str_arr[end], str_arr[start]
		start += 1
		end -= 1
	}
	return strings.Join(str_arr, "")
}

func FindNextGraterPal(num int) string {
	str := strconv.Itoa(num)
	str_arr := strings.Split(str, "")
	mid := ""
	if len(str)%2 != 0 {
		mid = str_arr[len(str)/2]
	}
	mid_index := len(str) / 2
	next_greater_pal := find_next_greater_pal(str_arr[:mid_index])
	return next_greater_pal + mid + reverse(next_greater_pal)
}

func find_next_greater_pal(str []string) string {
	index_count := make([]int, 10)
	for i := 0; i < 10; i++ {
		index_count[i] = 0
	}
	digit := 0
	pos := len(str) - 1
	for pos = len(str) - 1; pos >= 0; pos-- {
		num, _ := strconv.Atoi(str[pos])
		index_count[num] += 1
		if num < digit {
			digit = num
			break
		}
		digit = num
	}
	if pos == -1 {
		return "impossible."
	}
	for digit += 1; digit < 10; digit++ {
		if index_count[digit] > 0 {
			str[pos] = strconv.Itoa(digit)
			pos += 1
			index_count[digit] -= 1
			break
		}
	}
	for digit = 0; digit < 10; digit++ {
		if index_count[digit] > 0 {
			str[pos] = strconv.Itoa(digit)
			pos += 1
			index_count[digit] -= 1
		}
	}
	return strings.Join(str, "")
}

func NextGraterPalindrome() {
	pal := FindNextGraterPal(4697557964)
	fmt.Println(pal)
}
