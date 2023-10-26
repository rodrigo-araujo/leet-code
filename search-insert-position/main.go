package main

import "fmt"

func main() {

	var nums []int
	var target int
	var result int

	nums = []int{1, 3, 5, 6}
	target = 5
	result = searchInsert(nums, target)
	fmt.Println(result)

	nums = []int{1, 3, 5, 6}
	target = 2
	result = searchInsert(nums, target)
	fmt.Println(result)

	nums = []int{1, 3, 5, 6}
	target = 7
	result = searchInsert(nums, target)
	fmt.Println(result)

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	target = 4
	result = searchInsert(nums, target)
	fmt.Println(result)

}

func searchInsert(nums []int, target int) int {
	start := 0
	final := len(nums) - 1
	for start <= final {
		middle := (start + final) / 2

		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			final = middle - 1
		} else { // target > nums[meio]
			start = middle + 1
		}
	}

	return start
}
