func hasDuplicate(nums []int) bool {
	numsTracker := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := numsTracker[num]; ok { return true }
		
		numsTracker[num] = struct{}{}
	}

	return false    
}
