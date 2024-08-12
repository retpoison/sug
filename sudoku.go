package sug

import "fmt"

type Sudoku struct {
	Puzzle  [][]int
	Answers [][][]int
}

func NewSudoku() *Sudoku {
	s := new(Sudoku)
	s.MakeNewPuzzle()
	return s
}

func (s *Sudoku) PrintPuzzle() {
	fmt.Println("Puzzle:")
	PrintSudoku(s.Puzzle)
}

func (s *Sudoku) PrintAnswers() {
	for i, v := range s.Answers {
		fmt.Println("Answer", i+1, ":")
		PrintSudoku(v)
	}
}
