package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	Start Point
	End   Point
}

type Point struct {
	X int
	Y int
}

type Grid struct {
	Lines []Line
	SizeX int
	SizeY int
	Data  [][]int
}

func (g *Grid) Init(lines []Line) {
	g.Lines = lines
	for _, line := range g.Lines {
		if line.Start.X > g.SizeX {
			g.SizeX = line.Start.X
		}
		if line.Start.Y > g.SizeY {
			g.SizeY = line.Start.Y
		}
		if line.End.X > g.SizeX {
			g.SizeX = line.Start.X
		}
		if line.End.X > g.SizeY {
			g.SizeX = line.Start.X
		}
	}

	g.Data = make([][]int, g.SizeX)
	for i := range g.Data {
		g.Data[i] = make([]int, g.SizeY)
	}
}

func (g *Grid) DrawHV() {
	for _, line := range g.Lines {
		if line.Start.X == line.End.X {
			diff := lyy
			g.Data[line.Start.X][i]++
		}
	}
}

func NewPoint(point string) Point {
	comp := strings.Split(strings.TrimSpace(point), ",")

	x, err := strconv.ParseInt(comp[0], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	y, err := strconv.ParseInt(comp[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	return Point{X: int(x), Y: int(y)}
}

func NewLine(line string) Line {
	data := strings.Split(line, "->")
	return Line{Start: NewPoint(data[0]), End: NewPoint(data[1])}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []Line

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, NewLine(line))
	}

	grid := Grid{}
	grid.Init(lines)

	fmt.Printf("%#v", grid)
}
