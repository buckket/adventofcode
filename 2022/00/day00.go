package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func Part1(input io.Reader) string {
	return ""
}

func Part2(input io.Reader) string {
	return ""
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
