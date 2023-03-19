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
		rand:   rand.New(rand.NewSource(time.Now().UnixMicro())),
	}
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

func (p *Player) score() int {
	return p.total
}

func FixedStrategy(firstN, secondN, games int) (int, int) {
	firstWins, secondWins := 0, 0

	for i := 0; i < games; i++ {
		first := NewPlayer(true)
		second := NewPlayer(false)
		turn := true

		for {
			if turn {
				first.HoldAfterN(firstN)
			} else {
				second.HoldAfterN(secondN)
			}
			// switch turns
			turn = !turn

			if first.score() >= 100 {
				firstWins++
				break
			}

			if second.score() >= 100 {
				secondWins++
				break
			}
		}
	}

	return firstWins, secondWins
}
