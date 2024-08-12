package sug

func (s *Sudoku) MakeNewPuzzle() {
	s.Puzzle = nil
	s.Answers = nil
	s.MakePuzzle()
}

func (s *Sudoku) MakePuzzle() {
	p := make([][]int, Row)
	for i := range Row {
		p[i] = make([]int, Col)
	}
	for i := range Row {
		for j := range Col {
			p[i][j] = (j+1+i)%9 + 1
		}
	}
	Shuffle(p)
	s.Answers = append(s.Answers, p)
	cp := CopySudoku(p)
	RemoveCells(cp)
	s.Puzzle = cp
}
