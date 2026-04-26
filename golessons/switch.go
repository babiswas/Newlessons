package arr

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CollectNumber() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error occured:")
		panic("illigeal input provided.")
	}
	return num
}

func MathOperations(ops string) {
	switch ops {
	case "add":
		fmt.Println("This is add operation")
		fmt.Println("Enter first number:")
		num1 := CollectNumber()
		fmt.Println("Enter the second number:")
		num2 := CollectNumber()
		fmt.Println("The sum of the numbers is:", num1+num2)
	case "subs":
		fmt.Println("This is substarct operation:")
		fmt.Println("Enter the first number:")
		num1 := CollectNumber()
		fmt.Println("Enter the second number:")
		num2 := CollectNumber()
		fmt.Println("The difference of the number is:")
		if num1 >= num2 {
			fmt.Println(num1 - num2)
		} else {
			fmt.Println(num2 - num1)
		}
	case "multiply":
		fmt.Println("This is multiplication operation:")
		fmt.Println("Enter the first number:")
		num1 := CollectNumber()
		fmt.Println("Enter the second number:")
		num2 := CollectNumber()
		fmt.Println("The product of the numbers:")
		fmt.Println(num1 * num2)
	case "divide":
		fmt.Println("This is dividing the number:")
		fmt.Println("Enter the first number:")
		num1 := CollectNumber()
		fmt.Println("Enter the second number:")
		num2 := CollectNumber()
		fmt.Println("The divison of the numbers is:")
		if num1 >= num2 && num2 != 0 {
			fmt.Println(num1 / num2)
		} else if num1 != 0 && num2 > num1 {
			fmt.Println(num2 / num1)
		}
	default:
		fmt.Println("Invalid operation:")
	}
}

func MathOps() {
	fmt.Println("Enter operation:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ops := scanner.Text()
	MathOperations(ops)
}
