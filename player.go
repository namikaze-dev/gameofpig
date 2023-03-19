package main

import (
	"math/rand"
	"time"
)

type Player struct {
	player bool
	total  int
	rand   *rand.Rand
}

func NewPlayer(p bool) Player {
	return Player{
		player: p,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (p *Player) score() int {
	return p.total
}

// Rolls die untill current score >= N, returns current score.
func (p *Player) HoldAfterN(n int) int {
	var score int
	for score <= n {
		got := p.rand.Intn(7)
		// player scores nothing
		if got == 1 {
			return 0
		}
		score += got
	}
	p.total += score
	return score
}
