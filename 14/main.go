package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taskat/golang-utility/array"
)

var (
	endTime = 2503
)

type Reindeer struct {
	name string
	speed int
	flyTime int
	restTime int
}

func (r Reindeer) distance(time int) int {
	dist := 0
	for r.flyTime < time {
		dist += r.flyTime * r.speed
		time -= r.flyTime + r.restTime
	}
	if time > 0 {
		dist += time * r.speed
	}
	return dist
}

func CreateReindeers(data []byte) []Reindeer {
	lines := strings.Split(string(data), "\r\n")
	reindeers := make([]Reindeer, 0, len(lines))
	for _, line := range lines {
		words := strings.Split(line, " ")
		name := words[0]
		speed, _ := strconv.Atoi(words[3])
		flyTime, _ := strconv.Atoi(words[6])
		restTime, _ := strconv.Atoi(words[len(words) - 2])
		reindeers = append(reindeers, Reindeer{name: name, speed: speed,
			flyTime: flyTime, restTime: restTime})
	}
	return reindeers
}

func getDistance(reindeers []Reindeer, time int) map[Reindeer]int {
	distances := make(map[Reindeer]int)
	for _, reindeer := range reindeers {
		distances[reindeer] = reindeer.distance(time)
	}
	return distances
}

func mapToArr(distances map[Reindeer]int) []int {
	arr := make([]int, 0, len(distances))
	for _, value := range distances {
		arr = append(arr, value)
	}
	return arr
}

func getMax(distances map[Reindeer]int) []Reindeer {
	max := 0
	reindeers := make([]Reindeer, 0)
	for reindeer, dist := range distances {
		if dist == max {
			reindeers = append(reindeers, reindeer)
		}
		if dist > max {
			max = dist
			reindeers = make([]Reindeer, 0)
			reindeers = append(reindeers, reindeer)
		}
	}
	return reindeers
}

func getPoints(reindeers []Reindeer) map[Reindeer]int {
	points := make(map[Reindeer]int)
	for i := 1; i <= endTime; i++ {
		distances := getDistance(reindeers, i)
		winners := getMax(distances)
		for _, winner := range winners {
			points[winner]++
		}
	}
	return points
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
		reindeers := CreateReindeers(data)
		distances := getDistance(reindeers, endTime)

		fmt.Println("The solution is: ", array.Max(mapToArr(distances)))
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		reindeers := CreateReindeers(data)
		points := getPoints(reindeers)

		fmt.Println("The solution is: ", array.Max(mapToArr(points)))
	}
}
