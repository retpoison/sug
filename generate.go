package sug

func (s *Sudoku) MakeNewPuzzle(difficulty int) {
	s.Puzzle = nil
	s.Answers = nil
	s.MakePuzzle(difficulty)
}

func (s *Sudoku) MakePuzzle(difficulty int) {
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
	RemoveCells(cp, difficulty)
	s.Puzzle = cp
}
