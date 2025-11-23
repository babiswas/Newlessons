package main

import (
	"flag"
	"fmt"
	"gogym/Array"
	"gogym/Map"
	"gogym/SlidingWindow"
	"gogym/Util"
)

func main() {
	ops := flag.String("ops", "ssort", "Algorithm to execute")
	flag.Parse()
	algo := *ops
	switch algo {
	case "ssort":
		fmt.Println("Performing selection sort.")
		arr := []int{21, 23, 14, 16, 23, 10, 9}
		Array.SelectionSort(arr)
		fmt.Println(arr)
	case "isort":
		fmt.Println("Performing insertion sort.")
		arr := []int{26, 13, 16, 11, 45, 28, 19}
		Array.InsertionSort(arr, 0, len(arr)-1)
		fmt.Println(arr)
	case "bsort":
		fmt.Println("Performing bubble sort algorithm:")
		arr := []int{12, 18, 11, 23, 28, 29, 10}
		Array.BubbleSort(arr, 0, len(arr))
		fmt.Println(arr)
	case "prt":
		fmt.Println("Performing partition of the arr:")
		arr := []int{13, 18, 25, 10, 26, 11}
		index, arr := Array.Partition(arr, 0, len(arr)-1)
		fmt.Println("Array:", arr)
		fmt.Println("Partition index:", index)
	case "qsort":
		fmt.Println("Quick sort algo:")
		arr := []int{12, 11, 10, 5, 4, 23, 34, 2}
		Array.Qsort(arr, 0, len(arr)-1)
		fmt.Println(arr)
	case "spk":
		fmt.Println("Sort package example:")
		Util.SortPkExample()
	case "rotat":
		fmt.Println("Rotate an array:")
		arr := []int{12, 13, 18, 3, 9, 43, 29, 4}
		Array.RotateArray(arr, 3)
		fmt.Println("The state of the array after rotation.")
		fmt.Println(arr)
	case "rots":
		arr := []int{12, 13, 17, 19, 21, 10, 11}
		fmt.Println("Find the index of the key in rotated sorted:")
		for i := 0; i < len(arr); i++ {
			index := Array.SearchRotated(arr, 0, len(arr)-1, arr[i])
			fmt.Println(index)
		}
	case "secl":
		arr := []int{12, 14, 18, 2, 21, 24, 18, 16, 22}
		fmt.Println("Finding the second largest element in array:")
		num := Array.FindSecondLargest(arr)
		fmt.Println(num)
		fmt.Println(arr)
	case "fndm":
		fmt.Println("The minimum element in the rotated array is:")
		brr := []int{4, 5, 6, 7, 1, 2, 3}
		num := Array.FindMinm(brr)
		fmt.Println(num)
	case "mtime":
		fmt.Println("Minimum time to produce m items:")
		machines := []int{2, 4, 5}
		minm_time := Util.MinTimeMItem(machines, 7, 3)
		fmt.Println(minm_time)
	case "mmtime":
		fmt.Println("Minimum time to produce m items:")
		machines := []int{2, 4, 5}
		minm_time := Util.MinmTime(machines, 7, 3)
		fmt.Println(minm_time)
	case "countP":
		fmt.Println("Counting the minimum pages that can be allocated among student.")
		arr := []int{12, 34, 67, 90}
		num := Util.CalculatePages(arr, 2)
		fmt.Println(num)
	case "peake":
		fmt.Println("Finding the paeak element in an array:")
		arr := []int{1, 2, 4, 5, 7, 8, 3}
		elm := Util.FindPeakElement(arr)
		fmt.Println(arr[elm])
	case "bsrch":
		fmt.Println("Find key using binary search.")
		arr := []int{12, 16, 18, 20, 23, 26, 56, 78}
		index := Array.BinarySearch(arr, 78)
		fmt.Println(index)
	case "bsrchr":
		fmt.Println("Binary search reverse:")
		arr := []int{24, 23, 16, 14, 11, 8, 6}
		for i := 0; i < len(arr); i++ {
			index := Array.BinarySearchReverse(arr, arr[i])
			fmt.Println(index)
		}
	case "ordr":
		fmt.Println("Order agnostic search.")
		arr := []int{11, 14, 16, 17, 26, 28, 30}
		for i := 0; i < len(arr); i++ {
			index := Array.BinarySearch(arr, arr[i])
			fmt.Println(index)
		}
		fmt.Println("Descending order:")
		brr := []int{24, 21, 19, 16, 15, 12, 11}
		for i := 0; i < len(brr); i++ {
			index := Array.BinarySearchReverse(brr, brr[i])
			fmt.Println(index)
		}
	case "frsto":
		fmt.Println("First occurance of a number:")
		crr := []int{2, 4, 6, 6, 7, 8, 9, 11, 13}
		index := Array.FirstOccurance(crr, 6)
		fmt.Println(index)
		fmt.Println("Searching a non repeated element.")
		brr := []int{2, 4, 7, 9, 11, 18, 19, 21, 24}
		index = Array.FirstOccurance(brr, 4)
		fmt.Println(index)
	case "lasto":
		fmt.Println("Last occurance of a number:")
		arr := []int{4, 5, 10, 10, 10, 12, 15, 17}
		index := Array.LastOccurance(arr, 10)
		fmt.Println(index)
	case "cnte":
		fmt.Println("Find the first and last occurance:")
		arr := []int{2, 3, 4, 5, 5, 5, 6, 7, 9}
		lngth := Array.CountNumElements(arr, 5)
		fmt.Println(lngth)
	case "finde":
		fmt.Println("Find minimum element in array which is rotated.")
		arr := []int{10, 12, 14, 16, 6, 8, 9}
		element := Util.FindMinmelement(arr)
		fmt.Println(element)
	case "bsts":
		stock_price := []int{8, 3, 3, 10, 2, 10}
		fmt.Println("Best time to buy sell stock:")
		profit := SlidingWindow.BestTimeSellBuy(stock_price)
		fmt.Println(profit)
	case "lngs":
		fmt.Println("The substring with longest unique char.")
		str := "pwwkew"
		temp_lngth := SlidingWindow.LongestRepeat(str)
		fmt.Println(temp_lngth)
	case "minw":
		fmt.Println("Minimum window substring:")
		str1 := "ADOBECODEBANC"
		str2 := "ABC"
		lngth := SlidingWindow.MinWindowSubstr(str1, str2)
		fmt.Println(lngth)
	case "maxs":
		fmt.Println("Find maximum simillar substring.")
		str := "AABABBA"
		lngth := SlidingWindow.FindMaxSubSimillar(str, 1)
		fmt.Println(lngth)
	case "isan":
		fmt.Println("Verify if the two strings are anagrams.")
		temp := "cbaebabacd"
		substr := "abc"
		index_list := SlidingWindow.FindAnagram(temp, substr)
		fmt.Println(index_list)
	case "mpu":
		fmt.Println("Map utilities.")
		Map.MapUtils()
	default:
		fmt.Println("Invalid operation.")
	}
}
