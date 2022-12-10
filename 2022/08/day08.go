package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

type Row struct {
	Max int32

	Visible map[int]bool
	SizeMap map[int32]int

	Pos int
}

func (r *Row) Init() {
	r.Visible = make(map[int]bool)
	r.SizeMap = make(map[int32]int)
}

func (r *Row) Update(next int32) {
	if next > r.Max {
		// Bigger tree, we definitely can see it from the left
		r.Max = next
		r.Visible[r.Pos] = true
	}

	if next >= r.Max {
		// Reset, as we definitely can not see everything before this one from the right
		r.SizeMap = make(map[int32]int)
		r.SizeMap[next] = r.Pos
	} else {
		// Could be visible from the right, update SizeMap
		r.SizeMap[next] = r.Pos
	}

	r.Pos++
}

func (r *Row) Finish(i, j int) []int {
	// Sort all tree sizes we have seen in ascending order
	var keys []int
	for k, _ := range r.SizeMap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	// Iterate over all tree sizes
	// If a smaller tree is further to the right then the previous bigger ones we can see it
	var pos int
	for l := len(keys) - 1; l >= 0; l-- {
		if r.SizeMap[int32(keys[l])] > pos {
			r.Visible[r.SizeMap[int32(keys[l])]] = true
			pos = r.SizeMap[int32(keys[l])]
		}
	}

	// Create a list of all visible tree positions
	var result []int
	for k, _ := range r.Visible {
		result = append(result, i*k+j)
	}

	return result
}

func Part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var rows []Row
	var columns []Row

	var firstRow = true
	var lineCounter int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		rows = append(rows, Row{})
		rows[lineCounter].Init()

		for i, x := range line {
			rows[lineCounter].Update(x)
			if firstRow {
				columns = append(columns, Row{})
				columns[i].Init()
			}
			columns[i].Update(x)
		}
		lineCounter++
		firstRow = false
	}

	result := make(map[int]bool)
	for i, x := range rows {
		visible := x.Finish(1, len(rows)*i)
		for _, y := range visible {
			result[y] = true
		}
	}
	for i, x := range columns {
		visible := x.Finish(len(columns), i)
		for _, y := range visible {
			result[y] = true
		}
	}

	return fmt.Sprintf("%d", len(result))
}

func Part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var initialized bool
	var size int
	var forest []int

	var lineCount = 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		if !initialized {
			size = len(line)
			forest = make([]int, size*size)
			initialized = true
		}

		for i, v := range line {
			forest[lineCount*size+i] = int(v)
		}

		lineCount++
	}

	var score int
	for i := 0; i < len(forest); i++ {
		y := i / size
		x := i - y*size

		// Skip computation when on the edges
		if x == 0 || y == 0 || x == size-1 || y == size-1 {
			continue
		}

		var visibleL int
		for xL := x; xL > 0; xL-- {
			visibleL++
			if forest[i-(x-xL+1)] >= forest[i] {
				break
			}
		}

		var visibleR int
		for xR := x; xR < size-1; xR++ {
			visibleR++
			if forest[i+(xR-x+1)] >= forest[i] {
				break
			}
		}

		var visibleT int
		for yT := y; yT > 0; yT-- {
			visibleT++
			if forest[i-(y-yT+1)*size] >= forest[i] {
				break
			}
		}

		var visibleD int
		for yD := y; yD < size-1; yD++ {
			visibleD++
			if forest[i+(yD-y+1)*size] >= forest[i] {
				break
			}
		}

		if visibleL*visibleR*visibleT*visibleD > score {
			score = visibleL * visibleR * visibleT * visibleD
		}
	}

	return fmt.Sprintf("%d", score)
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
