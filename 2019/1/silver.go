package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var fuel = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel += int(math.Floor(float64(mass)/3) - 2)
	}
	fmt.Printf("Total fuel: %d\n", fuel)
}
