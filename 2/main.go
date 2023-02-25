package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taskat/golang-utility/intmath"
)

type box struct {
	sides []int
	faces []int
}

func createBoxes(data []byte) []box {
	lines := strings.Split(string(data), "\r\n")
	boxes := make([]box, len(lines))
	for i, line := range lines {
		dimensions := strings.Split(line, "x")
		sides := make([]int, len(dimensions))
		for i, dims := range dimensions {
			sides[i], _ = strconv.Atoi(dims)
		}
		b := box{sides: sides, faces: make([]int, 3)}
		b.faces[0] = sides[0] * sides[1]
		b.faces[1] = sides[0] * sides[2]
		b.faces[2] = sides[1] * sides[2]
		boxes[i] = b
	}
	return boxes
}

func getWrappingpaper(boxes []box) int {
	sizes := make([]int, len(boxes))
	for i, b := range boxes {
		sizes[i] = intmath.Sum(b.faces)*2 + intmath.Min(b.faces)
	}
	return intmath.Sum(sizes)
}

func getRibbon(boxes []box) int {
	ribbons := make([]int, len(boxes))
	for i, box := range boxes {
		ribbons[i] = intmath.Sum(intmath.RemoveFirst(box.sides, intmath.Max(box.sides)))
		ribbons[i] += intmath.Product(box.sides)
	}
	return intmath.Sum(ribbons)
}

func main() {
	part := 0
	validAnswer := false
	for !validAnswer {
		fmt.Println("Which part? (1 or 2)")
		//fmt.Scanf("%d\n", &part)
		part = 2
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
		boxes := createBoxes(data)
		size := getWrappingpaper(boxes)

		fmt.Println("The solution is: ", size)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		boxes := createBoxes(data)
		ribbon := getRibbon(boxes)

		fmt.Println("The solution is: ", ribbon)
	}
}
