package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(input io.Reader) string {
	var sum int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)

		nLeft := -1
		nRight := -1

		for i := 0; i < length; i++ {
			if line[i] <= '9' && line[i] >= '0' {
				nLeft = i
				break
			}
		}

		if nLeft == -1 {
			continue
		}

		for i := length - 1; i >= nLeft; i-- {
			if line[i] <= '9' && line[i] >= '0' {
				nRight = i
				break
			}
		}

		x, err := strconv.Atoi(string(line[nLeft]))
		if err != nil {
			log.Fatal()
		}

		y, err := strconv.Atoi(string(line[nRight]))
		if err != nil {
			log.Fatal()
		}

		sum += x*10 + y
	}

	return fmt.Sprintf("%d", sum)
}

type digit struct {
	key   string
	value int
}

func Part2(input io.Reader) string {
	var sum int

	digits := []digit{
		{key: "1", value: 1},
		{key: "2", value: 2},
		{key: "3", value: 3},
		{key: "4", value: 4},
		{key: "5", value: 5},
		{key: "6", value: 6},
		{key: "7", value: 7},
		{key: "8", value: 8},
		{key: "9", value: 9},
		{key: "one", value: 1},
		{key: "two", value: 2},
		{key: "three", value: 3},
		{key: "four", value: 4},
		{key: "five", value: 5},
		{key: "six", value: 6},
		{key: "seven", value: 7},
		{key: "eight", value: 8},
		{key: "nine", value: 9},
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		nLeft := -1
		nRight := -1
		vLeft := 0
		vRight := 0

		for _, d := range digits {
			left := strings.Index(line, d.key)
			if (nLeft == -1) || (left >= 0 && left < nLeft) {
				nLeft = left
				vLeft = d.value
			}

			right := strings.LastIndex(line, d.key)
			if right >= 0 && right > nRight {
				nRight = right
				vRight = d.value
			}
		}

		sum += vLeft*10 + vRight
	}

	return fmt.Sprintf("%d", sum)
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
