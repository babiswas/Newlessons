package arr

import (
	"fmt"
)

var lst = make([]string, 0)

func PrintallSubset(ip string, op string) {
	if len(ip) == 0 {
		lst = append(lst, op)
		return
	}
	op1 := op
	temp := string(ip[0])
	op2 := op + temp
	ip = ip[1:]
	PrintallSubset(ip, op1)
	PrintallSubset(ip, op2)
}

func PrintAll() {
	fmt.Println("Displaying all the subsets.")
	PrintallSubset("abcd", "")
	fmt.Println(lst)
}
