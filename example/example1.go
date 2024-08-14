package main

import (
	"github.com/retpoison/sug"
)

func main() {
	s := sug.NewSudoku(sug.Easy)
	s.PrintPuzzle()
	s.PrintAnswers()
}
