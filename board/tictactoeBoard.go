package board

import (
	"fmt"
	//"math"
)

// TicTacToeBoard basic 3x3 2D array
type TicTacToeBoard struct {
	board [3][3]int8
	Turns uint
}

func printState(val int8) string {
	switch val {
	case 0:
		return " "
	case 1:
		return "X"
	case -1:
		return "O"
	default:
		panic(fmt.Sprintf("Invalid board placement: %d", val))
	}
}

// PrettyPrintState print current tttB state
func (tttB TicTacToeBoard) PrettyPrintState() string {
	var line1 = fmt.Sprintf(" %s || %s || %s \n", printState(tttB.board[0][0]), printState(tttB.board[0][1]), printState(tttB.board[0][2]))
	var line2 = fmt.Sprintf(" %s || %s || %s \n", printState(tttB.board[1][0]), printState(tttB.board[1][1]), printState(tttB.board[1][2]))
	var line3 = fmt.Sprintf(" %s || %s || %s \n", printState(tttB.board[2][0]), printState(tttB.board[2][1]), printState(tttB.board[2][2]))
	var spacer = "=============\n"
	return line1 + spacer + line2 + spacer + line3

}

// UpdateState increment tttB by one move
func (tttB *TicTacToeBoard) UpdateState(player Player, move Move) (bool, error) {
	var illegal = (move.X > 3) || (move.Y > 3)
	if illegal {
		return false, fmt.Errorf("[ERROR]: illegal move; x: %d, y: %d", move.X, move.Y)
	}
	tttB.board[move.X][move.Y] = player.Mark
	return true, nil
}

// ResetState overwrite board with all 0's
func (tttB *TicTacToeBoard) ResetState() {
	for i := range tttB.board {
		for j := range tttB.board[i] {
			tttB.board[i][j] = 0
		}
	}
	tttB.Turns = 0
}

// InitializeBoard init method for TicTacToeBoard in place of exporting struct
func InitializeBoard() TicTacToeBoard {
	newBoard := TicTacToeBoard{}
	newBoard.ResetState()
	return newBoard
}

func p1orp2(p1 Player, p2 Player, sum int8) (bool, Player) {
	switch sum {
	case -3:
		return true, p2
	case 3:
		return true, p1
	default:
		return false, Player{}
	}
}

func abs(n int8) int8 {
	if n < 0 {
		return -n
	}
	return n
}

// Winner assess win condition of current board
// return winning Player struct if there is a winner
func (tttB TicTacToeBoard) Winner(p1 Player, p2 Player) (bool, Player) {
	if tttB.Turns < 5 {
		return false, Player{}
	}
	c1 := tttB.board[0][0] + tttB.board[0][1] + tttB.board[0][2]
	if abs(c1) == 3 {
		return p1orp2(p1, p2, c1)
	}
	c2 := tttB.board[1][0] + tttB.board[1][1] + tttB.board[1][2]
	if abs(c2) == 3 {
		return p1orp2(p1, p2, c2)
	}
	c3 := tttB.board[2][0] + tttB.board[2][1] + tttB.board[2][2]
	if abs(c3) == 3 {
		return p1orp2(p1, p2, c3)
	}
	r1 := tttB.board[0][0] + tttB.board[1][0] + tttB.board[2][0]
	if abs(r1) == 3 {
		return p1orp2(p1, p2, r1)
	}
	r2 := tttB.board[1][0] + tttB.board[1][1] + tttB.board[1][2]
	if abs(r2) == 3 {
		return p1orp2(p1, p2, r2)
	}
	r3 := tttB.board[2][0] + tttB.board[2][1] + tttB.board[2][2]
	if abs(r3) == 3 {
		return p1orp2(p1, p2, r3)
	}
	d1 := tttB.board[0][0] + tttB.board[1][1] + tttB.board[2][2]
	if abs(d1) == 3 {
		return p1orp2(p1, p2, d1)
	}
	d2 := tttB.board[2][0] + tttB.board[1][1] + tttB.board[0][2]
	if abs(d2) == 3 {
		return p1orp2(p1, p2, d2)
	}
	if tttB.Turns == 9 {
		return true, Player{}
	}
	return false, Player{}
}
