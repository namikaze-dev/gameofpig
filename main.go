package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	strategyFixed    = "fixed"
	strategyMixed    = "mixed"
	strategyCompound = "compound"
)

var (
	games  = 10
	logger = log.New(os.Stdout, "pig: ", 0)
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		logger.Fatalln("two arguments required: got", args)

	}

	a, b := args[0], args[1]
	gameType := resolveType(a, b)

	switch gameType {
	case strategyFixed:
		playStrategyFixed(a, b)
	case strategyMixed:
		fmt.Println("mixed")
	default:
		fmt.Println("compound")
	}
}

func playStrategyFixed(a, b string) {
	holdA, err := strconv.Atoi(a)
	if err != nil {
		logger.Fatalln("invalid not integer argument:", a)
	}

	holdB, err := strconv.Atoi(b)
	if err != nil {
		logger.Fatalln("invalid not integer argument:", b)
	}

	winsA, winsB := FixedStrategy(holdA, holdB, games)
	percentA := int(float32(winsA) / float32(games) * 100)
	percentB := int(float32(winsB) / float32(games) * 100)

	fmt.Printf("Holding at %v vs Holding at %v: wins: %v/%v (%v.0%%), losses: %v/%v (%v%%)\n",
		a, b, winsA, games, percentA, winsB, games, percentB,
	)
}