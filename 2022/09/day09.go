package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	return -1
}

type Point struct {
	X int
	Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func (p *Point) DistanceX(other Point) int {
	return p.X - other.X
}

func (p *Point) DistanceY(other Point) int {
	return p.Y - other.Y
}

type Rope struct {
	Points  []Point
	Visited map[string]bool
}

func (r *Rope) Update(pos int, deltaX, deltaY int) {
	r.Points[pos].X += deltaX
	r.Points[pos].Y += deltaY

	if pos < len(r.Points)-1 {
		dX := r.Points[pos].DistanceX(r.Points[pos+1])
		dY := r.Points[pos].DistanceY(r.Points[pos+1])
		if abs(dX) >= 2 || abs(dY) >= 2 {
			r.Update(pos+1, sign(dX), sign(dY))
		}
	}
}

func (r *Rope) Move(dir string) {
	var dX, dY int

	switch dir {
	case "R":
		dX++
	case "L":
		dX--
	case "U":
		dY++
	case "D":
		dY--
	}

	r.Update(0, dX, dY)
	r.Visited[r.Points[len(r.Points)-1].String()] = true
}

func ProcessInput(input io.Reader, rope *Rope) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		x, _ := strconv.Atoi(command[1])
		for i := 0; i < x; i++ {
			rope.Move(command[0])
		}
	}
}

func Part1(input io.Reader) string {
	rope := &Rope{
		Points:  make([]Point, 2),
		Visited: map[string]bool{},
	}
	ProcessInput(input, rope)
	return fmt.Sprintf("%d", len(rope.Visited))
}

func Part2(input io.Reader) string {
	rope := &Rope{
		Points:  make([]Point, 10),
		Visited: map[string]bool{},
	}
	ProcessInput(input, rope)
	return fmt.Sprintf("%d", len(rope.Visited))
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
