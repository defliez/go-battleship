package player

type Player struct {
	attemptsLeft, totalAttempts int
}

func NewPlayer(attempts int) *Player {
	return &Player{
		totalAttempts: 0,
		attemptsLeft: attempts,
	}
}

func (p *Player) DecreaseAttempt() {
	p.totalAttempts--
}

func (p *Player) GetTotalAttempts() int {
	return p.totalAttempts
}

func (p *Player) GetAttemptsLeft() int {
	return p.attemptsLeft
}
