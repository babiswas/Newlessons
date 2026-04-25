package arr

import (
	"container/list"
	"fmt"
)

var all_subset = list.New()

func AllSubsets(str string, start, end int, temp []string) {
	if start > end {
		all_subset.PushBack(temp)
		return
	}
	for i := start; i <= end; i++ {
		temp = append(temp, string(str[i]))
		AllSubsets(str, i+1, end, temp)
		temp = temp[:len(temp)-1]
		AllSubsets(str, i+1, end, temp)
	}
}

func GetAllSubset() {
	fmt.Println("All subsets of an arr:")
	temp_str := "abc"
	temp := make([]string, 0)
	fmt.Println("Collecting all subsets:")
	AllSubsets(temp_str, 0, len(temp_str)-1, temp)
	for el := all_subset.Front(); el != nil; el = el.Next() {
		sbst := el.Value.([]string)
		fmt.Println(sbst)
	}
}
