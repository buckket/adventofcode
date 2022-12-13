package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

type CPU struct {
	X int

	Instructions io.Reader

	Clock   <-chan bool
	Done    chan<- bool
	Databus chan<- int
}

func (c *CPU) Exec() {
	scanner := bufio.NewScanner(c.Instructions)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		switch line[:4] {
		case "noop":
			<-c.Clock
		case "addx":
			<-c.Clock
			x, _ := strconv.Atoi(line[5:])
			<-c.Clock
			c.X += x
			c.Databus <- c.X
		default:
			break
		}
	}

	c.Done <- true
}

func Part1(input io.Reader) string {
	clock := make(chan bool)
	done := make(chan bool)
	databus := make(chan int)

	cpu := CPU{X: 1, Instructions: input, Clock: clock, Done: done, Databus: databus}
	go cpu.Exec()

	var clockCounter int
	var solution int
	var x int

	clockCounter++

	for {
		select {
		case clock <- true:
			if clockCounter == 20 || (clockCounter+20)%40 == 0 {
				solution += clockCounter * x
			}
			clockCounter++
		case x = <-databus:
		case <-done:
			return fmt.Sprintf("%d", solution)
		}
	}
}

func Part2(input io.Reader) string {
	clock := make(chan bool)
	done := make(chan bool)
	databus := make(chan int)

	cpu := CPU{X: 1, Instructions: input, Clock: clock, Done: done, Databus: databus}
	go cpu.Exec()

	var clockCounter int
	var x int

	clockCounter++

	for {
		select {
		case clock <- true:
			if abs(x-(clockCounter-1)%40) <= 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			if clockCounter%40 == 0 {
				fmt.Printf("\n")
			}
			clockCounter++
		case x = <-databus:
		case <-done:
			return ""
		}
	}
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
