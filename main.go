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
		fmt.Println("fixed")
	case strategyMixed:
		fmt.Println("mixed")
	default:
		fmt.Println("compound")
	}
}


