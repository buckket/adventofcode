package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var data []int
	var bitSize = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if bitSize == 0 {
			bitSize = len(line)
		}
		x, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, int(x))
	}

	var out = 0
	for i := 0; i < bitSize; i++ {
		var t = 0
		for j := 0; j < len(data); j++ {
			if data[j]&(1<<i) > 0 {
				t++
			}
		}
		if t > len(data)/2 {
			out |= 1 << i
		}
	}

	fmt.Printf("Result: %d\n", out*(^out&((1<<bitSize)-1)))
}
