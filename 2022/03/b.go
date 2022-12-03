package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var prio int
	var counter int

	items := make(map[int32]int)

	scanner := bufio.NewScanner(os.Stdin)
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

	fmt.Printf("Prio sum: %d\n", prio)
}
