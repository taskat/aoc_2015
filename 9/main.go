package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taskat/golang-utility/unique"
)

type Path struct {
	cities []string
	dist int
}

func (d Path) other(s string) string {
	if d.cities[0] == s {
		return d.cities[1]
	}
	return d.cities[0]
}

func CreatePaths(data []byte) []Path {
	lines := strings.Split(string(data), "\r\n")
	paths := make([]Path, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		cities := strings.Split(parts[0], " to ")
		dist, _ := strconv.Atoi(parts[1])
		paths = append(paths, Path{cities: cities, dist: dist})
	}
	return paths
}

func contains(arr []string, s string) bool {
	for _, element := range arr {
		if element == s {
			return true
		}
	}
	return false
}

type city string 

func (c city) Equals(other unique.Item) bool {
	otherCity, ok := other.(city)
	return ok && c == otherCity
}

func getCities(paths []Path) unique.UniqueArray {
	arr := unique.Create(nil)
	for _, path := range paths {
		arr.Push(city(path.cities[0]))
		arr.Push(city(path.cities[1]))
	}
	return arr
}

func getAdjacent(paths []Path, from string) []Path {
	adjacents := make([]Path, 0)
	for _, path := range paths {
		if contains(path.cities, from) {
			adjacents = append(adjacents, path)
		}
	}
	return adjacents
}

func remove(paths []Path, from string) []Path {
	pathsCopy := make([]Path, len(paths))
	copy(pathsCopy, paths)
	for i := 0; i < len(pathsCopy); i++ {
		if contains(pathsCopy[i].cities, from) {
			pathsCopy = append(pathsCopy[:i], pathsCopy[i + 1:]...)
			i--
		}
	}
	return pathsCopy
}

func getRoute(paths []Path, from string, current int, found []int) []int {
	if len(paths) == 0 {
		return append(found, current)
	}
	adjacents := getAdjacent(paths, from)
	removed := remove(paths, from)
	for _, adjacent := range adjacents {
		found = getRoute(removed, adjacent.other(from), current + adjacent.dist, found)
	}
	return found
}

func getRoutes(paths []Path) unique.UniqueIntArray {
	cities := getCities(paths)
	routes := unique.CreateInts(nil)
	for i := 0; i < cities.Len(); i++ {
		startCity := string(cities.Get(i).(city))
		newRoutes := getRoute(paths, startCity, 0, nil)
		for _, route := range newRoutes {
			routes.Push(route)
		}
	}
	return routes
}

func getMin(routes unique.UniqueIntArray) int {
	min := routes.Get(0)
	for i := 1; i < routes.Len(); i++ {
		value := routes.Get(i)
		if value < min {
			min = value
		}
	}
	return min
}

func getMax(routes unique.UniqueIntArray) int {
	max := routes.Get(0)
	for i := 1; i < routes.Len(); i++ {
		value := routes.Get(i)
		if value > max {
			max = value
		}
	}
	return max
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
		paths := CreatePaths(data)
		routes := getRoutes(paths)
		min := getMin(routes)
		
		fmt.Println("The solution is: ", min)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		paths := CreatePaths(data)
		routes := getRoutes(paths)
		max := getMax(routes)

		fmt.Println("The solution is: ", max)
	}
}
