package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calcFuel(fuel, mass int) int {
	needed := int(math.Floor(float64(mass)/3) - 2)
	if needed <= 0 {
		return fuel
	} else {
		return calcFuel(fuel+needed, needed)
	}
}

func main() {
	var fuel = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel += calcFuel(0, mass)
	}
	fmt.Printf("Total fuel: %d\n", fuel)
}
