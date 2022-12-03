package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var prio int

	scanner := bufio.NewScanner(os.Stdin)
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

	fmt.Printf("Prio sum: %d\n", prio)
}
