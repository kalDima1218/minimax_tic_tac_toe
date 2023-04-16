package main

import "fmt"

// Define a function to evaluate the score of the board
func evaluate(board [3][3]string) int {
	// Check rows for a win
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == "X" {
				return 10
			} else if board[i][0] == "O" {
				return -10
			}
		}
	}

	// Check columns for a win
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			if board[0][i] == "X" {
				return 10
			} else if board[0][i] == "O" {
				return -10
			}
		}
	}

	// Check diagonals for a win
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == "X" {
			return 10
		} else if board[0][0] == "O" {
			return -10
		}
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == "X" {
			return 10
		} else if board[0][2] == "O" {
			return -10
		}
	}

	// If no one has won yet, return 0
	return 0
}

// Define the minimax function
func minimax(board [3][3]string, depth int, isMax bool) (int, [2]int) {
	// Evaluate the score of the board
	score := evaluate(board)

	// If the maximizer has won, return their score
	if score == 10 {
		return score, [2]int{}
	}

	// If the minimizer has won, return their score
	if score == -10 {
		return score, [2]int{}
	}

	// If there are no more moves and no one has won, return 0
	if !isMovesLeft(board) {
		return 0, [2]int{}
	}

	// If it's the maximizer's turn
	if isMax {
		// Initialize the best score and best move
		bestScore := -1000
		var bestMove [2]int

		// Loop through all the cells
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// If the cell is empty
				if board[i][j] == " " {
					// Make the move
					board[i][j] = "X"

					// Recursively call minimax and get the score and move
					score, _ := minimax(board, depth+1, !isMax)

					// Undo the move
					board[i][j] = " "

					// Update the best score and best move
					if score > bestScore {
						bestScore = score
						bestMove = [2]int{i, j}
					}
				}
			}
		}

		// Return the best score and best move
		return bestScore, bestMove
	} else { // If it's the minimizer's turn
		// Initialize the best score and best move
		bestScore := 1000
		var bestMove [2]int

		// Loop through all the cells
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// If the cell is empty
				if board[i][j] == " " {
					// Make the move
					board[i][j] = "O"

					// Recursively call minimax and get the score and move
					score, _ := minimax(board, depth+1, !isMax)

					// Undo the move
					board[i][j] = ""

					// Update the best score and best move
					if score < bestScore {
						bestScore = score
						bestMove = [2]int{i, j}
					}
				}
			}
		}

		// Return the best score and best move
		return bestScore, bestMove
	}
}

// Define a function to check if there are any moves left
func isMovesLeft(board [3][3]string) bool {
	// Loop through all the cells
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// If the cell is empty
			if board[i][j] == " " {
				// There is at least one move left
				return true
			}
		}
	}

	// There are no moves left
	return false
}

func main() {
	board := [3][3]string{
		{" ", "O", " "},
		{" ", "X", " "},
		{" ", " ", " "},
	}

	score, move := minimax(board, 2, true)

	fmt.Println(score, move)
}
