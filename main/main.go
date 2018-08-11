package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	b "tictactoe/board"
)

func main() {
	fmt.Printf("Welcome!\nPlease hold on while we set everything up.\n")
	tttB := b.InitializeBoard()
	p1 := b.InitPlayer(1)
	p2 := b.InitPlayer(-1)

	var turn int8 = 1
	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("Things look good to go\n")
	fmt.Printf("Please enter your moves in the form <x,y>\n")

	gameEnd, winner := tttB.Winner(p1, p2)
	for !gameEnd {
		fmt.Printf(tttB.PrettyPrintState())
		move := b.Move{}
		switch turn % 2 {
		case 0:
			fmt.Printf("Player2>")
			input, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			} else {
				s := strings.Split(strings.Split(input, "\n")[0], ",")
				move.X, _ = strconv.Atoi(s[0])
				move.Y, _ = strconv.Atoi(s[1])
				fmt.Printf("Your move\nx: %d\ny: %d\n", move.X, move.Y)
				tttB.UpdateState(p2, move)
			}
		case 1:
			fmt.Printf("Player1>")
			input, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			} else {
				s := strings.Split(strings.Split(input, "\n")[0], ",")
				move.X, _ = strconv.Atoi(s[0])
				move.Y, _ = strconv.Atoi(s[1])
				fmt.Printf("Your move\nx: %d\ny: %d\n", move.X, move.Y)
				tttB.UpdateState(p1, move)
			}
		}
		turn++
		tttB.Turns++
		gameEnd, winner = tttB.Winner(p1, p2)
	}
	if winner != (b.Player{}) {
		fmt.Printf("Congratulations, Player%d!\n", winner.Mark)
	} else {
		fmt.Printf("Boo. Ties are no fun. Try again.")
	}
}
