func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, val := range nums {
		numsMap[val] = i // {1: 0, 3: 1, 4: 2, 2: 3}
	}

	for i, val := range nums {
		compliment := target - val // 5

		j, complimentExists := numsMap[compliment]

		if complimentExists && i != j {
			return []int{i, j}
		}
	}
	
	return []int{}
}
