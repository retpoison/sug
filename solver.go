package sug

import (
	"errors"
)

var (
	Ans [][][]int
)

func Solve(puzzle [][]int) ([][][]int, error) {
	Ans = nil
	Ans = make([][][]int, 0)
	if !IsValidSudoku(puzzle) {
		return Ans, errors.New("Puzzle is not valid.")
	}
	su := CopySudoku(puzzle)
	FillDefinitBlocks(su)
	if i, j := GetFirstZero(su); i == -1 && j == -1 {
		Ans = append(Ans, su)
		return Ans, nil
	}
	TryAllPossibleValues(puzzle)
	return Ans, nil
}

func TryAllPossibleValues(puzzle [][]int) bool {
	zi, zj := GetFirstZero(puzzle)
	if zi == -1 && zj == -1 {
		return true
	}
	b := append(GetValidNumbers(puzzle, zi, zj), 10)
	for _, v := range b {
		if v == 10 {
			puzzle[zi][zj] = 0
			return false
		}
		puzzle[zi][zj] = v

		status := TryAllPossibleValues(puzzle)
		if status {
			nzi, nzj := GetFirstZero(puzzle)
			if nzi == -1 && nzj == -1 {
				cp := CopySudoku(puzzle)
				Ans = append(Ans, cp)
			}
		}
	}
	puzzle[zi][zj] = 0
	return true
}

func FillDefinitBlocks(s [][]int) {
fill:
	oneFlag := false
	for i := range Row {
		for j := range Col {
			b := GetValidNumbers(s, i, j)
			if s[i][j] == 0 && len(b) == 1 {
				s[i][j] = b[0]
				oneFlag = true
			}
		}
	}
	if oneFlag {
		goto fill
	}
}
