func runeCounts(str string) [26]int {
	var counts [26]int

	for _, v := range str { counts[v - 'a']++ }

	return counts
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	
	for _, str := range strs {
		runeCount := runeCounts(str)

		groups[runeCount] = append(groups[runeCount], str)
	}
	
	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
