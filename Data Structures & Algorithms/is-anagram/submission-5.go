func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }

	counts := make(map[byte]int)

	for i := range len(s) {
		counts[s[i]]++
		counts[t[i]]--
	}

	for _, val := range counts {
		if val != 0 { return false }
	}

	return true
}
