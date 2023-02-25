package main

import (
	"fmt"
	"io/ioutil"
)

type Pos struct {
	x, y int
}

func createGrid(data []byte) map[Pos]int {
	grid := make(map[Pos]int)
	current := Pos{x: 0, y: 0}
	grid[current]++
	for _, next := range string(data) {
		switch next {
		case '<': current.x--
		case '>': current.x++
		case 'v': current.y--
		case '^': current.y++
		}
		grid[current]++
	}
	return grid
}

func move(dir rune, current Pos) Pos {
	switch dir {
		case '<': current.x--
		case '>': current.x++
		case 'v': current.y--
		case '^': current.y++
		}
	return current
}

func createDoubleGrid(data []byte) map[Pos]int {
	grid := make(map[Pos]int)
	santaCurrent := Pos{x: 0, y: 0}
	roboCurrent := Pos{x: 0, y: 0}
	grid[santaCurrent]++
	grid[roboCurrent]++
	for i, next := range string(data) {
		if i % 2 == 0 {
			santaCurrent = move(next, santaCurrent)
			grid[santaCurrent]++
		} else {
			roboCurrent = move(next, roboCurrent)
			grid[roboCurrent]++
		}
	}
	return grid
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
		grid := createGrid(data)

		fmt.Println("The solution is: ", len(grid))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		grid := createDoubleGrid(data)

		fmt.Println("The solution is: ", len(grid))
	}
}
