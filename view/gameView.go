package view

import (
	"fmt"

	"valedv.com/battleship/board"
	"valedv.com/battleship/ship"
)

// If revealShips is true then occupied cells show “O”, otherwise they print as “~”.
// func ShowGrid(b *board.Board, revealShips bool) {
// ShowGrid renders only hits and misses to stdout.
func ShowGrid(b *board.Board) {
	size := b.Size

	// column headers
	fmt.Print("   ")
	for x := 0; x < size; x++ {
		fmt.Printf("%2d", x)
	}
	fmt.Println()

	for y := 0; y < size; y++ {
		// row header
		fmt.Printf("%2d ", y)
		for x := 0; x < size; x++ {
			c := ship.Coordinate{X: x, Y: y}
			st := b.Grid[c]
			var ch string
			switch st {
			case board.StateHit:
				ch = "X"
			case board.StateMissed:
				ch = "·"
			default: // StateNone or StateOccupied
				ch = "~"
			}
			fmt.Printf("%2s", ch)
		}
		fmt.Println()
	}
	fmt.Println()
}

func ShowStats(hitCount, missCount int, b *board.Board) {
	total := hitCount + missCount
	accuracy := 0.0
	if total > 0 {
		accuracy = float64(hitCount) / float64(total) * 100
	}
	sunk := 0
	for _, s := range b.Ships {
		if s.IsSunk() {
			sunk++
		}
	}
	fmt.Printf("Shots: %d  Hits: %d  Misses: %d  Accuracy: %.1f%%  Ships sunk: %d/%d\n\n",
		total, hitCount, missCount, accuracy, sunk, len(b.Ships))
}
