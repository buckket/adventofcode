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

		var noun, verb int

	out:
		for noun = 0; noun < 99; noun++ {
			for verb = 0; verb < 99; verb++ {
				tmp := make([]string, len(intcode))
				copy(tmp, intcode)
				tmp[1] = strconv.Itoa(noun)
				tmp[2] = strconv.Itoa(verb)
				res := execute(tmp)
				if res[0] == "19690720" {
					break out
				}
			}
		}
		fmt.Printf("noun: %d, verb: %d, res: %d\n", noun, verb, 100*noun+verb)
	}
}
