package main_test

import (
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