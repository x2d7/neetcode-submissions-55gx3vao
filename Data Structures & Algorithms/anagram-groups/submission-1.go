func getStringProfile(str string) [26]int {
	var result [26]int
	for _, c := range str {
		result[c-'a'] += 1
	}

	return result
}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	m := make(map[[26]int][]string, len(strs))

	for _, str := range strs {
		s := getStringProfile(str)
		m[s] = append(m[s], str)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
