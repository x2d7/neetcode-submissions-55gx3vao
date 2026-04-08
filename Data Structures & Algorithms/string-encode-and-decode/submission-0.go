import "io"

type Solution struct{}

func readNRunes(r *strings.Reader, n int) string {
    var sb strings.Builder
    for i := 0; i < n; i++ {
        ch, _, err := r.ReadRune()
        if err == io.EOF {
            break
        }
        sb.WriteRune(ch)
    }
    return sb.String()
}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		l := rune(len(str))
		sb.WriteRune(l)
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0)
	
	sr := strings.NewReader(encoded)
	for {
		l, _, err := sr.ReadRune()
		if err != nil {
			break
		}
		str := readNRunes(sr, int(l))
		result = append(result, str)
	}

	return result
}
