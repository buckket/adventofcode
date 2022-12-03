package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getValue(code []string, pos int) int {
	if pos > len(code) {
		panic("Invalid pos")
	}
	v, err := strconv.Atoi(code[pos])
	if err != nil {
		panic("Atoi failed")
	}
	return v
}

func writeValue(code []string, pos, value int) {
	if pos > len(code) {
		panic("Invalid pos")
	}
	code[pos] = strconv.Itoa(value)
}

func execute(code []string) []string {
	pos := 0
	max := len(code)

	for pos < max && code[pos] != "99" {
		if code[pos] == "1" {
			a := getValue(code, getValue(code, pos+1))
			b := getValue(code, getValue(code, pos+2))
			t := getValue(code, pos+3)
			writeValue(code, t, a+b)
		} else if code[pos] == "2" {
			a := getValue(code, getValue(code, pos+1))
			b := getValue(code, getValue(code, pos+2))
			t := getValue(code, pos+3)
			writeValue(code, t, a*b)
		} else {
			panic("Unknown opcode")
		}
		pos += 4
	}
	return code
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		intcode := strings.Split(scanner.Text(), ",")
		intcode[1] = "12"
		intcode[2] = "2"
		fmt.Printf("%s\n", strings.Join(execute(intcode), ","))
	}
}
