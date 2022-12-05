package main

import (
	"bufio"
	"fmt"
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

func (s Stacks) Move(old, new int) {
	x, _ := s[old].Pop()
	s[new].Push(x)
}

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
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

		for i := 0; i < amount; i++ {
			stacks.Move(from-1, to-1)
		}

	}

	var output string
	for _, stack := range stacks {
		x, _ := stack.Pop()
		output += x

	}
	fmt.Printf("%s\n", output)
}
