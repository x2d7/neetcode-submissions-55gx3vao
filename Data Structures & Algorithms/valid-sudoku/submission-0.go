type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(e T) bool{
	_, ok := s.m[e]
	if !ok {
		s.m[e] = struct{}{}
	}
	return !ok
}

type state struct {
	rows     [9]Set[byte]
	columns  [9]Set[byte]
	sections [9]Set[byte]
}

func NewState() state {
	var rows, columns, sections [9]Set[byte]
	for i, _ := range rows {
		rows[i] = NewSet[byte]()
	}
	for i, _ := range columns {
		columns[i] = NewSet[byte]()
	}
	for i, _ := range sections {
		sections[i] = NewSet[byte]()
	}

	return state{rows, columns, sections}
}

func isValidSudoku(board [][]byte) bool {
	state := NewState()
	emptyByte := byte('.')

	for rowIdx, row := range board {
		for colIdx, e := range row {
			if e == emptyByte {
				continue
			}
			
			secIdx := (rowIdx / 3) * 3 + (colIdx / 3)

			if ok := state.rows[rowIdx].Add(e); !ok {
				return false
			}

			if ok := state.columns[colIdx].Add(e); !ok {
				return false
			}

			if ok := state.sections[secIdx].Add(e); !ok {
				return false
			}
		}
	}
	return true
}
