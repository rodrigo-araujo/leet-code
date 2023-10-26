package main

import "fmt"

func main() {
	var result int
	result = climbStairs(1)
	fmt.Println(result)

	result = climbStairs(2)
	fmt.Println(result)

	result = climbStairs(3)
	fmt.Println(result)

	result = climbStairs(4)
	fmt.Println(result)

	result = climbStairs(5)
	fmt.Println(result)

	result = climbStairs(6)
	fmt.Println(result)

	result = climbStairs(7)
	fmt.Println(result)

	result = climbStairs(8)
	fmt.Println(result)
}

func climbStairs(n int) int {
	previous := 1
	last := 1

	for i := 0; i < n-1; i++ {
		//temp := previous
		previous = previous + last
		last = previous - last
	}

	return previous
}
