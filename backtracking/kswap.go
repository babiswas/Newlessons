package arr

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func findmax(arr []int, start, end int) int {
	max := math.MinInt
	for j := start; j <= end; j++ {
		if arr[j] > max {
			max = arr[j]
		}
	}
	return max
}

func create_num(arr *[]int) int {
	sum := ""
	for _, val := range *arr {
		sum += strconv.Itoa(val)
	}
	number, _ := strconv.Atoi(sum)
	return number
}

func find_max(temp []int, K int, start int, end int, maxnum *int) {
	if (start == end) || (K == 0) {
		return
	}
	for i := start + 1; i <= end; i++ {
		mxnum := findmax(temp, i, end)
		if temp[start] < temp[i] && temp[i] == mxnum {
			temp[start], temp[i] = temp[i], temp[start]
			number := create_num(&temp)
			if number > (*maxnum) {
				*maxnum = number
			}
			find_max(temp, K-1, start+1, end, maxnum)
			temp[start], temp[i] = temp[i], temp[start]
		}
		find_max(temp, K, start+1, end, maxnum)
	}
}

func KSwap(str string, K int) {
	temp := make([]int, len(str))
	str_arr := strings.Split(str, "")
	for i := 0; i < len(str); i++ {
		temp[i], _ = strconv.Atoi(str_arr[i])
	}
	maxnum := -1
	find_max(temp, K, 0, len(str)-1, &maxnum)
	fmt.Println("Maximum number:", maxnum)
}
