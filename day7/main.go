package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type HandStrength int64

const (
	High HandStrength = iota
	One
	Two
	Tree
	Full
	Four
	Five
)

type PokerHand struct {
	cards    string
	handType int
	bid      int
}

func splitLine(line, t string) []string {
	r := regexp.MustCompile(t)
	return r.Split(line, -1)
}

func checkType(hand string) {
}

func insertHands(hands []PokerHand, h PokerHand) []PokerHand {
	if len(hands) == 0 {
		return append(hands, h)
	}
	return hands
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error while reading file: %v\n", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	hands := []PokerHand{}
	for scanner.Scan() {
		line := splitLine(strings.TrimSpace(scanner.Text()), " ")
		bid, _ := strconv.Atoi(line[1])
		hand := PokerHand{line[0], bid}
		insertHands(hands, hand)
	}
}
