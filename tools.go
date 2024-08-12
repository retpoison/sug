package sug

import (
	"fmt"
	"math/rand/v2"
)

func CopySudoku(s [][]int) [][]int {
	var cp [][]int = make([][]int, len(s))
	for i := range len(s) {
		cp[i] = make([]int, len(s[i]))
		copy(cp[i], s[i])
	}
	return cp
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GetValidNumbers(s [][]int, i, j int) []int {
	validNumbers := []int{}
	if s[i][j] != 0 {
		return validNumbers
	}
	inNumbers := GetNumbersInCol(s, j)
	inNumbers = append(inNumbers, GetNumbersInRow(s, i)...)
	inNumbers = append(inNumbers, GetNumbersInSubSquare(s, i, j)...)
	for i := 1; i < 10; i++ {
		for _, v := range inNumbers {
			if i == v {
				goto next
			}
		}
		validNumbers = append(validNumbers, i)
	next:
	}
	return validNumbers
}

func GetNumbersInCol(s [][]int, col int) []int {
	inNumbers := []int{}
	for i := range Row {
		if s[i][col] != 0 {
			inNumbers = append(inNumbers, s[i][col])
		}
	}
	return inNumbers
}

func GetNumbersInRow(s [][]int, row int) []int {
	inNumbers := []int{}
	for j := range Col {
		if s[row][j] != 0 {
			inNumbers = append(inNumbers, s[row][j])
		}
	}
	return inNumbers
}

func GetNumbersInSubSquare(s [][]int, i, j int) []int {
	centers := []int{1, 4, 7}
	cell := []int{-1, 0, 1}
	inNumbers := []int{}
	var sqi, sqj int
	for _, v := range centers {
		if Abs(v-i) < 2 {
			sqi = v
		}
		if Abs(v-j) < 2 {
			sqj = v
		}
	}
	for _, sx := range cell {
		for _, sy := range cell {
			v := s[sqi+sx][sqj+sy]
			if v != 0 {
				inNumbers = append(inNumbers, v)
			}
		}
	}
	return inNumbers
}

func GetFirstZero(s [][]int) (int, int) {
	for i := range Row {
		for j := range Col {
			if s[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func PrintSudoku(s [][]int) {
	for i := range Row {
		for j := range Col {
			if j%3 == 0 && j != 0 {
				fmt.Print("|")
			}
			fmt.Printf("%d ", s[i][j])
		}
		fmt.Print("\n")
		if i%3 == 2 && i != 8 {
			fmt.Print("-------------------\n")
		}
	}
}

func Shuffle(s [][]int) {
	n := 100
start:
	for _ = range rand.IntN(n) {
		SwapRow(s)
		SwapCol(s)
	}
	if !IsValidSudoku(s) {
		n = 10
		goto start
	}
}

func SwapRow(s [][]int) {
randr:
	r1 := rand.IntN(Row)
	r2 := rand.IntN(Row)
	if r1 == r2 {
		goto randr
	}
	t := s[r1]
	s[r1] = s[r2]
	s[r2] = t
}

func SwapCol(s [][]int) {
randc:
	c1 := rand.IntN(Col)
	c2 := rand.IntN(Col)
	if c1 == c2 {
		goto randc
	}
	for i := range Row {
		t := s[i][c1]
		s[i][c1] = s[i][c2]
		s[i][c2] = t
	}
}

func IsValidSudoku(s [][]int) bool {
	if !IsSquaresValid(s) || !IsRowsColsValid(s) {
		return false
	}
	return true
}

func IsSquaresValid(s [][]int) bool {
	centers := []int{1, 4, 7}
	for _, i := range centers {
		for _, j := range centers {
			if !IsNumbersValid(GetNumbersInSubSquare(s, i, j)) {
				return false
			}
		}
	}

	return true
}

func IsRowsColsValid(s [][]int) bool {
	for i := range Row {
		if !IsNumbersValid(GetNumbersInRow(s, i)) ||
			!IsNumbersValid(GetNumbersInCol(s, i)) {
			return false
		}
	}
	return true
}

func IsNumbersValid(numbers []int) bool {
	cnumbers := [9]int{0}
	for _, v := range numbers {
		cnumbers[v-1]++
	}
	for _, v := range cnumbers {
		if v > 1 {
			return false
		}
	}
	return true
}

func RemoveCells(puzzle [][]int) {
	var r, c, l int
	for {
		ans, err := Solve(puzzle)
		if err != nil {
			puzzle[r][c] = l
		} else if len(ans) > 1 {
			puzzle[r][c] = l
			break
		}
	start:
		r = rand.IntN(Row)
		c = rand.IntN(Col)
		if puzzle[r][c] == 0 {
			goto start
		}
		l = puzzle[r][c]
		puzzle[r][c] = 0
	}
}
