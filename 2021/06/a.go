package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func step(state map[int]int) map[int]int {
	newState := make(map[int]int)
	for i := 8; i >= 0; i-- {
		if i == 8 {
			newState[i] = state[0]
		} else if i == 6 {
			newState[i] = state[0] + state[i+1]
		} else {
			newState[i] = state[i+1]
		}
	}

	return newState
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	state := make(map[int]int)
	for i := 0; i < 9; i++ {
		state[i] = 0
	}

	for scanner.Scan() {
		line := scanner.Text()
		stateStr := strings.Split(line, ",")
		for _, element := range stateStr {
			e, err := strconv.ParseInt(element, 10, 32)
			if err != nil {
				continue
			}
			state[int(e)]++
		}
	}

	for i := 0; i < 80; i++ {
		state = step(state)
	}

	result := 0
	for _, v := range state {
		result += v
	}
	fmt.Printf("Result: %d\n", result)
}
