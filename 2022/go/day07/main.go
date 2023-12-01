package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	maxDirSize      = 100_000
	totalDiskSpace  = 70_000_000
	updateDiskSpace = 30_000_000
)

var (
	cdReg   = regexp.MustCompile(`^\$ cd (/|..|[A-Za-z_]+)$`)
	lsReg   = regexp.MustCompile(`^\$ ls$`)
	dirReg  = regexp.MustCompile(`^dir (\w+)$`)
	fileReg = regexp.MustCompile(`^(\d+) ([0-9A-Za-z_\.]+)$`)
)

func main() {
	f, err := os.Open("../data/day07")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var current *Node

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if matches := cdReg.FindStringSubmatch(line); len(matches) == 2 {
			dirName := matches[1]

			if current == nil { // root
				current = NewNode(dirName, nil, nil)
			} else if dirName == ".." {
				current = current.parent
			} else {
				var ok bool
				current, ok = current.children[dirName]
				if !ok {
					log.Fatalf("unable to find child %v in %v", dirName, current.name)
				}
			}
		} else if lsReg.MatchString(line) {
			// nothing to do
		} else if matches := dirReg.FindStringSubmatch(line); len(matches) == 2 {
			dirName := matches[1]

			fmt.Printf("adding directory %v to %v", dirName, current.name)
			current.AddChild(dirName, nil)
		} else if matches := fileReg.FindStringSubmatch(line); len(matches) == 3 {
			fileName := matches[2]
			size, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("invalid file size: %v for file: %v", matches[1], fileName)
			}

			current.AddChild(fileName, &size)
		}
	}

	root := current.Root()

	dirs := root.Collect(func(n *Node) bool {
		return n.size == nil && n.SizeOf() <= maxDirSize
	})

	var totalSize int
	for _, d := range dirs {
		totalSize += d.SizeOf()
		fmt.Printf("%v -> %d\n", d.name, d.SizeOf())
	}
	fmt.Printf("sum of directories having size up to %v: %v\n", maxDirSize, totalSize)

	usedDiskSpace := root.SizeOf()
	fmt.Println("total used disk space", usedDiskSpace)

	availableDiskSpace := totalDiskSpace - usedDiskSpace
	fmt.Println("available disk space", availableDiskSpace)

	candidates := root.Collect(func(n *Node) bool {
		return n.size == nil && n.SizeOf() > updateDiskSpace-availableDiskSpace
	})

	var min = math.MaxInt
	for _, d := range candidates {
		size := d.SizeOf()
		if size < min {
			min = size
		}
		fmt.Printf("%v -> %d\n", d.name, size)
	}

	fmt.Println("size of the smallest directory that can be deleted", min)
}

type Node struct {
	parent *Node
	name   string
	size   *int

	children map[string]*Node
}

func NewNode(name string, parent *Node, size *int) *Node {
	return &Node{
		parent:   parent,
		name:     name,
		size:     size,
		children: make(map[string]*Node),
	}
}

func (n *Node) Root() *Node {
	if n.parent == nil {
		return n
	}

	return n.parent.Root()
}

func (n *Node) SizeOf() int {
	var total int

	if n.size != nil {
		total += *n.size
	}

	for _, c := range n.children {
		total += c.SizeOf()
	}

	return total
}

func (n *Node) AddChild(name string, size *int) {
	n.children[name] = NewNode(name, n, size)
}

func (n *Node) Collect(predicate func(*Node) bool) []*Node {
	var acc []*Node

	if predicate(n) {
		acc = append(acc, n)
	}

	for _, child := range n.children {
		acc = append(acc, child.Collect(predicate)...)
	}

	return acc
}
