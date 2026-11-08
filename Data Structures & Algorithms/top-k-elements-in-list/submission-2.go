func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 { return []int{nums[0]} }

	count := make(map[int]int)
	result := make(map[int]struct{})

	for _, num := range nums {
		count[num]++
		
		minKey := 0
		minVal := len(nums) + 1

		for key := range result {
			if count[key] < minVal {
				minKey = key
				minVal = count[key] 
			}
		}

		if _, ok := result[num]; !ok {
			if len(result) < k {
				result[num] = struct{}{}
			} else if count[num] > count[minKey] {
				result[num] = struct{}{}
				delete(result, minKey)
			}	
		}

	}

	finalResult := []int{}

	for key := range result {
		finalResult = append(finalResult, key)
	}

	return finalResult
}
