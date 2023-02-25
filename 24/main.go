package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
	sum  int
}

func NewStack() Stack {
	return Stack{data: make([]int, 0), sum: 0}
}

func (s Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() int {
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.sum -= item
	return item
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
	s.sum += item
}

func (s Stack) Top() int {
	return s.data[len(s.data)-1]
}

func createPackets(data []byte) []int {
	lines := strings.Split(string(data), "\r\n")
	packets := make([]int, 0, len(lines))
	for _, line := range lines {
		size, _ := strconv.Atoi(line)
		packets = append(packets, size)
	}
	return packets
}

func getGroupSize(packets []int, groups int) int {
	sum := 0
	for _, packet := range packets {
		sum += packet
	}
	return sum / groups
}

func getQE(packets []int) int {
	qe := 1
	for _, packet := range packets {
		qe *= packet
	}
	return qe
}

func reverse(arr []int) []int {
	newArr := make([]int, len(arr))
	for i, value := range arr {
		newArr[len(newArr) - i - 1] = value
	}
	return newArr
}

func findGroups(packets []int, size, all, into int) [][]int {
	firstGroup := NewStack()
	remaining := NewStack()
	groups := make([][]int, 0)
	for i := len(packets) - 1; i >= 0 ; i-- {
		firstGroup.Push(packets[i])
		if firstGroup.sum == size {
			if into == 2 {
				return [][]int{firstGroup.data, append(remaining.data, packets[:i]...)}
			}
			newPackets := make([]int, i)
			copy(newPackets, packets[:i])
			newPackets = append(newPackets, reverse(remaining.data)...)
			if len(findGroups(newPackets, size, all, into - 1)) > 0 {
				copied := make([]int, firstGroup.Len())
				copy(copied, firstGroup.data)
				groups = append(groups, copied)
				if into < all {
					return groups
				}
			}
		}
		if i == 0 {	
			firstGroup.Pop()
			i++
			for length := firstGroup.Len(); length == firstGroup.Len(); {
				if firstGroup.Len() == 0 {
					return groups
				}
				if remaining.Len() == 0 || firstGroup.Top() < remaining.Top() {
					remaining.Push(firstGroup.Pop())
				} else {
					remaining.Pop()
					i++
				}
			}
		} else if firstGroup.sum >= size {
			for length := firstGroup.Len(); length == firstGroup.Len(); {
				if firstGroup.Len() == 0 {
					return groups
				}
				if remaining.Len() == 0 || firstGroup.Top() < remaining.Top() {
					remaining.Push(firstGroup.Pop())
				} else {
					remaining.Pop()
					i++
				}
			}
		}
	}
	return groups
}

func getMinSizedGroups(groups [][]int) [][]int{
	size := len(groups[0])
	mins := [][]int{groups[0]}
	for _, group := range groups {
		if len(group) < size {
			size = len(group)
			mins = [][]int{group}
		} else if len(group) == size {
			mins = append(mins, group)
		}
	}
	return mins
}

func getMinQE(groups [][]int) int {
	qe := make([]int, len(groups))
	for i, group := range groups {
		qe[i] = getQE(group)
	}
	minQe := qe[0]
	minIdx := 0
	for i, currentQe := range qe {
		if currentQe < minQe {
			minQe = currentQe
			minIdx = i
		}
	}
	fmt.Println(minIdx)
	fmt.Println(groups[minIdx])
	return minQe
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
		packets := createPackets(data)
		groups := findGroups(packets, getGroupSize(packets, 3), 3, 3)
		minGroups := getMinSizedGroups(groups)
		minQe := getMinQE(minGroups)

		fmt.Println("The solution is: ", minQe)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		packets := createPackets(data)
		groups := findGroups(packets, getGroupSize(packets, 4), 4, 4)
		minGroups := getMinSizedGroups(groups)
		minQe := getMinQE(minGroups)

		fmt.Println("The solution is: ", minQe)
	}
}
