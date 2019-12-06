package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func createVertices(instructions string) (ver []Vertex, offsetX, offsetY, dimX, dimY int) {
	var a, b Vertex
	var minX, minY, maxX, maxY int

	ver = append(ver, Vertex{X: 0, Y: 0,})

	instr := strings.Split(instructions, ",")
	for _, vec := range instr {
		l, err := strconv.Atoi(vec[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch vec[0] {
		case 'U':
			b = Vertex{X: a.X, Y: a.Y + l,}
		case 'D':
			b = Vertex{X: a.X, Y: a.Y - l,}
		case 'R':
			b = Vertex{X: a.X + l, Y: a.Y,}
		case 'L':
			b = Vertex{X: a.X - l, Y: a.Y,}
		}

		if b.X < minX {
			minX = b.X
		}
		if b.X > maxX {
			maxX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
		if b.Y > maxY {
			maxY = b.Y
		}

		ver = append(ver, b)
		a = b
	}

	if minX < 0 {
		offsetX = -minX
	}
	if minY < 0 {
		offsetY = -minY
	}

	maxX++
	maxY++

	return ver, offsetX, offsetY, maxX, maxY
}

func checkIntersection(grid [][]int, dist, x, y, oX, oY, id int) int {
	if grid[x][y] > 0 && grid[x][y] != id {
		t := int(math.Abs(float64(x-oX)) + math.Abs(float64(y-oY)))
		if t < dist && t > 0 {
			dist = t
		}
	}
	return dist
}

func drawLines(grid [][]int, ver []Vertex, oX, oY, id int) (dist int) {
	dist = len(grid) + len(grid[0])
	for i := 0; i < len(ver)-1; i++ {
		if d := ver[i+1].X - ver[i].X; d != 0 {
			sign := int(math.Copysign(1, float64(d)))
			for j := 0; j < int(math.Abs(float64(d))); j++ {
				dist = checkIntersection(grid, dist, oX+ver[i].X+sign*j, oY+ver[i].Y, oX, oY, id)
				grid[oX+ver[i].X+sign*j][oY+ver[i].Y] = id
			}
		}
		if d := ver[i+1].Y - ver[i].Y; d != 0 {
			sign := int(math.Copysign(1, float64(d)))
			for j := 0; j < int(math.Abs(float64(d))); j++ {
				dist = checkIntersection(grid, dist, oX+ver[i].X, oY+ver[i].Y+sign*j, oX, oY, id)
				grid[oX+ver[i].X][oY+ver[i].Y+sign*j] = id
			}
		}
	}
	dist = checkIntersection(grid, dist, oX+ver[len(ver)-1].X, oY+ver[len(ver)-1].Y, oX, oY, id)
	grid[oX+ver[len(ver)-1].X][oY+ver[len(ver)-1].Y] = id
	return dist
}

func greaterValue(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	wire1, _ := ioutil.ReadFile("wire1")
	wire2, _ := ioutil.ReadFile("wire2")
	a, oX1, oY1, dX1, dY1 := createVertices(strings.TrimSuffix(string(wire1), "\n"))
	b, oX2, oY2, dX2, dY2 := createVertices(strings.TrimSuffix(string(wire2), "\n"))

	var oX, oY, dX, dY int
	oX = greaterValue(oX1, oX2)
	oY = greaterValue(oY1, oY2)
	dX = greaterValue(dX1, dX2)
	dY = greaterValue(dY1, dY2)
	dX = dX + oX
	dY = dY + oY

	grid := make([][]int, dX)
	for i := 0; i < dX; i++ {
		grid[i] = make([]int, dY)
	}

	drawLines(grid, a, oX, oY,1)
	dist := drawLines(grid, b, oX, oY, 2)
	fmt.Printf("Closest intersection: %d\n", dist)

	for y := dY - 1; y >= 0; y-- {
		for x := 0; x < dX; x++ {
			fmt.Printf("%d", grid[x][y])
		}
		fmt.Print("\n")
	}

}
