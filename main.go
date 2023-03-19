package main

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
