package Util

import (
	"fmt"
	"log"
	"sort"
)

//Sort numbers

func SortPkExample() {
	log.Println("Sorting integers:")
	arr := []int{12, 4, 5, 13, 18, 9, 10}
	fmt.Println("Before sorting the array:")
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Println("Before sorting the strings array is:")
	str_arr := []string{"hello", "bello", "mello", "tello", "cello"}
	fmt.Println(str_arr)
	sort.Strings(str_arr)
	fmt.Println(str_arr)

	type Person struct {
		name string
		age  int
	}

	person_arr := []Person{
		{"bapan", 32},
		{"tapan", 66},
		{"capan", 22},
		{"dapan", 45},
	}
	fmt.Println("Person array before sorting:", person_arr)
	sort.Slice(person_arr, func(i, j int) bool {
		return person_arr[i].age < person_arr[j].age
	})
	fmt.Println("Person array after sorting:", person_arr)
	fmt.Println(person_arr)
}
