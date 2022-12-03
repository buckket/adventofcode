package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x = 0
	var y = 0
	var aim = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		t, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}

		switch s[0] {
		case "forward":
			x += t
			y += aim * t
		case "down":
			aim += t
		case "up":
			aim -= t

		}
	}

	fmt.Printf("Result: %d\n", x*y)
}
