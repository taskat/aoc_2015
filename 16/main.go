package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Sue map[string]int

func CreateSues(data []byte) []Sue {
	lines := strings.Split(string(data), "\r\n")
	sues := make([]Sue, 0, len(lines))
	for _, line := range lines {
		words := strings.Split(line, " ")
		line = strings.Join(words[2:], " ")
		properties := strings.Split(line, ", ")
		s := make(Sue)
		for _, proproperty := range properties {
			parts := strings.Split(proproperty, ": ")
			num, _ := strconv.Atoi(parts[1])
			s[parts[0]] = num
		}
		sues = append(sues, s)
	}
	return sues
}

func CreateGoal() Sue {
	sue := make(Sue)
	sue["children"] = 3
	sue["cats"] = 7
	sue["samoyeds"] = 2
	sue["pomeranians"] = 3
	sue["akitas"] = 0
	sue["vizslas"] = 0
	sue["goldfish"] = 5
	sue["trees"] = 3
	sue["cars"] = 2
	sue["perfumes"] = 1
	return sue
}

func match1(sue, goal Sue) bool {
	for key, value := range sue {
		if value != goal[key] {
			return false
		}
	}
	return true
}

func match2(sue, goal Sue) bool {
	for key, value := range sue {
		switch key {
		case "cats", "trees" :
			if goal[key] >= value {
				return false
			}
		case "pomeranians", "goldfish":
			if goal[key] <= value {
				return false
			}
		default:
			if value != goal[key] {
				return false
			}
		}
	}
	return true
}

func find(sues []Sue, goal Sue, match func(Sue, Sue) bool) int {
	for i, sue := range sues {
		if match(sue, goal) {
			return i
		}
	}
	return -1
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
		sues := CreateSues(data)
		goal := CreateGoal()
		index := find(sues, goal, match1)

		fmt.Println("The solution is: ", index+1)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		sues := CreateSues(data)
		goal := CreateGoal()
		index := find(sues, goal, match2)

		fmt.Println("The solution is: ", index + 1)
	}
}
