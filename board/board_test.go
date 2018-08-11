package board

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

var p1 = Player{
	mark: 1,
	id:   uuid.New(),
}
var p2 = Player{
	mark: -1,
	id:   uuid.New(),
}

func TestInit(t *testing.T) {
	newBoard := TicTacToeBoard{}
	newBoard.ResetState()
	blankBoard := TicTacToeBoard{
		board: [3][3]int8{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
		Turns: 0,
	}
	if newBoard != blankBoard {
		t.Errorf("Arrays did not match\nReset board:\n%s", newBoard.PrettyPrintState())
	}
}

func TestPrettyPrint(t *testing.T) {
	newBoard := TicTacToeBoard{}
	newBoard.ResetState()
	blankBoard := newBoard.PrettyPrintState()
	line1 := fmt.Sprintf("   ||   ||   \n")
	line2 := fmt.Sprintf("   ||   ||   \n")
	line3 := fmt.Sprintf("   ||   ||   \n")
	spacer := "=============\n"
	blankBoardString := line1 + spacer + line2 + spacer + line3
	if blankBoardString != blankBoard {
		t.Errorf("Blank board string does not match intended result\nBlank board string:\n%s", blankBoard)
	}
}

func TestUpdateState(t *testing.T) {
	newBoard := TicTacToeBoard{}
	newBoard.ResetState()
	m1 := Move{
		X: 0,
		Y: 0,
	}
	m2 := Move{
		X: 1,
		Y: 1,
	}
	newBoard.UpdateState(p1, m1)
	newBoard.UpdateState(p2, m2)
	blankBoard := TicTacToeBoard{
		board: [3][3]int8{
			{1, 0, 0},
			{0, -1, 0},
			{0, 0, 0},
		},
		Turns: 0,
	}
	if newBoard != blankBoard {
		t.Errorf("Board does not match expected value\nPrinted value:\n%s", newBoard.PrettyPrintState())
	}
}

func TestWin(t *testing.T) {
	winBoard := TicTacToeBoard{
		board: [3][3]int8{
			{1, 1, 1},
			{-1, -1, 0},
			{-1, 0, 0},
		},
		Turns: 6,
	}
	gameEnd, winner := winBoard.Winner(p1, p2)
	if !gameEnd && winner != p1 {
		t.Errorf("Winner not returned\nGame state: %t\nWinner: %d", gameEnd, winner.id)
	}
}

func TestTie(t *testing.T) {
	tieBoard := TicTacToeBoard{
		board: [3][3]int8{
			{1, -1, 1},
			{-1, -1, 1},
			{-1, 1, -1},
		},
		Turns: 9,
	}
	gameEnd, tie := tieBoard.Winner(p1, p2)
	if !gameEnd && tie != (Player{}) {
		t.Errorf("Tie not returned\nGame state: %t", gameEnd)
	}
}
