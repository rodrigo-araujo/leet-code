package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)

	nums = []int{3, 2, 4}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result)

	nums = []int{3, 3}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	missing := make(map[int]int, len(nums))

	for currentIndex, value := range nums {
		m := target - value
		missingKey, found := missing[value]
		if found {
			return []int{currentIndex, missingKey}
		}
		missing[m] = currentIndex

	}

	return nil
}
