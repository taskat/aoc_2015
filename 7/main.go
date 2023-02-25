package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type gateFunc func(a, b uint16) uint16

type Command struct {
	input1, input2 string
	output         string
	gate           gateFunc
}

func value(a, b uint16) uint16 {
	return a
}

func and(a, b uint16) uint16 {
	return a & b
}

func or(a, b uint16) uint16 {
	return a | b
}

func leftShift(a, b uint16) uint16 {
	return a << b
}

func rightShift(a, b uint16) uint16 {
	return a >> b
}

func not(a, b uint16) uint16 {
	return ^a
}

func createCommand(line string) Command {
	parts := strings.Split(line, " -> ")
	command := Command{output: parts[1]}
	tokens := strings.Split(parts[0], " ")
	switch len(tokens) {
	case 1:
		command.input1 = tokens[0]
		command.input2 = ""
		command.gate = value
	case 2:
		command.input1 = tokens[1]
		command.input2 = ""
		command.gate = not
	case 3:
		command.input1 = tokens[0]
		command.input2 = tokens[2]
		switch tokens[1] {
		case "AND":
			command.gate = and
		case "OR":
			command.gate = or
		case "LSHIFT":
			command.gate = leftShift
		case "RSHIFT":
			command.gate = rightShift
		}
	}
	return command
}

func createCommands(data []byte) []Command {
	lines := strings.Split(string(data), "\r\n")
	commands := make([]Command, len(lines))
	for i, line := range lines {
		commands[i] = createCommand(line)
	}
	return commands
}

func runCommands(commands []Command, overrides map[string]uint16) map[string]uint16 {
	wires := make(map[string]uint16)
	for 0 < len(commands) {
		for i := 0; i < len(commands); i++ {
			command := commands[i]
			a64, err := strconv.ParseInt(command.input1, 10, 16)
			a := uint16(a64)
			var ok bool
			if err != nil && command.input1 != "" {
				a, ok = overrides[command.input1]
				if !ok {
					a, ok = wires[command.input1]
					if !ok {
						continue
					}
				}
			}
			b64, err := strconv.ParseInt(command.input2, 10, 16)
			b := uint16(b64)
			if err != nil && command.input2 != "" {
				b, ok = overrides[command.input2]
				if !ok {
					b, ok = wires[command.input2]
					if !ok {
						continue
					}
				}
			}
			wires[command.output] = command.gate(a, b)
			commands = append(commands[:i], commands[i+1:]...)
			i--
		}
	}
	return wires
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
		commands := createCommands(data)
		wires := runCommands(commands, make(map[string]uint16))

		fmt.Println("The solution is: ", wires["a"])
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		commands := createCommands(data)
		commandsCopy := make([]Command, len(commands))
		copy(commandsCopy, commands)
		wires := runCommands(commandsCopy, make(map[string]uint16))
		wires = runCommands(commands, map[string]uint16{"b": wires["a"]})

		fmt.Println("The solution is: ", wires["a"])
	}
}
