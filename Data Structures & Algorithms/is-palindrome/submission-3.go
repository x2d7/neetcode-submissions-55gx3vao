func ReverseString(s string) string {
    runes := []rune(s)
    var sb strings.Builder
    sb.Grow(len(runes))
    for i := len(runes) - 1; i >= 0; i-- {
        sb.WriteRune(runes[i])
    }
    return sb.String()
}

func isPalindrome(s string) bool {
	var sb strings.Builder

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = 'a' + r - 'A'
		}
		
		if !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')) {
			continue
		} 

		sb.WriteRune(r)
	}


	s = sb.String()

	return ReverseString(s) == s
}
