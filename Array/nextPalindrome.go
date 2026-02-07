package algo

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(str string) string {
	temp_str := strings.Split(str, "")
	start_index := 0
	end_index := len(str) - 1
	for start_index < end_index {
		temp_str[start_index], temp_str[end_index] = temp_str[end_index], temp_str[start_index]
		start_index += 1
		end_index -= 1
	}
	return strings.Join(temp_str, "")
}

func NextGreater(str []string) string {
	count := make([]int, 10)
	for i := 0; i < len(str); i++ {
		count[i] = 0
	}
	pos := 0
	digit := 0
	for pos = len(str) - 1; pos >= 0; pos-- {
		temp_num, _ := strconv.Atoi(str[pos])
		count[temp_num] += 1
		if digit > temp_num {
			digit = temp_num
			break
		}
		digit = temp_num
	}
	if pos == -1 {
		return "Impossible"
	}
	for digit = digit + 1; digit < 10; digit++ {
		if count[digit] > 0 {
			temp_str := strconv.Itoa(digit)
			str[pos] = temp_str
			count[digit] -= 1
			pos += 1
			break
		}
	}

	for digit = 0; digit < 10; digit++ {
		for count[digit] > 0 {
			temp_str := strconv.Itoa(digit)
			str[pos] = temp_str
			count[digit] -= 1
			pos += 1
		}
	}
	return strings.Join(str, "")

}

func NextGreaterPalindrome(str string) string {
	temp_nums := strings.Split(str, "")
	mid_index := -1
	mid := ""
	if len(str)%2 == 0 {
		mid = ""
		mid_index = len(str) / 2
	} else {
		mid = temp_nums[len(str)/2]
		mid_index = len(str) / 2
	}

	result := NextGreater(temp_nums[0:mid_index])
	fmt.Println("Obtained result:", result)
	if result == "Impossible" {
		return ""
	} else {
		return result + mid + reverse(result)
	}
}
