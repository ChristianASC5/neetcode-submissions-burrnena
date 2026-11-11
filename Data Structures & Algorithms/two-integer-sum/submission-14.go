func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i, num := range nums {
		compliment := target - num

		if _, ok := seen[compliment]; ok {
			return []int{seen[compliment], i}
		}

		seen[num] = i
	}

	return []int{}
}
