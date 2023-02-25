package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func turn(s string) string {
	newS := make([]byte, 0, len(s))
	lastByte := s[0]
	sameCount := 1
	for i := 1; i < len(s); i++ {
		if s[i] == lastByte {
			sameCount++
		} else {
			newS = append(newS, strconv.Itoa(sameCount)[0])
			newS = append(newS, lastByte)
			lastByte = s[i]
			sameCount = 1
		}
	}
	newS = append(newS, strconv.Itoa(sameCount)[0])
	newS = append(newS, lastByte)
	return string(newS)
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
		s := string(data)
		for i := 0; i < 40; i++ {
			s = turn(s)
		}

		fmt.Println("The solution is: ", len(s))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		s := string(data)
		for i := 0; i < 50; i++ {
			s = turn(s)
		}

		fmt.Println("The solution is: ", len(s))
	}
}
