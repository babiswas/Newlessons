package Map

import (
	"fmt"
	"maps"
	"strings"
)

func MapUtils() {
	mp := make(map[string]int)
	mp["hello"] = 1
	mp["bello"] = 2
	mp["tello"] = 3

	for key, val := range mp {
		fmt.Println(key, "=", val)
	}
	fmt.Println("The length of the map is:", len(mp))
	fmt.Println(mp)

	tp := make(map[string]int)
	tp["hello"] = 1
	tp["bello"] = 2
	tp["tello"] = 3

	fmt.Println(maps.Equal(mp, tp))

	clear(mp)
	fmt.Println(mp)

	fmt.Println(maps.Equal(mp, tp))
	np := make(map[string]int)
	str := "The quick brown fox jumped over lazy dog."
	for index, val := range strings.Split(str, " ") {
		fmt.Println(index)
		_, ok := np[val]
		if ok {
			np[val] += 1
		} else {
			np[val] = 1
		}
	}
	fmt.Println(np)
	delete(np, "fox")
	fmt.Println(np)
}
