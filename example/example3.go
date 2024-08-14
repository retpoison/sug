package main

import (
	"fmt"
	"log"
	"time"

	"github.com/retpoison/sug"
)

func main() {
	l := "52...6.........7.13...........4..8..6......5...........418.........3..2...87....."
	var s = make([][]int, sug.Row)
	for i := range sug.Row {
		s[i] = make([]int, sug.Col)
	}

	var num int
	for i := range sug.Row * sug.Col {
		ch := l[i]
		if ch == '.' {
			num = 0
		} else {
			num = int(ch) - '0'
		}
		s[i/sug.Row][i%sug.Col] = num
	}

	start := time.Now()
	ans, err := sug.Solve(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle:")
	sug.PrintSudoku(s)
	for i := range len(ans) {
		fmt.Println("Answer", i+1, ":")
		sug.PrintSudoku(ans[i])
	}
	fmt.Println("Solved in: ", time.Since(start))
}
