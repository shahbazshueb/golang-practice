package main

import "fmt"

func main() {
	// create a tic-tac-toe board with a winning move
	board := [][]string{}
	win(&board)
	for i := 0; i < 3; i++ {
		fmt.Println(board[i])
	}
}

func win(board *[][]string) {
	row1 := []string{" ", "o", "x"}
	row2 := []string{"o", "x", " "}
	row3 := []string{"x", " ", "o"}
	*board = append(*board, row1)
	*board = append(*board, row2)
	*board = append(*board, row3)

}
