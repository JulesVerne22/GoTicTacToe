package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/JulesVerne22/GoTicTacToe/game"
	"github.com/JulesVerne22/GoTicTacToe/player"
)

func printIntro() {
	fmt.Println("------------------------------------------------")
	fmt.Println("   Hi, and welcome to a game of Tic Tac Toe!")
	fmt.Println("------------------------------------------------")
	fmt.Println("\nFirst I need to know who's playing. I'll need each",
		"player to enter their name and the character they want to use",
		"during the game.")
	fmt.Println("\nAfter entering your responses press Enter...")
}

func clearTerm(system string) {
	if system == "linux" || system == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if system == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func playAgain(play string) string {
	for strings.ToLower(play) != "y" && strings.ToLower(play) != "n" {
		fmt.Println("Would you like to play again? (y/n)")
		fmt.Scanln(&play)
	}

	return play
}

func main() {
	var board game.Board
	board.ClearBoard()

	var player1 player.Player
	var player2 player.Player

	printIntro()

	player1.SetupPlayer("Player1", []string{}, []string{})
	player2.SetupPlayer("Player2", []string{player1.Name}, []string{string(player1.Char)})

	var play string
	var row, col int

	for {
		clearTerm(runtime.GOOS)
		fmt.Printf("%s's Turn\n", player1.Name)
		fmt.Printf("-------------------------------\n\n")
		board.PrintBoard()

		row, col = player1.GetPlay(board)
		board.PlaceChar(player1.Char, row, col)
		if board.CheckWin(player1.Char, row, col) {
			clearTerm(runtime.GOOS)
			fmt.Printf("%s's Turn\n", player1.Name)
			fmt.Printf("-------------------------------\n\n")
			board.PrintBoard()
			player1.PrintWin()
			board.ClearBoard()
			play = playAgain(play)
		} else if board.CheckTie() {
			clearTerm(runtime.GOOS)
			fmt.Printf("%s's Turn\n", player1.Name)
			fmt.Printf("-------------------------------\n\n")
			board.PrintBoard()
			board.PrintTie()
			board.ClearBoard()
			play = playAgain(play)
		}

		if strings.ToLower(play) == "n" {
			break
		}
		play = ""

		clearTerm(runtime.GOOS)
		fmt.Printf("%s's Turn\n", player2.Name)
		fmt.Printf("-------------------------------\n\n")
		board.PrintBoard()

		row, col = player2.GetPlay(board)
		board.PlaceChar(player2.Char, row, col)
		if board.CheckWin(player2.Char, row, col) {
			clearTerm(runtime.GOOS)
			fmt.Printf("%s's Turn\n", player2.Name)
			fmt.Printf("-------------------------------\n\n")
			board.PrintBoard()
			player2.PrintWin()
			board.ClearBoard()
			play = playAgain(play)
		} else if board.CheckTie() {
			clearTerm(runtime.GOOS)
			fmt.Printf("%s's Turn\n", player2.Name)
			fmt.Printf("-------------------------------\n\n")
			board.PrintBoard()
			board.PrintTie()
			board.ClearBoard()
			play = playAgain(play)
		}

		if strings.ToLower(play) == "n" {
			break
		}
		play = ""
	}
}
