package main

import "strconv"

func play(first, second *Player, firstN, secondN int) (int, int) {
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
			return 1, 0
		}

		if second.score() >= 100 {
			return 0, 1
		}
	}
}

func CompoundStrategyHelper(min, max, firstN, games int) (int, int) {
	firstWins, secondWins := 0, 0
	for secondN := min; secondN <= max; secondN++ {
		if secondN == firstN {
			continue
		}

		for i := 0; i < games; i++ {
			first := NewPlayer(true)
			second := NewPlayer(false)

			a, b := play(&first, &second, firstN, secondN)
			firstWins += a
			secondWins += b
		}
	}

	return firstWins, secondWins
}

func resolveType(a, b string) string {
	_, err := strconv.Atoi(a)
	_, err2 := strconv.Atoi(b)

	if err == nil && err2 == nil {
		return strategyFixed
	}

	if err == nil || err2 == nil {
		return strategyMixed
	}

	return strategyCompound
}
