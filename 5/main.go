package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getStrings(data []byte) []string {
	lines := strings.Split(string(data), "\r\n")
	return lines
}

func hasThreeVowels(s string) bool {
	return strings.Count(s, "a") + strings.Count(s, "e") + strings.Count(s, "i") +
		strings.Count(s, "o") + strings.Count(s, "u") >= 3
}

func hasDouble(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i - 1] {
			return true
		}
	}
	return false
}

func notContains(s string) bool {
	naughties := []string{"ab", "cd", "pq", "xy"}
	for _, naughty := range naughties {
		if strings.Contains(s, naughty) {
			return false
		}
	}
	return true
}

func isNice(s string) bool {
	return hasThreeVowels(s) && hasDouble(s) && notContains(s)
}

func hasDoubleRepeat(s string) bool {
	for i := 1; i < len(s); i++ {
		if strings.Count(s, s[i - 1 : i + 1]) >= 2 {
			return true
		}
	}
	return false
}

func hasLetterWithGap(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i - 2] {
			return true
		}
	}
	return false
}

func isNewNice(s string) bool {
	return hasDoubleRepeat(s) && hasLetterWithGap(s)
}

func filterNices(strings []string, filter func(string) bool) []string {
	nices := make([]string, 0)
	for _, s := range strings {
		if filter(s) {
			nices = append(nices, s)
		}
	}
	return nices
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
		nices := filterNices(strings, isNice)

		fmt.Println("The solution is: ", len(nices))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		strings := getStrings(data)
		nices := filterNices(strings, isNewNice)

		fmt.Println("The solution is: ", len(nices))
	}
}
