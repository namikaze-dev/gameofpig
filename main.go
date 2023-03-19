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

func MixedStrategy(firstN, secondMin, secondMax, games int) ([]int, []int) {
	var firstTotal []int
	var secondTotal []int

	for secondN := secondMin; secondN <= secondMax; secondN++ {
		if secondN == firstN {
			continue
		}

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

		firstTotal = append(firstTotal, firstWins)
		secondTotal = append(secondTotal, secondWins)
	}

	return firstTotal, secondTotal
}

func CompoundStrategy(firstMin, firstMax, secondMin, secondMax, games int) ([]int, []int) {
	var firstTotal []int
	var secondTotal []int

	for firstN := firstMin; firstN <= firstMax; firstN++ {

		firstWins, secondWins := 0, 0
		for secondN := secondMin; secondN <= secondMax; secondN++ {
			if secondN == firstN {
				continue
			}

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
		}

		firstTotal = append(firstTotal, firstWins)
		secondTotal = append(secondTotal, secondWins)
	}

	return firstTotal, secondTotal
}
