package algo

import "math"

func Miops(str string) int {
	newstr := str + str
	alt1, alt2 := "", ""
	min_lngth := math.MinInt
	for i := 0; i < len(newstr); i++ {
		if i%2 == 0 {
			alt1 += "1"
			alt2 += "0"
		} else {
			alt1 += "0"
			alt2 += "1"
		}
	}
	left, right := 0, 0
	n1, n2 := 0, 0
	for right < len(newstr) {
		if newstr[right] != alt1[right] {
			n1 += 1
		}
		if newstr[right] != alt2[right] {
			n2 += 1
		}
		if (right - left + 1) > len(str) {
			if newstr[left] != alt1[left] {
				n1 -= 1
			}
			if newstr[left] != alt2[left] {
				n2 -= 1
			}
		}
		if (right - left + 1) == len(str) {
			if n1 > n2 {
				min_lngth = n2
			} else {
				min_lngth = n1
			}
		}
		right += 1
	}
	return min_lngth
}
