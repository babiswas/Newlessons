package tree

import "fmt"

type Rob struct {
	val   int
	left  *Rob
	right *Rob
}

func (r *Rob) Insert(val int) {
	if r.val > val {
		if r.left != nil {
			r.left.Insert(val)
		} else {
			r.left = &Rob{val, nil, nil}
		}
	} else {
		if r.right != nil {
			r.right.Insert(val)
		} else {
			r.right = &Rob{val, nil, nil}
		}
	}
}

func Maxy(arr []int) int {
	maxm := 0
	for _, num := range arr {
		if maxm < num {
			maxm = num
		}
	}
	return maxm
}

func (r *Rob) MaxProfit() (int, int) {
	if r == nil {
		return 0, 0
	}
	leftSum, rightSum := []int{0, 0}, []int{0, 0}
	if r.left != nil {
		withRoot, withoutRoot := r.left.MaxProfit()
		leftSum[0] = withRoot
		leftSum[1] = withoutRoot
	}
	if r.right != nil {
		withRoot, withoutRoot := r.right.MaxProfit()
		rightSum[0] = withRoot
		rightSum[1] = withoutRoot
	}
	result1, result2 := (r.val + leftSum[1] + rightSum[1]), (Maxy(leftSum) + Maxy(rightSum))
	return result1, result2
}

func (r *Rob) Profit() int {
	res1, res2 := r.MaxProfit()
	if res1 > res2 {
		return res1
	} else {
		return res2
	}
}

func RobHouse() {
	node := Rob{15, nil, nil}
	node.Insert(22)
	node.Insert(20)
	node.Insert(26)
	node.Insert(8)
	node.Insert(10)
	node.Insert(6)
	fmt.Println("The maximum profit is:")
	profit := node.Profit()
	fmt.Println(profit)
}
