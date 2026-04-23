package arr

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var result = make(map[int]bool)

func construct_num(arr string) int {
	result, _ := strconv.Atoi(arr)
	return result
}

func CollectGreater(temp string, N int, start, end int) {
	if N == 0 {
		num := construct_num(temp)
		result[num] = true
		return
	}

	for i := start; i <= end; i++ {
		last_num := -1
		if temp != "" {
			last_num, _ = strconv.Atoi(strings.Split(temp, "")[len(temp)-1])
		}
		if (temp == "" && i != 0) || ((temp != "") && i > last_num) {
			temp += strconv.Itoa(i)
			CollectGreater(temp, N-1, i+1, end)
			temp = strings.Join(strings.Split(temp, "")[:len(temp)-1], "")
		}
		CollectGreater(temp, N, i+1, end)
	}
}

func FindNumbers() {
	fmt.Println("Find the list of numbers with increasing digit.")
	CollectGreater("", 3, 0, 9)
	result_list := make([]int, 0)
	for key, _ := range result {
		result_list = append(result_list, key)
	}
	sort.Ints(result_list)
	fmt.Println(result_list)
}
