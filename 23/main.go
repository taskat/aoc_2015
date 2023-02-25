package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Computer map[string]int

type Instruction interface {
	Do(c *Computer) int
}

type Hlf struct {
	register string
}

func (h Hlf) Do(c *Computer) int {
	(*c)[h.register] /= 2
	return 1
}

type Tpl struct {
	register string
}

func (t Tpl) Do(c *Computer) int {
	(*c)[t.register] *= 3
	return 1
}

type Inc struct {
	register string
}

func (i Inc) Do(c *Computer) int {
	(*c)[i.register]++
	return 1
}

type Jmp struct {
	offset int
}

func (j Jmp) Do(c *Computer) int {
	return j.offset
}

type Jie struct {
	register string
	offset int
}

func (j Jie) Do(c *Computer) int {
	if (*c)[j.register] % 2 == 0 {
		return j.offset
	}
	return 1
}

type Jio struct {
	register string
	offset int
}

func (j Jio) Do(c *Computer) int {
	if (*c)[j.register] == 1 {
		return j.offset
	}
	return 1
}

func createInstructions(data []byte) []Instruction {
	lines := strings.Split(string(data), "\r\n")
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		var instruction Instruction
		switch line[:3] {
		case "hlf": instruction = Hlf{register: line[4:]}
		case "tpl": instruction = Tpl{register: line[4:]}
		case "inc": instruction = Inc{register: line[4:]}
		case "jmp":
			offset, _ := strconv.Atoi(line[4:])
			instruction = Jmp{offset: offset}
		case "jie":
			register := line[4:5]
			offset, _ := strconv.Atoi(line[7:])
			instruction = Jie{register: register, offset: offset}
		case "jio":
			register := line[4:5]
			offset, _ := strconv.Atoi(line[7:])
			instruction = Jio{register: register, offset: offset}
		}
		instructions = append(instructions, instruction)
	}
	return instructions
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
		instructions := createInstructions(data)
		computer := Computer{"a": 0, "b": 0}
		for i := 0; i < len(instructions); {
			if i < 0 || i >= len(instructions) {
				break
			}
			i += instructions[i].Do(&computer)
		}

		fmt.Println("The solution is: ", computer["b"])
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		instructions := createInstructions(data)
		computer := Computer{"a": 1, "b": 0}
		for i := 0; i < len(instructions); {
			if i < 0 || i >= len(instructions) {
				break
			}
			i += instructions[i].Do(&computer)
		}

		fmt.Println("The solution is: ", computer["b"])
	}
}
