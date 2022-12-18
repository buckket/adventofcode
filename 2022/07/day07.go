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

const (
	Dir = iota
	File
)

type Node struct {
	Name     string
	Size     int
	NodeType int
	Parent   *Node
	Children map[string]*Node
}

func (n *Node) UpdateChild(name string, nodeType int, size int) {
	node, ok := n.Children[name]
	if !ok {
		node = &Node{Name: name, NodeType: nodeType, Size: size, Parent: n, Children: map[string]*Node{}}
		n.Children[name] = node
		n.Size += size
		n.UpdateParentSize(size)
	}
}

func (n *Node) GetOrCreateDir(name string) *Node {
	node, ok := n.Children[name]
	if !ok {
		node = &Node{Name: name, NodeType: Dir, Parent: n, Children: map[string]*Node{}}
		n.Children[name] = node
	}
	return node
}

func (n *Node) UpdateParentSize(size int) {
	node := n
	for node.Parent != nil {
		node.Parent.Size += size
		node = node.Parent
	}
}

func (n *Node) Print(level int) {
	fmt.Printf("%s (type=%d, size=%d)\n", strings.Repeat("\t", level)+"- "+n.Name, n.NodeType, n.Size)
	for _, c := range n.Children {
		c.Print(level + 1)
	}
}

func (n *Node) FindDirsPart1(sizes int) int {
	var csizes int
	for _, c := range n.Children {
		if c.NodeType != Dir {
			continue
		}
		csizes += c.FindDirsPart1(sizes)
	}
	sizes += csizes

	if n.NodeType == Dir && n.Size <= 100000 {
		sizes += n.Size
	}

	return sizes
}

func (n *Node) FindDirsPart2(biggerThan, currentMatch int) int {
	for _, c := range n.Children {
		if c.NodeType != Dir {
			continue
		}
		currentMatch = c.FindDirsPart2(biggerThan, currentMatch)
	}

	if n.NodeType == Dir && n.Size >= biggerThan && n.Size <= currentMatch {
		return n.Size
	}

	return currentMatch
}

type Filesystem struct {
	RootNode   *Node
	ActiveNode *Node
}

func (f *Filesystem) Init() {
	f.RootNode = &Node{Name: "/", NodeType: Dir, Children: map[string]*Node{}}
}

func (f *Filesystem) Command(lineCache []string) {
	command := strings.Split(lineCache[0], " ")
	switch command[1] {
	case "cd":
		switch command[2] {
		case "/":
			f.ActiveNode = f.RootNode
		case "..":
			if f.ActiveNode.Parent != nil {
				f.ActiveNode = f.ActiveNode.Parent
			}
		default:
			newNode := f.ActiveNode.GetOrCreateDir(command[2])
			f.ActiveNode = newNode
		}
	case "ls":
		for _, dirLine := range lineCache {
			if strings.HasPrefix(dirLine, "$") {
				continue
			}
			info := strings.Split(dirLine, " ")
			if info[0] == "dir" {
				f.ActiveNode.UpdateChild(info[1], Dir, 0)
			} else {
				size, _ := strconv.Atoi(info[0])
				f.ActiveNode.UpdateChild(info[1], File, size)
			}
		}
	}
}

func ProcessInput(input io.Reader, fs *Filesystem) {
	var lineCache []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lineCache = append(lineCache, line)
		if !strings.HasPrefix(line, "$") {
			continue
		}
		fs.Command(lineCache)
		lineCache = []string{line}
	}
	fs.Command(lineCache)
	fs.RootNode.Print(0)
}

func Part1(input io.Reader) string {
	fs := &Filesystem{}
	fs.Init()

	ProcessInput(input, fs)

	return fmt.Sprintf("%d", fs.RootNode.FindDirsPart1(0))
}

func Part2(input io.Reader) string {
	fs := &Filesystem{}
	fs.Init()

	ProcessInput(input, fs)

	biggerThan := 30000000 - (70000000 - fs.RootNode.Size)
	return fmt.Sprintf("%d", fs.RootNode.FindDirsPart2(biggerThan, fs.RootNode.Size))
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

// TODO: Could be optimized by implementing the container/heap interface
