package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var depths []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, x)
	}

	var c = 0
	for i := 0; i < len(depths)-3; i++ {
		if (depths[i] + depths[i+1] + depths[i+2]) < (depths[i+1] + depths[i+2] + depths[i+3]) {
			c++
		}
	}

	fmt.Printf("Larger than previous: %d\n", c)
}
