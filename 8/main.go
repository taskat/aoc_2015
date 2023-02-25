package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getStrings(data []byte) []string {
	return strings.Split(string(data), "\r\n")
}

func codeLength(strings []string) int {
	sum := 0
	for _, s := range strings {
		sum += len(s)
	}
	return sum
}

func memoryLength(strings []string) int {
	sum := 0
	for _, s := range strings {
		s, _ = strconv.Unquote(s)
		sum += len(s)
	}
	return sum
}

func encodedLength(strings []string) int {
	sum := 0
	for _, s := range strings {
		sum += len(strconv.Quote(s))
	}
	return sum
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
		strings := getStrings(data)

		fmt.Println("The solution is: ", codeLength(strings) - memoryLength(strings))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		strings := getStrings(data)

		fmt.Println("The solution is: ", encodedLength(strings) - codeLength(strings))
	}
}
