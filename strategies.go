package main

func FixedStrategy(firstN, secondN, games int) (int, int) {
	firstWins, secondWins := 0, 0

	for i := 0; i < games; i++ {
		first := NewPlayer(true)
		second := NewPlayer(false)

		a, b := play(&first, &second, firstN, secondN)
		firstWins += a
		secondWins += b
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

			a, b := play(&first, &second, firstN, secondN)
			firstWins += a
			secondWins += b
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
		firstWins, secondWins := CompoundStrategyHelper(secondMin, secondMax, firstN, games)
		firstTotal = append(firstTotal, firstWins)
		secondTotal = append(secondTotal, secondWins)
	}

	return firstTotal, secondTotal
}