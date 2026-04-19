package main

import (
	"fmt"
	"strings"
)

var str_lst = make([][]string, 0)

func Permute(str []string, start, end int) {
	if start == end {
		temp := make([]string, len(str))
		copy(temp, str)
		str_lst = append(str_lst, temp)
		return
	}

	for i := start; i <= end; i++ {
		str[start], str[i] = str[i], str[start]
		Permute(str, start+1, end)
		str[start], str[i] = str[i], str[start]
	}
}

func main() {
	temp_str := "abc"
	temp := strings.Split(temp_str, "")
	Permute(temp, 0, 2)
	fmt.Println(str_lst)
}
