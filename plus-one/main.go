package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var digits []int
	var result []int

	digits = []int{1, 2, 3}
	result = plusOne(digits)
	fmt.Println(result)

	digits = []int{4, 3, 2, 1}
	result = plusOne(digits)
	fmt.Println(result)

	digits = []int{9}
	result = plusOne(digits)
	fmt.Println(result)

	digits = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	result = plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	last := len(digits) - 1
	carry := 1
	for x := last; x >= 0; x-- {
		tmp := strconv.Itoa(digits[x] + carry)

		parts := strings.Split(tmp, "")
		if len(parts) == 1 {
			value, _ := strconv.Atoi(parts[0])
			result[x+1] = value
			carry = 0
		} else {
			value, _ := strconv.Atoi(parts[1])
			result[x+1] = value
			carry = 1
		}
	}

	if carry == 1 {
		result[0] = 1
	} else {
		result = result[1:]
	}

	return result
}
