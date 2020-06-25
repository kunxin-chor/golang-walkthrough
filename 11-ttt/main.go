package main

import "fmt"

func main() {
	var board [3][3]int
	turn := 0
	stop := false
	// while the game has not ended
	for stop == false {
		fmt.Println("Turn", turn)
		displayBoard(board[:])
		fmt.Printf("Player %d, please enter move, or -1 -1 to quit: ", getPlayer(turn))

		row, col := 0, 0
		fmt.Scanf("%d %d", &row, &col)
		if row == -1 || col == -1 {
			stop = true
		}
		if (board[row][col] == 0) {
			board[row][col] = getPlayer(turn)
			// increase turn
			turn++;

		} else {
			fmt.Println("Invalid move, please try again")
		}

		// has anyone won? (-1 is none, 0 or 1)
		winner := gameEnded(board[:])
		if winner != -1 {
			stop = true
			fmt.Printf("Player %d has won!", winner+1)
		}

	}
}

func getPlayer(turn int) int {
	if turn % 2 == 0 {
		return 1
	} else {
		return 2
	}
}

func displayBoard(board [][3]int) {
	for _, r := range board {
		for _, c := range r {
			fmt.Printf("%d ", c)
		}
		fmt.Println()
	}
}

func gameEnded(board [][3]int) int {
	// check if who have won

	// check rows
	for i :=0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
	}

	// check cols
	for j :=0; j < 3; j++ {
		if board[0][j] == board[1][j] && board[0][j] == board[2][j] {
			return board[0][j]
		}
	}

	// check diagionals
	if board[0][0] == board[1][1] && board[0][0 ]== board[2][2] {
		return board[0][0]
	}

	if board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}

	// no result
	return -1
}
