package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	i := len(*s) - 1
	if i < 0 {
		return "", false
	}
	x := (*s)[i]
	*s = (*s)[:i]
	return x, true
}

type Stacks []Stack

func (s Stacks) Move(amount, old, new int) {
	tmp := make([]string, amount)
	for i := 0; i < amount; i++ {
		tmp[i], _ = s[old].Pop()
	}
	for i := amount - 1; i >= 0; i-- {
		s[new].Push(tmp[i])
	}
}

func Common(input io.Reader, single bool) string {
	var lines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	stackNumberLine := lines[len(lines)-1]
	re := regexp.MustCompile(`^(?:\d|\s)+(\d+)\s*$`)
	stackCount, _ := strconv.Atoi(re.FindStringSubmatch(stackNumberLine)[1])

	stacks := make(Stacks, stackCount)
	for i := 0; i < stackCount; i++ {
		stacks[i] = Stack{}
	}

	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j < stackCount; j++ {
			if len(lines[i]) < 4*j {
				continue
			}
			x := string(lines[i][1+4*j])
			if x != " " {
				stacks[j].Push(x)
			}
		}
	}

	re = regexp.MustCompile(`^move (?P<amount>\d+) from (?P<from>\d+) to (?P<to>\d+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		matches := re.FindStringSubmatch(line)
		amount, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])

		if single {
			stacks.Move(amount, from-1, to-1)

		} else {
			for i := 0; i < amount; i++ {
				stacks.Move(1, from-1, to-1)
			}
		}
	}

	var output string
	for _, stack := range stacks {
		x, _ := stack.Pop()
		output += x

	}

	return fmt.Sprintf("%s", output)
}

func Part1(input io.Reader) string {
	return Common(input, false)
}

func Part2(input io.Reader) string {
	return Common(input, true)
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
