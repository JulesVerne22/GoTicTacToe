package game

import (
	"testing"
)

func TestBoardSetup(t *testing.T) {
	var board Board

	if len(board) != 3 {
		t.Error("board does not have 3 rows, it has", len(board), "rows")
	}

	if len(board) == 0 {
		t.Error("board does not have any rows")
	} else if len(board[0]) != 3 {
        t.Error("board does not have 3 columns, it has", len(board[0]), "columns")
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] != 0 {
                t.Error("board is not filled with 0's, the value is", string(board[i][j]))
            }
        }
    }
}

func TestClearBoard(t *testing.T) {
    var board Board

    board.ClearBoard()

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] != ' ' {
                t.Error("board is not cleared, the value is", string(board[i][j]))
            }
        }
    }
}

func TestPlaceChar(t *testing.T) {
    var board Board

    board.ClearBoard()

    board.PlaceChar('X', 0, 0)

    if board[0][0] != 'X' {
        t.Error("placed char does not match, expected X and got", string(board[0][0]))
    }
}

func TestCheckWin(t *testing.T) {
    var board Board

    board.ClearBoard()

    if board.CheckWin('X', 0, 0) {
        t.Error("board shows win when cleared")
    }

    board.PlaceChar('X', 0, 0) 
    board.PlaceChar('X', 0, 1) 
    board.PlaceChar('X', 0, 2) 
    board.PlaceChar('X', 1, 0)
    board.PlaceChar('X', 2, 0)
    board.PlaceChar('X', 1, 1)
    board.PlaceChar('X', 2, 2)

    if !board.CheckWin('X', 0, 2) {
        t.Error("doesn't show win with 3 matching chars in a row")
    }

    if !board.CheckWin('X', 2, 0) {
        t.Error("board doesn't show win with 3 matching chars in a column")
    }

    if !board.CheckWin('X', 2, 2) {
        t.Error("board doesn't show win with 3 matching chars on a diagonal")
    }
}

func TestCheckTie(t *testing.T) {
    var board Board

    board.ClearBoard()

    if board.CheckTie() {
        t.Error("board shows a tie when the board is cleared")
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            board.PlaceChar('X', i, j)
        }
    }

    if !board.CheckTie() {
        t.Error("board does not show a tie when filled")
    }
}
