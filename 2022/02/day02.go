package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var moveLut = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"X": 0,
	"Y": 1,
	"Z": 2}

func Part1(input io.Reader) string {
	// LUT: 0 for Rock, 1 for Paper, and 2 for Scissors
	var scoreLut = [3][3]int{
		{3 + 1, 6 + 2, 0 + 3},
		{0 + 1, 3 + 2, 6 + 3},
		{6 + 1, 0 + 2, 3 + 3}}

	var score int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		if len(moves) != 2 {
			break
		}
		score += scoreLut[moveLut[moves[0]]][moveLut[moves[1]]]
	}

	return fmt.Sprintf("%d", score)
}

func Part2(input io.Reader) string {
	// First index: 0 for Rock, 1 for Paper, and 2 for Scissors
	// Second index: 0 for Loss, 1 for Draw, and 2 for Win
	var scoreLut = [3][3]int{
		{0 + 3, 3 + 1, 6 + 2},
		{0 + 1, 3 + 2, 6 + 3},
		{0 + 2, 3 + 3, 6 + 1}}

	var score int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		if len(moves) != 2 {
			break
		}
		score += scoreLut[moveLut[moves[0]]][moveLut[moves[1]]]
	}

	return fmt.Sprintf("%d", score)
}

func main() {
	var partFlag = flag.Int("p", 0, "select part")
	flag.Parse()

	switch *partFlag {
	case 1:
		fmt.Println(Part1(os.Stdin))
	case 2:
		fmt.Println(Part2(os.Stdin))
	default:
		fmt.Println(fmt.Errorf("unknown part number %d", *partFlag))
	}
}
