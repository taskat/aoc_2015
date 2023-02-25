package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taskat/golang-utility/intmath"
)

type pos struct {
	x, y int
}

type grid [][]int

func (g grid) turn(from, to pos, f modifier) {
	for y := from.y; y <= to.y; y++ {
		for x := from.x; x <= to.x; x++ {
			g[y][x] = f(g[y][x])
		}
	}
}

func (g grid) sum() int {
	sum := 0
	for _, row := range g {
		sum += intmath.Sum(row)
	}
	return sum
}

func CreateGrid() grid {
	grid := make(grid, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	return grid
}

func posFromString(s string) pos {
	coords := strings.Split(s, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return pos{x: x, y: y}
}

func getPos(line string) (pos, pos) {
	parts := strings.Split(line, " through ")
	firsts := strings.Split(parts[0], " ")
	startString := firsts[len(firsts) - 1]
	start := posFromString(startString)
	end := posFromString(parts[1])
	return start, end
}

type modifier func(int) int

func on1(int) int {
	return 1
}

func off1(int) int {
	return 0
}

func toggle1(n int) int {
	return 1 - n
}

func execute1(lines []string) grid {
	g := CreateGrid()
	for _, line := range lines {
		start, end := getPos(line)
		switch {
		case strings.Contains(line, "turn on"):
			g.turn(start, end, on1)
		case strings.Contains(line, "turn off"):
			g.turn(start, end, off1)
		case strings.Contains(line, "toggle"):
			g.turn(start, end, toggle1)
		}
	}
	return g
}

func on2(n int) int {
	return n + 1
}

func off2(n int) int {
	n--
	if n < 0 {
		return 0
	}
	return n
}

func toggle2(n int) int {
	return n + 2
}

func execute2(lines []string) grid {
	g := CreateGrid()
	for _, line := range lines {
		start, end := getPos(line)
		switch {
		case strings.Contains(line, "turn on"):
			g.turn(start, end, on2)
		case strings.Contains(line, "turn off"):
			g.turn(start, end, off2)
		case strings.Contains(line, "toggle"):
			g.turn(start, end, toggle2)
		}
	}
	return g
}

func main() {
	part := 0
	validAnswer := false
	for !validAnswer {
		fmt.Println("Which part? (1 or 2)")
		fmt.Scanf("%d\n", &part)
		if part < 1 || part > 2 {
			fmt.Println("Invalid answer!")
			continue
		}
		validAnswer = true
	}
	switch part {
	case 1:
		fmt.Println("Solving part 1")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(data), "\r\n")
		g := execute1(lines)
		on := g.sum()

		fmt.Println("The solution is: ", on)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(data), "\r\n")
		g := execute2(lines)
		brightness := g.sum()

		fmt.Println("The solution is: ", brightness)
	}
}
