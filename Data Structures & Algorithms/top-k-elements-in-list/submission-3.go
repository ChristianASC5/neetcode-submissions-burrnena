func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, val := range nums {
		counts[val]++
	}

	buckets := make([][]int, len(nums) + 1)
	
	for num, count := range counts {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, n := range buckets[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
