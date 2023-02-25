package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func getCoord(data []byte) (int, int) {
	r, _ := regexp.Compile(`\d\d\d\d`)
	numbers := r.FindAllString(string(data), 2)
	row, _ := strconv.Atoi(numbers[0])
	col, _ := strconv.Atoi(numbers[1])
	return row, col
}

func getN(row, col int) int {
	end := row + col - 1
	sum := 0 
	for i := 0; i < end; i++ {
		sum += i
	}
	return sum + col
}

func calculateNth(n int) int {
	start := 20151125
	multiplier := 252533
	divisor := 33554393
	for i := 1; i < n; i++ {
		start *= multiplier
		start %= divisor
	}
	return start
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
		row, col := getCoord(data)
		n := getN(row, col)
		code := calculateNth(n)

		fmt.Println("The solution is: ", code)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("The solution is: ", data)
	}
}
