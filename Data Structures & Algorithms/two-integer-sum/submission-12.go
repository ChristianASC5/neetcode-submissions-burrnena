func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for j, val := range nums {
		compliment := target - val

		i, complimentSeen := numsMap[compliment]
		
		if complimentSeen {
			return []int{i, j}
		}

		numsMap[val] = j
	}
	
	return []int{}
}
