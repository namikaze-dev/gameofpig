package main_test

import (
	"reflect"
	"testing"

	"github.com/namikaze-dev/gameofpig"
)

func TestPlayer(t *testing.T) {
	// true for player 1, false for player two 
	player := main.NewPlayer(true)

	// hold after 10
	want := 10
	got := player.HoldAfterN(want)

	if got < want && got != 0 {
		t.Errorf("got %v want >= %v", got, want)
	}

	// hold after 99, increase chance of scoring 0
	want = 99
	got = player.HoldAfterN(want)

	if got < want && got != 0 {
		t.Errorf("got %v want >= %v", got, want)
	}
}

func TestFixedStrategyGamePlay(t *testing.T) {
	firstN, secondN, games := 10, 15, 10
	firstWins, secondWins := main.FixedStrategy(firstN, secondN, games)

	got := firstWins + secondWins
	if got != games {
		t.Errorf("total games: got %v want %v", got, games)
	}

	// higher chance of rand results 1000 games 
	firstN, secondN, games = 10, 15, 1000
	firstWins, secondWins = main.FixedStrategy(firstN, secondN, games)

	got = firstWins + secondWins
	if got != games {
		t.Errorf("total games: got %v want %v", got, games)
	}

	// NOTE: there's a small chance that wins would be equal, expect occasional failure
	if firstWins == secondWins {
		t.Errorf("wins should not be equal, got %v and %v", firstWins, secondWins)
	}
}

func TestMixedStrategyGamePlay(t *testing.T) {
	firstN, secondMin, secondMax, games := 21, 1, 100, 10
	firstTotal, secondTotal := main.MixedStrategy(firstN, secondMin, secondMax, games)

	got := len(firstTotal)
	if got != (secondMax - secondMin) {
		t.Errorf("first player total result: got %v want %v", got, games)
	}

	got = len(secondTotal)
	if got != (secondMax - secondMin) {
		t.Errorf("second player total result: got %v want %v", got, games)
	}

	// total wins should not be equal
	if reflect.DeepEqual(firstTotal, secondTotal) {
		t.Errorf("total wins: got %v and %v", firstTotal, secondTotal)
	}
}