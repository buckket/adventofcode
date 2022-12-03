package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cal int
	var calMax int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if cal > calMax {
				calMax = cal
			}
			cal = 0
		}
		cal += x
	}

	fmt.Printf("Most calories: %d\n", calMax)
}
