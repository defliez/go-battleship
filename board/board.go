package board

import (
	"math/rand"
	"time"

	"valedv.com/battleship/ship"
)

type Board struct {
	Size  		int
	Ships 		[]*ship.Ship
	Grid		map[ship.Coordinate]bool
	DisplayGrid map[ship.Coordinate]
}

func InitializeBoard(size int) *Board {
	grid := make(map[ship.Coordinate]bool, size*size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			grid[ship.Coordinate{X: x, Y: y}] = false
		}
	}
	return &Board{
		Size:  size,
		Ships: []*ship.Ship{},
		Grid:  grid,
	}
}

func (b *Board) PlaceShips(lengths []int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, length := range lengths {
		for {
			horiz := rand.Intn(2) == 0
			var x0, y0 int
			if horiz {
				x0 = rng.Intn(b.Size - length + 1)
				y0 = rng.Intn(b.Size)
			} else {
				x0 = rng.Intn(b.Size)
				y0 = rng.Intn(b.Size - length + 1)
			}

			// build candidate coords
			coords := make([]ship.Coordinate, length)
			for i := 0; i < length; i++ {
				if horiz {
					coords[i] = ship.Coordinate{X: x0 + i, Y: y0}
				} else {
					coords[i] = ship.Coordinate{X: x0, Y: y0 + i}
				}
			}

			if b.canPlace(coords) {
				b.addShip(coords)
				break
			}
			// else retry
		}
	}
}

func (b *Board) canPlace(coords []ship.Coordinate) bool {
	for _, c := range coords {
		if b.Grid[c] {
			return false
		}
	}
	return true
}

func (b *Board) addShip(coords []ship.Coordinate) {
	for _, c := range coords {
		b.Grid[c] = true
	}
	b.Ships = append(b.Ships, ship.NewShip(coords))
}

func GetBoard() (
