func countLetters(s string) map[rune]int {
	result := make(map[rune]int)
	for _, r := range s {
		result[r] += 1
	}

	return result
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	r1, r2 := countLetters(s), countLetters(t)
	for k, v1 := range r1 {
		v2 := r2[k]
		if v1 != v2 {
			return false
		}
	}
	return true
}
