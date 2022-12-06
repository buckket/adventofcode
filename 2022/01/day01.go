package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func Part1(input io.Reader) string {
	var cal int
	var calMax int

	scanner := bufio.NewScanner(input)
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
	if cal > calMax {
		calMax = cal
	}
	cal = 0

	return fmt.Sprintf("%d", calMax)
}

func Part2(input io.Reader) string {
	var cal int
	var calMax [3]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil{
			if cal > calMax[0] {
				calMax[0] = cal
				sort.Ints(calMax[:])
			}
			cal = 0
		}
		cal += x
	}
	if cal > calMax[0] {
		calMax[0] = cal
		sort.Ints(calMax[:])
	}

	return fmt.Sprintf("%d", calMax[0]+calMax[1]+calMax[2])
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
