package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var l, c int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if x > l && l > 0 {
			c++
		}
		l = x
	}
	fmt.Printf("Larger than previous: %d\n", c)
}
