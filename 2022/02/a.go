package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won

func main() {
	// LUT: 0 for Rock, 1 for Paper, and 2 for Scissors
	var score_lut = [3][3]int{
		{3 + 1, 6 + 2, 0 + 3},
		{0 + 1, 3 + 2, 6 + 3},
		{6 + 1, 0 + 2, 3 + 3}}

	var move_lut = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2}

	var score int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		if len(moves) != 2 {
			break
		}
		score += score_lut[move_lut[moves[0]]][move_lut[moves[1]]]
	}

	fmt.Printf("Score: %d\n", score)
}
