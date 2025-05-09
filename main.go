package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"valedv.com/battleship/board"
	"valedv.com/battleship/player"
	"valedv.com/battleship/ship"
	"valedv.com/battleship/view"
)

var scanner = bufio.NewScanner(os.Stdin)
var size, difficulty int
var currPlayer player.Player
var currBoard board.Board

func main() {
	Greeting()
	for {
		fmt.Print("Options:\n1. Start game\n0. Quit\n>>> ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "0":
			os.Exit(0)
		case "1":
			InitializeGame()
		default:
			fmt.Println("Not a valid option, try again.")
		}
	}
}

func Greeting() {
	fmt.Printf("Welcome to Battleship in Go :))\n")
}

func InitializeGame() {

	for {
		fmt.Print("Please enter size\n>>> ")

		scanner.Scan()
		inputStr := strings.TrimSpace(scanner.Text())
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Invalid input. You must enter a valid number.")
			continue
		}

		if input < 1 {
			fmt.Println("Invalid input. You must enter a number larger than 0")
			continue
		}
		size = input
		break
	}

	for {
		fmt.Println("1. Easy\n2. Medium\n3. Hard")
		fmt.Print("Select difficulty\n>>> ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			difficulty = 50
			break
		case "2":
			difficulty = 25
			break
		case "3":
			difficulty = 10
			break
		default:
			fmt.Print("Invalid option. Try again\n>>> ")
		}
		if difficulty != -1 {
			break
		}
	}

	fmt.Printf("These are your chosen settings:\nSize: %d\nDifficulty: %d\n", size, difficulty)
	fmt.Print("Would you like to start the game?\n1. Yes\n2. No\n>>> ")

	scanner.Scan()

	input := scanner.Text()

	switch input {
	case "1":
		Game()
	case "2":
		return
	}
}

func Game() {
	fmt.Printf("game has started!!!\noptions:\nsize: %d\ndifficulty: %d \n\n", size, difficulty)

	currPlayer := player.NewPlayer(difficulty)
	currBoard := board.InitializeBoard(size)
	currBoard.PlaceShips([]int{5, 4, 3, 3, 2})

	for {
		if currPlayer.GetAttemptsLeft() < 1 {
			break
		}

		view.ShowGrid(currBoard)

		fmt.Println("Enter X coords: ")
		scanner.Scan()
		inputStr := strings.TrimSpace(scanner.Text())
		xinput, err := strconv.Atoi(inputStr)

		if err != nil {
			fmt.Println("Invalid input. You must enter a valid number.")
			continue
		}

		if xinput < 1 && xinput > size {
			fmt.Println("Invalid input. You must enter a number that on the board")
			continue
		}

		fmt.Println("Enter Y coords: ")
		scanner.Scan()
		inputStr = strings.TrimSpace(scanner.Text())
		yinput, err := strconv.Atoi(inputStr)

		if err != nil {
			fmt.Println("Invalid input. You must enter a valid number.")
			continue
		}

		if yinput < 1 && yinput > size {
			fmt.Println("Invalid input. You must enter a number that on the board")
			continue
		}

		currBoard.Attack(ship.Coordinate{X: xinput, Y: yinput})
	}

}
