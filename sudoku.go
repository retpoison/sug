package sug

import "fmt"

const (
	Row = 9
	Col = 9

	/* from https://github.com/fedeztk/sku/blob/master/pkg/sudoku/sudoku.go */
	Easy   = 42
	Medium = 36
	Hard   = 27
	Expert = 25
)

type Sudoku struct {
	Puzzle  [][]int
	Answers [][][]int
}

func NewSudoku(difficulty int) *Sudoku {
	s := new(Sudoku)
	s.MakeNewPuzzle(difficulty)
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
