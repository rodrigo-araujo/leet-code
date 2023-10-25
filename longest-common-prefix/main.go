package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println("result: ", result)

	strs = []string{"dog", "racecar", "car"}
	result = longestCommonPrefix(strs)
	fmt.Println("result: ", result)

}

func longestCommonPrefix(strs []string) string {
	prefix := ""

	for key, word := range strs {
		if key == 0 {
			prefix = word
			continue
		}

		size := len(prefix)
		if len(word) < size {
			size = len(word)
			prefix = prefix[0:size]
		}

		for x := 0; x < size; x++ {
			if word[x] != prefix[x] {
				prefix = prefix[0:x]
				break
			}
		}

		if prefix == "" {
			break
		}
	}
	return prefix
}
