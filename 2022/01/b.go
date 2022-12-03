package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var cal int
	var calMax [3]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if cal > calMax[0] {
				calMax[0] = cal
				sort.Ints(calMax[:])
			}
			cal = 0
		}
		cal += x
	}

	fmt.Printf("Most calories (top three): %d\n", calMax[0]+calMax[1]+calMax[2])
}
