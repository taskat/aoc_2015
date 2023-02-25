package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taskat/golang-utility/array"
)

var eggnog = 150

func CreateContainers(data []byte) []int {
	lines := strings.Split(string(data), "\r\n")
	containers := make([]int, 0, len(lines))
	for _, line := range lines {
		container, _ := strconv.Atoi(line)
		containers = append(containers, container)
	}
	return containers
}

func getCombinations(containers []int) [][]int {
	combinations := make([][]int, 0)
	var findSolutions func(containers []int, goal int, used []int)
	findSolutions = func(containers []int, goal int, used []int) {
		if goal < 0 {
			return
		}
		if goal == 0 {
			combinations = append(combinations, make([]int, len(used)))
			copy(combinations[len(combinations) - 1], used)
			return
		}
		for i, container := range containers {
			findSolutions(containers[i + 1:],
				goal - container, array.AppendInt(used, container))
		}
	}
	findSolutions(containers, eggnog, nil)
	return combinations
}

func findMinContainers(combinations [][]int) int {
	min := len(combinations[0])
	for _, combination := range combinations {
		if len(combination) < min {
			min = len(combination)
		}
	}
	return min
}

func getMinimalCombinations(combinations [][]int) [][]int {
	min := findMinContainers(combinations)
	result := make([][]int, 0)
	for _, combination := range combinations {
		if len(combination) == min {
			result = append(result, combination)
		}
	}
	return result
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
		containers := CreateContainers(data)
		combinations := getCombinations(containers)
		
		fmt.Println("The solution is: ", len(combinations))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		containers := CreateContainers(data)
		combinations := getCombinations(containers)
		minimal := getMinimalCombinations(combinations)

		fmt.Println("The solution is: ", len(minimal))
	}
}
