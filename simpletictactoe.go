// Inspired from the Tour of Go exercise, I wanted to try making
// a very simple command line tic tac toe
// https://go.dev/tour/moretypes/14

package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Simple command line coloring helpers from
// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

type Board [][]string

var guide Board = [][]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

type MoveCoordinates struct{ x, y int }

// The user will give input with 1-9 so this
// will help map to the board slice of slices
// which the script access like [x][y]
var numToCoord = map[int]MoveCoordinates{
	1: {0, 0},
	2: {0, 1},
	3: {0, 2},
	4: {1, 0},
	5: {1, 1},
	6: {1, 2},
	7: {2, 0},
	8: {2, 1},
	9: {2, 2},
}
var EmptyCoords = MoveCoordinates{-1, -1}

var settingsPlayerCount = 1

// No computer logic implemented yet
// var settingsDifficulty = "random"
var settingsHumanPlayer = "X"

func main() {
	// Use time to seed rand
	rand.Seed(time.Now().UnixNano())
	getSettings()

	promptTurn(initBoard(), 0, "", EmptyCoords)
}

func getSettings() {
	playerCount := getCharInput("Players: (1) or 2")

	playerInt, conversionErr := strconv.Atoi(playerCount)
	if conversionErr != nil || (playerInt != 1 && playerInt != 2) {
		fmt.Println("Players: ", settingsPlayerCount)
	} else {
		settingsPlayerCount = playerInt
	}

	if settingsPlayerCount == 1 {
		// fmt.Println("Right now, the only computer difficulty is 'Random Play'")

		playAs := getCharInput("Play as (X) or O?")
		if strings.ToLower(playAs) == "o" {
			settingsHumanPlayer = "O"
		} else {
			settingsHumanPlayer = "X"
		}
		fmt.Println("Playing as " + settingsHumanPlayer)
	}
}

func initBoard() Board {
	return [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}

func showBoard(board Board, lastMove MoveCoordinates) {
	fmt.Println("Game     Options")
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if j > 0 {
				fmt.Printf(" ")
			}
			// Highlight the last move in Green
			highlight := ""
			if lastMove.x == i && lastMove.y == j {
				highlight = Green
			}
			fmt.Printf(highlight+"%s"+Reset, board[i][j])
		}
		fmt.Printf("    ")

		// Show a guide with the available numbers, maps to
		// the positions on the board
		for j := 0; j < len(guide[i]); j++ {
			if j > 0 {
				fmt.Printf(" ")
			}
			// Only show numbers left to pick
			if board[i][j] == "_" {
				fmt.Printf("%s", guide[i][j])
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func promptTurn(board Board, turn int, msg string, lastMove MoveCoordinates) {

	fmt.Println("------------------------")
	if turn == 0 {
		fmt.Println("Welcome to Tic Tac Toe")
	}

	// Determine X or O based on the turn
	player := "X"
	if turn%2 == 1 {
		player = "O"
	}

	showBoard(board, lastMove)

	// Take status from outside to print below the board
	fmt.Println(msg)

	if player == "O" && CheckForWin("X", board) {
		showGameOver("X")
		return
	}
	if player == "X" && CheckForWin("O", board) {
		showGameOver("O")
		return
	}
	if turn == 9 {
		showGameOver("")
		return
	}

	// Determine if turn is taken by player(s) or computer
	if settingsPlayerCount == 2 || (settingsHumanPlayer == player) {
		getUserInputForTurn(board, turn, player, lastMove)
	} else {
		takeComputerTurn(board, turn, player)
	}
}

// Get selection from the user
// Validate selection
// Restart turn with message if invalid selection or,
// Update board and advance to next turn
func getUserInputForTurn(board Board, turn int, player string, lastMove MoveCoordinates) {
	mv := getCharInput(player + "'s Turn. Select 1-9:")
	mvAsInt, conversionErr := strconv.Atoi(mv)
	if conversionErr != nil || (mvAsInt < 1 || mvAsInt > 9) {
		promptTurn(board, turn, Red+"Invalid space (must be 1-9)..."+Reset, lastMove)
		return
	}

	if mvAsInt < 1 || mvAsInt > 9 {
		promptTurn(board, turn, Red+"Invalid space (must be 1-9)..."+Reset, lastMove)
		return
	}
	coords := numToCoord[mvAsInt]
	curValue := board[coords.x][coords.y]
	if curValue != "_" {
		promptTurn(board, turn, Red+"Pick an open space..."+Reset, lastMove)
		return
	}

	board[coords.x][coords.y] = player

	promptTurn(board, turn+1, "", coords)
}

func takeComputerTurn(board Board, turn int, player string) {
	if turn > 8 {
		log.Fatal("Game should've ended by now...")
	}
	fmt.Print("Computer's turn")
	for x := 0; x < 3; x++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
	move := 0
	var coords MoveCoordinates
	for move == 0 {
		randomPick := rand.Intn(9) + 1
		coords = numToCoord[randomPick]
		curValue := board[coords.x][coords.y]
		if curValue == "_" {
			move = randomPick
		}
	}
	board[coords.x][coords.y] = player

	promptTurn(board, turn+1, "", coords)
}

func showGameOver(winner string) {
	if winner != "" {
		fmt.Println(Green+"Congratulations", winner, ", you've won!"+Reset)
	} else {
		fmt.Println(Yellow + "The game has ended in a tie" + Reset)
	}

	playAgain := getCharInput("Play again? [Y]es [N]o")
	if strings.ToLower(playAgain) != "n" && strings.ToLower(playAgain) != "no" {
		promptTurn(initBoard(), 0, "", EmptyCoords)
	}
}

func CheckForWin(player string, board Board) bool {
	result := false
	switch {
	case board[0][0] == player && board[0][1] == player && board[0][2] == player:
		result = true
	case board[1][0] == player && board[1][1] == player && board[1][2] == player:
		result = true
	case board[2][0] == player && board[2][1] == player && board[2][2] == player:
		result = true
	case board[0][0] == player && board[1][0] == player && board[2][0] == player:
		result = true
	case board[0][1] == player && board[1][1] == player && board[2][1] == player:
		result = true
	case board[0][2] == player && board[1][2] == player && board[2][2] == player:
		result = true
	case board[0][0] == player && board[1][1] == player && board[2][2] == player:
		result = true
	case board[0][2] == player && board[1][1] == player && board[2][0] == player:
		result = true
	}
	return result
}

// Helper that displays a `msg` and gets user's char input
func getCharInput(msg string) string {
	fmt.Println(msg)
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		if err.Error() != "unexpected newline" {
			fmt.Println(err)
		}
	}
	return input
}
