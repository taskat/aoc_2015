package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func floor(s string) int {
	up := strings.Count(s, "(")
	down := strings.Count(s, ")")
	return up - down
}

func firstBasement(s string) int {
	currentFloor := 0
	for i, r := range s {
		switch r {
		case '(': currentFloor++
		case ')': currentFloor--
		}
		if currentFloor < 0 {
			return i + 1
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
		result := floor(string(data))

		fmt.Println("The solution is: ", result)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		result := firstBasement(string(data))

		fmt.Println("The solution is: ", result)
	}
}
