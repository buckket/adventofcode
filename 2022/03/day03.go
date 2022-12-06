package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func Part1(input io.Reader) string {
	var prio int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rucksack := scanner.Text()

		rucksack_len := len(rucksack)
		if rucksack_len == 0 {
			break
		}

		items := make(map[int32]int)
		for i, code := range rucksack {

			if code >= 97 {
				code -= 97
			} else {
				code -= 65 - 26
			}

			if _, ok := items[code]; !ok {
				items[code] = 0
			}

			if i > rucksack_len/2-1 {
				items[code] |= 1
			} else {
				items[code] |= 2
			}

			if items[code] == 3 {
				prio += int(code) + 1
				break
			}
		}
	}

	return fmt.Sprintf("%d", prio)
}

func Part2(input io.Reader) string {
	var prio int
	var counter int

	items := make(map[int32]int)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if counter > 2 {
			counter = 0
			items = make(map[int32]int)
		}

		rucksack := scanner.Text()
		if len(rucksack) == 0 {
			break
		}

		for _, code := range rucksack {
			if code >= 97 {
				code -= 97
			} else {
				code -= 65 - 26
			}

			if _, ok := items[code]; !ok {
				items[code] = 0
			}
			items[code] |= 1 << counter

			if items[code] == 7 {
				prio += int(code) + 1
				break
			}
		}
		counter++
	}

	return fmt.Sprintf("%d", prio)
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
