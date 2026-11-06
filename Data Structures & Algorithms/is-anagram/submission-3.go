func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }

	counts := make(map[rune]int)

	for _, val := range s { counts[val]++ }
	for _, val := range t { counts[val]-- }

	for _, val := range counts{
		if val != 0 { return false }
	}

	return true

}
