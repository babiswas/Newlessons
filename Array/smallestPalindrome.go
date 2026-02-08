package algo

import (
	"fmt"
	"strconv"
	"strings"
)

func is_palindrome(str string) bool {
	temp_str := strings.Split(str, "")
	start_index := 0
	end_index := 0
	for start_index < end_index {
		if temp_str[start_index] == temp_str[end_index] {
			start_index += 1
			end_index -= 1
		} else {
			return false
		}
	}
	return true
}

func isAllNines(str string) bool {
	temp_str := strings.Split(str, "")
	for _, val := range temp_str {
		if val != "9" {
			return false
		}
	}
	return true
}

func SmallestPalindrome(num int) string {
	num_str := strconv.Itoa(num)

	reverse := func(str string) string {
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

	mid_index := len(num_str) / 2

	first_part := string(num_str[0:mid_index])
	mid := ""
	last_part := ""

	if len(num_str)%2 != 0 {
		mid = string(num_str[mid_index])
		last_part = string(num_str[mid_index+1:])
	} else {
		last_part = string(num_str[mid_index:])
	}

	last_num, _ := strconv.Atoi(last_part)

	if isAllNines(num_str) && is_palindrome(num_str) {
		num, _ := strconv.Atoi(num_str)
		num += 1
		new_num_str := strconv.Itoa(num)
		fmt.Println(new_num_str)

		if len(new_num_str)%2 == 0 {
			mid_index := len(new_num_str) / 2
			first_part := string(new_num_str[:mid_index])
			last_part := reverse(first_part)
			return first_part + last_part
		} else {
			mid_index := len(new_num_str) / 2
			first_part := string(new_num_str[:mid_index])
			mid := string(new_num_str[mid_index])
			last_part := reverse(first_part)
			return first_part + mid + last_part
		}
	} else {
		if len(num_str)%2 == 0 {
			if strings.Contains(first_part, "9") {
				last_part_num := reverse(first_part)
				new_last_num, _ := strconv.Atoi(last_part_num)
				if new_last_num > last_num {
					return first_part + last_part_num
				} else {
					first_part_num, _ := strconv.Atoi(first_part)
					first_part_num += 1
					first_part = strconv.Itoa(first_part_num)
					last_part = reverse(first_part)
					return first_part + last_part
				}
			} else {
				fmt.Println("We are in not all 9 block:")
				new_last_part := reverse(first_part)
				num_last, _ := strconv.Atoi(new_last_part)
				fmt.Println("New last part:", new_last_part)
				fmt.Println("Last num", last_num)
				if num_last > last_num {
					return first_part + new_last_part
				} else {
					first_mid_num, _ := strconv.Atoi(first_part)
					first_mid_num += 1
					newnum := strconv.Itoa(first_mid_num) + last_part
					new_mid_index := len(newnum) / 2
					first_part = string(newnum[:new_mid_index])
					new_last_part := reverse(first_part)
					return first_part + new_last_part
				}
			}
		} else {
			fmt.Println("Odd string block")
			if mid == "9" {
				new_last_part := reverse(first_part)
				num_last, _ := strconv.Atoi(new_last_part)
				if num_last > last_num {
					return first_part + mid + new_last_part
				} else {
					first_mid := first_part + mid
					first_mid_num, _ := strconv.Atoi(first_mid)
					first_mid_num += 1
					newnum := strconv.Itoa(first_mid_num) + last_part
					new_mid_index := len(newnum) / 2
					mid = string(newnum[new_mid_index])
					first_part = string(newnum[:new_mid_index])
					last_part = string(newnum[new_mid_index+1:])
					new_last_part := reverse(first_part)
					return first_part + mid + new_last_part
				}
			} else {
				new_last_part := reverse(first_part)
				num_last, _ := strconv.Atoi(new_last_part)
				if num_last > last_num {
					return first_part + mid + new_last_part
				} else if last_num > num_last {
					mid_num, _ := strconv.Atoi(mid)
					mid_num += 1
					mid := strconv.Itoa(mid_num)
					return first_part + mid + new_last_part
				}

			}
		}
	}
	return ""
}
