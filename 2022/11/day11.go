package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GeneratorOp(op string) func(x, y int) int {
	switch op {
	case "+":
		return func(x, y int) int {
			return x + y
		}
	case "*":
		return func(x, y int) int {
			return x * y
		}
	}
	return func(x, y int) int {
		return 0
	}
}

type Monkeys []Monkey

func (m Monkeys) Len() int           { return len(m) }
func (m Monkeys) Less(i, j int) bool { return m[i].Counter < m[j].Counter }
func (m Monkeys) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type Monkey struct {
	Items     []int
	Operation func(x int) int
	Test      func(x int) int
	Counter   int
	Divisor   int
}

func (m *Monkey) Throw(worryDiv int) (dest, item int, ok bool) {
	if len(m.Items) == 0 {
		return 0, 0, false
	}
	m.Counter++

	item = m.Items[0]
	if len(m.Items) > 1 {
		m.Items = m.Items[1:]
	} else {
		m.Items = nil
	}

	item = m.Operation(item) / worryDiv
	dest = m.Test(item)
	return dest, item, true
}

func ParseMonkey(buffer []string) Monkey {
	m := Monkey{}

	itemsStr := strings.Split(buffer[1][18:], ", ")
	for _, item := range itemsStr {
		i, _ := strconv.Atoi(item)
		m.Items = append(m.Items, i)
	}

	opStr := strings.Split(buffer[2][19:], " ")
	if opStr[0] == "old" && opStr[2] == "old" {
		m.Operation = func(x int) int {
			return GeneratorOp(opStr[1])(x, x)
		}
	} else {
		c, _ := strconv.Atoi(opStr[2])
		m.Operation = func(x int) int {
			return GeneratorOp(opStr[1])(x, c)
		}
	}

	divBy, _ := strconv.Atoi(buffer[3][21:])
	trueM, _ := strconv.Atoi(buffer[4][29:])
	falseM, _ := strconv.Atoi(buffer[5][30:])
	m.Test = func(x int) int {
		if x%divBy == 0 {
			return trueM
		}
		return falseM
	}
	m.Divisor = divBy

	return m
}

func ProcessInput(input io.Reader) Monkeys {
	var monkeys Monkeys
	var monkeyBuffer []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			monkeyBuffer = append(monkeyBuffer, line)
			continue
		}
		monkeys = append(monkeys, ParseMonkey(monkeyBuffer))
		monkeyBuffer = nil
	}
	monkeys = append(monkeys, ParseMonkey(monkeyBuffer))

	return monkeys
}

func Part1(input io.Reader) string {
	monkeys := ProcessInput(input)

	for i := 0; i < 20; i++ {
		for j, _ := range monkeys {
			for true {
				dest, item, ok := monkeys[j].Throw(3)
				if !ok {
					break
				}
				monkeys[dest].Items = append(monkeys[dest].Items, item)
			}
		}
	}

	sort.Sort(monkeys)
	return fmt.Sprintf("%d", monkeys[len(monkeys)-1].Counter*monkeys[len(monkeys)-2].Counter)
}

func Part2(input io.Reader) string {
	monkeys := ProcessInput(input)

	// TODO: Refresh memory on the following topics:
	//  - https://en.wikipedia.org/wiki/Coprime_integers
	//  - https://en.wikipedia.org/wiki/Chinese_remainder_theorem
	//  - https://de.wikipedia.org/wiki/Restklassenring

	lcm := 1
	for j, _ := range monkeys {
		lcm *= monkeys[j].Divisor
	}

	for i := 0; i < 10000; i++ {
		for j, _ := range monkeys {
			for true {
				dest, item, ok := monkeys[j].Throw(1)
				item = item % lcm
				if !ok {
					break
				}
				monkeys[dest].Items = append(monkeys[dest].Items, item)
			}
		}
	}

	sort.Sort(monkeys)
	return fmt.Sprintf("%d", monkeys[len(monkeys)-1].Counter*monkeys[len(monkeys)-2].Counter)
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
