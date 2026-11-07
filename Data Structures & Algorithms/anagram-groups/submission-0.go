func runeCounts(str string) [26]int {
	var counts [26]int

	for _, v := range str { counts[v - 'a']++ }

	return counts
}

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	
	for _, str := range strs {
		runeCount := runeCounts(str)

		if _, ok := anagramMap[runeCount]; !ok {
			anagramMap[runeCount] = []string{str}
		} else {
			anagramMap[runeCount] = append(anagramMap[runeCount], str)
		}
	}
	
	var anagramArray [][]string

	for _, value := range anagramMap {
		anagramArray = append(anagramArray, value)
	}

	return anagramArray
}
