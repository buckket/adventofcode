package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var data []int
	var bitSize = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if bitSize == 0 {
			bitSize = len(line)
		}
		x, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, int(x))
	}

	oxygen := calculate(data, bitSize, 1)
	scrubber := calculate(data, bitSize, 0)

	fmt.Printf("Oxygen: %d, CO2 Scrubber: %d, Result: %d\n", oxygen, scrubber, oxygen*scrubber)
}

func calculate(data []int, bitSize int, criteria int) int {
	for i := bitSize - 1; i >= 0; i-- {
		var startingZero []int
		var startingOne []int
		for j := 0; j < len(data); j++ {
			if data[j]&(1<<i) > 0 {
				startingOne = append(startingOne, data[j])
			} else {
				startingZero = append(startingZero, data[j])
			}
		}

		switch criteria {
		case 1:
			if len(startingOne) >= len(startingZero) {
				data = startingOne
			} else {
				data = startingZero
			}
		case 0:
			if len(startingZero) <= len(startingOne) {
				data = startingZero
			} else {
				data = startingOne
			}
		default:
			panic("Unknown criteria")
		}

		if len(data) == 1 {
			return data[0]
		}

	}

	return -1
}
