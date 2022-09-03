package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintBoard(board *[3][3]string) {
	b := *board
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
}

func IsBoardFull(board *[3][3]string) bool {
	for _, r := range *board {
		for _, c := range r {
			if c == "_" {
				return false
			}
		}
	}
	return true
}

func IsWin(board *[3][3]string, user string) bool {
	b := *board

	row1 := b[0][0] == b[0][1] && b[0][0] == b[0][2] && b[0][0] == user
	row2 := b[1][0] == b[1][1] && b[1][0] == b[1][2] && b[1][0] == user
	row3 := b[2][0] == b[2][1] && b[2][0] == b[2][2] && b[2][0] == user

	col1 := b[0][0] == b[1][0] && b[0][0] == b[2][0] && b[0][0] == user
	col2 := b[0][1] == b[1][1] && b[0][1] == b[2][1] && b[0][1] == user
	col3 := b[0][2] == b[1][2] && b[0][2] == b[2][2] && b[0][2] == user

	l_angle := b[0][0] == b[1][1] && b[0][0] == b[2][2] && b[0][0] == user
	r_angle := b[0][2] == b[1][1] && b[0][2] == b[2][0] && b[2][0] == user

	return row1 || row2 || row3 || col1 || col2 || col3 || l_angle || r_angle
}

func SwitchCurrentUser(currentPlayer *string) {
	switch *currentPlayer {
	case "x":
		*currentPlayer = "o"
	case "o":
		*currentPlayer = "x"
	}
}

func WriteBoard(board *[3][3]string, row int, col int, user string) {
	b := *board
	b[row][col] = user
	*board = b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentPlayer := "x"
	board := [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	PrintBoard(&board)

	for {

		fmt.Printf("(%v) Player (row/col): ", currentPlayer)
		text, _ := reader.ReadString('\n')
		move := strings.Split(strings.TrimSpace(text), "/")

		if len(move) != 2 {
			fmt.Printf("Please Enter row / col. e.g 1/1, 2/2 \n")
			continue
		}
		row, _ := strconv.Atoi(strings.TrimSpace(move[0]))
		col, _ := strconv.Atoi(strings.TrimSpace(move[1]))
		if 3 < row || row < 1 || 3 < col || col < 1 {
			fmt.Printf("Please Enter row / col between 1-3. e.g 1/1, 2/3 \n")
			continue
		}
		row -= 1
		col -= 1

		if board[row][col] != "_" {
			fmt.Println("Try Another This field is fill-up.")
			continue
		}
		WriteBoard(&board, row, col, currentPlayer)
		PrintBoard(&board)

		if IsWin(&board, "x") {
			fmt.Println("Congratulation X Win.")
			break
		}

		if IsWin(&board, "o") {
			fmt.Println("Congratulation O Win.")
			break
		}

		if IsBoardFull(&board) {
			fmt.Println("Tie No One Win.")
			break
		}

		SwitchCurrentUser(&currentPlayer)
	}
}
