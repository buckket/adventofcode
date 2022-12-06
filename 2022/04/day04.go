package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1(input io.Reader) string {
	var counter int
	var ranges [4]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		elves := strings.Split(line, ",")
		for i, elve := range elves {
			sections := strings.Split(elve, "-")
			ranges[2*i], _ = strconv.Atoi(sections[0])
			ranges[2*i+1], _ = strconv.Atoi(sections[1])
		}

		if (ranges[0] <= ranges[2] && ranges[1] >= ranges[3]) || (ranges[0] >= ranges[2] && ranges[1] <= ranges[3]) {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter)
}

func Part2(input io.Reader) string {
	var counter int
	var ranges [4]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		min := math.MaxInt
		max := 0

		elves := strings.Split(line, ",")
		for i, elve := range elves {
			sections := strings.Split(elve, "-")
			ranges[2*i], _ = strconv.Atoi(sections[0])
			ranges[2*i+1], _ = strconv.Atoi(sections[1])
			if ranges[2*i] <= min {
				min = ranges[2*i]
			}
			if ranges[2*i+1] >= max {
				max = ranges[2*i+1]
			}

		}

		if ranges[1]-ranges[0]+ranges[3]-ranges[2] >= max-min {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter)
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
