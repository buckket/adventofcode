package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func countUnique(str string) int {
	var unique int
	keys := make(map[int32]bool)
	for _, entry := range str {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique++
		}
	}
	return unique
}

func findStart(input io.Reader, distinctChars int) string {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		signal := scanner.Text()
		if len(signal) == 0 {
			break
		}

		// Warning: This code does not work with unicode runes
		// We would have to convert the signal to a []rune first

		// Solution 1
		// O(distinctChars * len(signal))
		for i := 0; i < len(signal)-distinctChars-1; i++ {
			if countUnique(signal[i:i+distinctChars]) >= distinctChars {
				//return fmt.Sprintf("%d", i+distinctChars)
			}
		}

		// Solution 2
		// O(len(signal)
		var unique int
		keys := make(map[uint8]int)

		for i := 0; i < len(signal); i++ {
			if keys[signal[i]]++; keys[signal[i]] == 1 {
				unique++
			}
			if i >= distinctChars {
				if keys[signal[i-distinctChars]]--; keys[signal[i-distinctChars]] == 0 {
					unique--
				}
			}
			if unique >= distinctChars {
				return fmt.Sprintf("%d", i+1)
			}
		}

	}

	return ""
}

func Part1(input io.Reader) string {
	return findStart(input, 4)
}

func Part2(input io.Reader) string {
	return findStart(input, 14)
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
