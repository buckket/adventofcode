package bingo

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Board struct {
	Points         map[int]Point
	MarkedInRow    []int
	MarkedInColumn []int
	Size           int
}

func (b *Board) Init(lines []string) {
	b.Size = len(lines)
	b.MarkedInRow = make([]int, b.Size)
	b.MarkedInColumn = make([]int, b.Size)
	b.Points = make(map[int]Point)

	for i, line := range lines {
		numbers := strings.Fields(line)
		for j, number := range numbers {
			n, err := strconv.ParseInt(number, 10, 32)
			if err != nil {
				panic("ParseInt failed")
			}
			b.Points[int(n)] = Point{X: j, Y: i}
		}
	}
}

func (b *Board) Check(number int) (result int, bingo bool) {
	if p, ok := b.Points[number]; ok {
		b.MarkedInRow[p.Y]++
		b.MarkedInColumn[p.X]++

		if b.MarkedInColumn[p.X] >= b.Size || b.MarkedInRow[p.Y] >= b.Size {
			for k := range b.Points {
				result += k
			}
			result = (result - number) * number
			return result, true
		}

		delete(b.Points, number)
	}

	return 0, false
}
