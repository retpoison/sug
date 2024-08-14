package main

import (
	"fmt"
	"log"
	"time"

	"github.com/retpoison/sug"
)

func main() {
	s := [][]int{
		[]int{5, 0, 3, 0, 6, 0, 0, 0, 0},
		[]int{2, 0, 0, 0, 0, 0, 4, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 1, 0, 8},
		[]int{0, 0, 0, 4, 0, 0, 8, 0, 7},
		[]int{0, 0, 0, 0, 0, 0, 0, 3, 0},
		[]int{6, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 7, 0, 8, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 5, 0, 0, 2, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0},
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
