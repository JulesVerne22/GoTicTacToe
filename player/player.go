package player

import (
	"fmt"
	"strconv"

	"github.com/JulesVerne22/GoTicTacToe/game"
)

type Player struct {
	Name string
	Char rune
}

func (p Player) GetPlay(board game.Board) (row int, col int) {
	fmt.Println("Choose a square by typing its number and press enter")
	fmt.Println()

	count := 1
	avail := map[int]bool{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == ' ' && j != len(board[i])-1 {
				fmt.Printf("%d | ", count)
				avail[count] = true
			} else if board[i][j] == ' ' {
				fmt.Printf("%d\n", count)
				avail[count] = true
			} else if j != len(board[i])-1 {
				fmt.Printf("  | ")
			} else {
				fmt.Printf(" \n")
			}
			count++
		}

		if i != len(board)-1 {
			fmt.Println("---------")
		} else {
			fmt.Println()
		}
	}

	var choiceStr string
	var choice int
	var err error = nil
	for choiceStr == "" {
		fmt.Printf("Square: ")
		fmt.Scanln(&choiceStr)

		choice, err = strconv.Atoi(choiceStr)
		if err != nil {
			choiceStr = ""
		} else if !avail[choice] {
			choiceStr = ""
		}
	}

	return (choice - 1) / 3, (choice - 1) % 3
}

func (p Player) PrintWin() {
	fmt.Printf("Congratulations, %s! You are the winner!!\n", p.Name)
}

func (p *Player) SetupPlayer(playerNum string, disNames []string, disChars []string) {
	var pName string
	var pCharStr string

	for pName == "" || searchString(disNames, pName) {
		fmt.Print(playerNum, " Name: ")
		fmt.Scanln(&pName)
	}

	for pCharStr == "" || searchString(disChars, string(pCharStr)) {
		fmt.Print(playerNum, " Character: ")
		fmt.Scanln(&pCharStr)
	}

	pChar := []rune(pCharStr)

	p.Name = pName
	p.Char = pChar[0]
}

func searchString(strings []string, search string) bool {
	for _, v := range strings {
		if search == v {
			return true
		}
	}
	return false
}
