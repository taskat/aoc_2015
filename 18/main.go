package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type grid [][]bool

func (g grid) countOn() int {
	count := 0
	for _, row := range g {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	return count
}

func (g grid) String() string {
	lines := make([]string, 0, len(g))
	for _, row := range g {
		fields := make([]string, 0, len(row))
		for _, light := range row {
			r := '.'
			if light {
				r = '#'
			}
			fields = append(fields, string(r))
		}
		lines = append(lines, strings.Join(fields, ""))
	}
	return strings.Join(lines, "\n")
}

func CreateGrid(data []byte) grid {
	lines := strings.Split(string(data), "\r\n")
	g := make(grid, 0, len(lines))
	for _, line := range lines {
		row := make([]bool, 0, len(line))
		for _, r := range line {
			if r == '#' {
				row = append(row, true)
			}
			if r == '.' {
				row = append(row, false)
			}
		}
		g = append(g, row)
	}
	return g
}

func countOnNeighbors(g grid, i, j int) int {
	iStart := i - 1
	if iStart < 0 {
		iStart = 0
	}
	iEnd := i + 1
	if iEnd > len(g) - 1 {
		iEnd = len(g) - 1
	}
	jStart := j - 1
	if jStart < 0 {
		jStart = 0
	}
	jEnd := j + 1
	if jEnd > len(g[0]) - 1 {
		jEnd = len(g[0]) - 1
	}
	count := 0
	for iIter := iStart; iIter <= iEnd; iIter++ {
		for jIter := jStart; jIter <= jEnd; jIter++ {
			if i == iIter && j == jIter {
				continue
			}
			if g[iIter][jIter] {
				count++
			}
		}
	}
	return count
}

func next(g grid) grid {
	newGrid := make(grid, 0, len(g))
	for _, row := range g {
		newRow := make([]bool, len(row))
		newGrid = append(newGrid, newRow)
	}
	for i, row := range newGrid {
		for j := range row {
			neighbors := countOnNeighbors(g, i, j)
			if g[i][j] {
				if neighbors == 2 || neighbors == 3 {
					newGrid[i][j] = true
				} else {
					newGrid[i][j] = false
				}
			} else {
				if neighbors == 3 {
					newGrid[i][j] = true
				} else {
					newGrid[i][j] = false
				}
			}
		}
	}
	return newGrid
}

func setCornersOn(g grid) grid {
	g[0][0] = true
	g[0][len(g[0]) - 1] = true
	g[len(g) - 1][0] = true
	g[len(g) - 1][len(g[0]) - 1] = true
	return g
}

func step(g grid, steps int, cornersAlwaysOn bool) grid {
	for i := 0; i < steps; i++ {
		g = next(g)
		if cornersAlwaysOn {
			g = setCornersOn(g)
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
		grid := CreateGrid(data)
		grid = step(grid, 100, false)

		fmt.Println("The solution is: ", grid.countOn())
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		grid := CreateGrid(data)
		grid = setCornersOn(grid)
		grid = step(grid, 100, true)

		fmt.Println("The solution is: ", grid.countOn())
	}
}
