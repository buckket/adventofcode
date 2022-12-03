package main

import (
	"bufio"
	"fmt"
	"github.com/buckket/adventofcode/2021/07/helper"
	"os"
	"strconv"
	"strings"
)

func main() {
	var state []int
	var stateMax int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		stateStr := strings.Split(line, ",")
		for _, element := range stateStr {
			e, err := strconv.ParseInt(element, 10, 32)
			if err != nil {
				continue
			}
			state = append(state, int(e))
			if state[len(state)-1] > stateMax {
				stateMax = state[len(state)-1]
			}
		}
	}

	var fuel int
	pos := helper.Mean(state)
	for _, e := range state {
		n := helper.AbsInt(e - pos)
		fuel += ((n * n) + n) / 2
	}

	fmt.Printf("Result: %d\n", fuel)
}
