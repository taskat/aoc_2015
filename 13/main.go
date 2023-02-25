package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taskat/golang-utility/array"
	"github.com/taskat/golang-utility/unique"
)

type pair struct {
	people [2]string
	happiness int
}

func (p pair) Equals(other unique.Item) bool {
	otherPair, ok := other.(pair)
	return ok && (p.people[0] == otherPair.people[0] && p.people[1] == otherPair.people[1]) ||
		(p.people[0] == otherPair.people[1] && p.people[1] == otherPair.people[0])
}

func createPairs(data []byte) unique.UniqueArray {
	arr := unique.Create(nil)
	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		line = strings.Replace(line, ".", "", 1)
		words := strings.Split(line, " ")
		p := pair{}
		p.people[0] = words[0]
		p.people[1] = words[len(words) - 1]
		p.happiness, _ = strconv.Atoi(words[3])
		if strings.Contains(line, "lose") {
			p.happiness *= -1
		}
		if !arr.Push(p) {
			index := arr.GetIndex(p)
			other := arr.Get(index)
			otherPair := other.(pair)
			otherPair.happiness += p.happiness
			arr.Set(otherPair, index)
		}
	}
	return arr
}

func getPeople(arr unique.UniqueArray) unique.UniqueStringArray {
	people := unique.CreatStrings(nil)
	for i := 0; i < arr.Len(); i++ {
		people.Push(arr.Get(i).(pair).people[0])
		people.Push(arr.Get(i).(pair).people[1])
	}
	return people
}

func getOrders(arr unique.UniqueArray) [][]string {
	people := getPeople(arr)
	return array.PermutateString(people.GetData())
}

func getHappiness(orders [][]string, pairs unique.UniqueArray) []int {
	happinesses := make([]int, 0, len(orders))
	for _, order := range orders {
		currentHappiness := 0
		for i := 1; i < len(order); i++ {
			idx := pairs.GetIndex(pair{people: [2]string{order[i - 1], order[i]}})
			currentHappiness += pairs.Get(idx).(pair).happiness
		}
		idx := pairs.GetIndex(pair{people: [2]string{order[0], order[len(order) - 1]}})
		currentHappiness += pairs.Get(idx).(pair).happiness
		happinesses = append(happinesses, currentHappiness)
	}
	return happinesses
}

func addMe(pairs unique.UniqueArray) unique.UniqueArray {
	people := getPeople(pairs)
	for _, person := range people.GetData() {
		pairs.Push(pair{people: [2]string{person, "Me"}, happiness: 0})
	}
	return pairs
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
		pairs := createPairs(data)
		orders := getOrders(pairs)
		happinesses := getHappiness(orders, pairs)

		fmt.Println("The solution is: ", array.Max(happinesses))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		pairs := createPairs(data)
		pairs = addMe(pairs)
		orders := getOrders(pairs)
		happinesses := getHappiness(orders, pairs)

		fmt.Println("The solution is: ", array.Max(happinesses))
	}
}
