package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var counter int
	var ranges [4]int

	scanner := bufio.NewScanner(os.Stdin)
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

	fmt.Printf("Overlap: %d\n", counter)
}
