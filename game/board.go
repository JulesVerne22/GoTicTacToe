package game

import (
	"fmt"
)

type Board [3][3]rune

func (b *Board) ClearBoard() {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = ' '
		}
	}
}

func (b *Board) PlaceChar(char rune, row int, col int) {
	b[row][col] = char
}

func (b Board) CheckTie() bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}

func (b Board) PrintTie() {
	fmt.Printf("Oh no... It was a tie...\n")
}

func (b Board) PrintBoard() {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if j != len(b[i])-1 {
				fmt.Printf("%s | ", string(b[i][j]))
			} else {
				fmt.Printf("%s\n", string(b[i][j]))
			}
		}

		if i != len(b)-1 {
			fmt.Println("---------")
		} else {
			fmt.Println()
		}
	}
}

func (b Board) CheckWin(char rune, row int, col int) bool {
	switch {
	case row == 0 && col == 0:
		if b[0][1] == char && b[0][2] == char {
			return true
		} else if b[1][0] == char && b[2][0] == char {
			return true
		} else if b[1][1] == char && b[2][2] == char {
			return true
		}
	case row == 0 && col == 1:
		if b[0][0] == char && b[0][2] == char {
			return true
		} else if b[1][1] == char && b[2][1] == char {
			return true
		}
	case row == 0 && col == 2:
		if b[0][0] == char && b[0][1] == char {
			return true
		} else if b[1][2] == char && b[2][2] == char {
			return true
		} else if b[1][1] == char && b[2][0] == char {
			return true
		}
	case row == 1 && col == 0:
		if b[1][1] == char && b[1][2] == char {
			return true
		} else if b[0][0] == char && b[2][0] == char {
			return true
		}
	case row == 1 && col == 1:
		if b[1][0] == char && b[1][2] == char {
			return true
		} else if b[0][1] == char && b[2][1] == char {
			return true
		} else if b[0][0] == char && b[2][2] == char {
			return true
		} else if b[0][2] == char && b[2][0] == char {
			return true
		}
	case row == 1 && col == 2:
		if b[1][0] == char && b[1][1] == char {
			return true
		} else if b[0][2] == char && b[2][2] == char {
			return true
		}
	case row == 2 && col == 0:
		if b[2][1] == char && b[2][2] == char {
			return true
		} else if b[0][0] == char && b[1][0] == char {
			return true
		} else if b[1][1] == char && b[0][2] == char {
			return true
		}
	case row == 2 && col == 1:
		if b[2][0] == char && b[2][2] == char {
			return true
		} else if b[0][1] == char && b[1][1] == char {
			return true
		}
	case row == 2 && col == 2:
		if b[2][0] == char && b[2][1] == char {
			return true
		} else if b[0][2] == char && b[1][2] == char {
			return true
		} else if b[0][0] == char && b[1][1] == char {
			return true
		}
	}
	return false
}
