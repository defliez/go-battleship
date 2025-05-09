package ship

type Coordinate struct {
	X, Y int
}

type Ship struct {
	Positions []Coordinate
	hits      map[Coordinate]bool
}

func NewShip(positions []Coordinate) *Ship {
	return &Ship{
		Positions: positions,
		hits:      make(map[Coordinate]bool),
	}
}

func (s *Ship) Hit(coord Coordinate) bool {
	for _, p := range s.Positions {
		if p == coord {
			s.hits[coord] = true
			return true
		}
	}
	return false
}

func (s *Ship) IsSunk() bool {
	return len(s.hits) == len(s.Positions)
}
