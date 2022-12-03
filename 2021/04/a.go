package main

import (
	"bufio"
	"fmt"
	"github.com/buckket/adventofcode/2021/04/bingo"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []int
	var boards []bingo.Board

	ok := scanner.Scan()
	if !ok {
		panic("Invalid input")
	}
	numbersString := strings.Split(scanner.Text(), ",")
	for _, number := range numbersString {
		n, err := strconv.ParseInt(number, 10, 32)
		if err != nil {
			panic("ParseInt failed")
		}
		numbers = append(numbers, int(n))
	}

	var boardLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			boardLines = append(boardLines, line)
		} else {
			board := bingo.Board{}
			board.Init(boardLines)
			boards = append(boards, board)
			boardLines = nil
		}
	}

	for _, number := range numbers {
		for i, board := range boards {
			if result, bingo := board.Check(number); bingo {
				fmt.Printf("Bingo on board %d: %d\n", i, result)
				return
			}
		}
	}
}
