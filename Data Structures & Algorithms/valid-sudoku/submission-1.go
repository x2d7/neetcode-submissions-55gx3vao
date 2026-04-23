type Bitmask uint16

func (m *Bitmask) Add(e byte) bool {
	selected := Bitmask(1 << int(e - '1'))

	ok := *m & selected > 0
	if !ok {
		*m |= selected
	}
	return !ok
}

type state struct {
	rows     [9]Bitmask
	columns  [9]Bitmask
	sections [9]Bitmask
}

func NewState() state {
	var rows, columns, sections [9]Bitmask
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
