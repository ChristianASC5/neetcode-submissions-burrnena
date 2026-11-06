func hasDuplicate(nums []int) bool {
	numsMap := make(map[int]int)
 
    for _, num := range nums {
		if _, numExists := numsMap[num]; numExists {
			return true
		}

		numsMap[num] = 1
	}

	return false
}
