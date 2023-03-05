package main

import (
	"fmt"
	"os"
	"os/exec"
)

var board = [3][3]string{}

var plO = "o" // 1
var plX = "x" // 2

var player1 string
var player2 string

var currentPlayer = ""

func main() {

	for {
		var playerSide string

		var choose string

		fmt.Println("You're first player, choose you side:")

		fmt.Printf("1. - %s\n", plX)
		fmt.Printf("2. - %s\n", plO)

		fmt.Scan(&choose)

		if choose == plX {
			player1 = plX
			player2 = plO
		} else {
			player1 = plO
			player2 = plX
		}
		playerSide = player2

		currentPlayer = plX

		fmt.Printf("You're second player, you're side is: %s\n", playerSide)

		// Game start
	mg:
		for {
			// Making move

			PrintBoard()

			var row, col int

			fmt.Println("Enter column:")

			fmt.Scan(&col)

			fmt.Println("Enter row:")

			fmt.Scan(&row)

			if row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != "" {
				fmt.Println("Invalid input, please, try again")
				continue
			}

			MakeMove(row, col, currentPlayer)

			// Победа?
			if WinChecker() {
				PrintBoard()
				if currentPlayer == player1 {
					fmt.Println("Player №1 WIN!!!")
				} else {
					fmt.Println("Player №2 WIN!!!")
				}
				break mg
			}

			// Ничья?
			if DrawChecker() {
				PrintBoard()
				fmt.Println("Draw :((")
				break mg
			}

			SwitchPlayers()
			ClearScreen()
		}
		board = [3][3]string{}
	}
}

func PrintBoard() {
	fmt.Println("  1 2 3")

	for x, row := range board {
		fmt.Print(x + 1)
		for _, col := range row {
			if col == "" {
				fmt.Print(" -")
			} else {
				fmt.Print(" " + col)
			}
		}
		fmt.Println()
	}
}

func MakeMove(row, col int, pl string) {
	// Игрок выбирает сначала индекс ряда, а потом номер индекс в этом ряду
	// В конце вызывать проверку на победу

	if pl == "o" {
		board[row-1][col-1] = plO
	} else {
		board[row-1][col-1] = plX
	}
}

func SwitchPlayers() {
	if currentPlayer == plX {
		currentPlayer = plO
	} else {
		currentPlayer = plX
	}
}

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func WinChecker() bool {

	for i := 0; i < 3; i++ {
		// По вертикали
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		// По горизонтали
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && currentPlayer == board[2][i] {
			return true
		}
	}

	// По диагонали обычной
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	// По диагонали обратной
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

func DrawChecker() bool {
	for _, row := range board {
		for _, col := range row {
			if col == "" {
				return false
			}
		}
	}
	return true
}
