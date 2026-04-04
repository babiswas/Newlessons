package arr

import (
	"fmt"
	"sort"
)

func FindDuplicate(arr []int) int {
	mp := make(map[int]bool)
	for index := range len(arr) {
		ok, _ := mp[arr[index]]
		if !ok {
			mp[arr[index]] = true
		} else {
			return arr[index]
		}
	}
	return -1
}

func FindDuplicatesv2(arr []int) bool {
	status := false
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				status = true
				break
			}
		}
	}
	return status
}

func FindDuplicatesV3(arr []int) bool {
	sort.Ints(arr)
	fmt.Println("After sorting the array:", arr)
	status := false
	counter := 0
	temp := 0
	temp = arr[0]
	counter += 1
	for i := 1; i < len(arr); i++ {
		if temp == arr[i] {
			counter += 1
		} else {
			counter = 1
		}
		if counter > 1 {
			status = true
			break
		}
	}
	return status
}
