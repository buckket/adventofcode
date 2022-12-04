package main

import (
	"bufio"
	"fmt"
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

	fmt.Printf("Fully contained: %d\n", counter)
}
