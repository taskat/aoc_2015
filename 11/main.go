package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/taskat/golang-utility/unique"
)

func hasIncreasing(password string) bool {
	for i := 2; i < len(password); i++ {
		if password[i-2] == password[i]-2 && password[i-1] == password[i]-1 {
			return true
		}
	}
	return false
}

func notContains(password string) bool {
	return !strings.ContainsAny(password, "iol")
}

func containsDoublePair(password string) bool {
	pairs := unique.CreateInts(nil)
	for i := 1; i < len(password); i++ {
		if password[i] == password[i-1] {
			pairs.Push(int(password[i]))
			i++
		}
	}
	return pairs.Len() >= 2
}

func isValid(password string) bool {
	return notContains(password) && hasIncreasing(password) && containsDoublePair(password)
}

func increment(old []byte, i int) {
	old[i]++
	if old[i] == 'z'+1 {
		old[i] = 'a'
		if i != 0 {
			increment(old, i-1)
		}
	}
}

func next(old []byte) {
	increment(old, len(old) - 1)
	for !isValid(string(old)) {
		increment(old, len(old) - 1)
	}
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
		password := data
		next(password)

		fmt.Println("The solution is: ", string(password))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		password := data
		next(password)
		next(password)

		fmt.Println("The solution is: ", string(password))
	}
}
