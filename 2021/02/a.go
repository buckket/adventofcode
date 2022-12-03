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
		case "down":
			y += t
		case "up":
			y -= t

		}
	}

	fmt.Printf("Result: %d\n", x*y)
}
