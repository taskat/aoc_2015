package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/taskat/golang-utility/array"
	"github.com/taskat/golang-utility/unique"
)

type replacement struct {
	from string
	to   []string
}

type molecule struct {
	atoms []string
	cost  int
}

func (m molecule) Equals(other unique.Item) bool {
	otherMolecule, ok := other.(molecule)
	if !ok || len(m.atoms) != len(otherMolecule.atoms) {
		return false
	}
	for i, atom := range m.atoms {
		if atom != otherMolecule.atoms[i] {
			return false
		}
	}
	return true
}

func (m molecule) replace(rule replacement) []molecule {
	children := make([]molecule, 0)
	for i, atom := range m.atoms {
		if atom == rule.from {
			atomsCopy := make([]string, len(m.atoms))
			copy(atomsCopy, m.atoms)
			newChild := append(append(atomsCopy[:i], rule.to...), m.atoms[i+1:]...)
			children = append(children, molecule{atoms: newChild, cost: m.cost + 1})
		}
	}
	return children
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func toMolecule(s string) molecule {
	atoms := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if isUpper(s[i]) {
				if isUpper(s[i+1]) {
					atoms = append(atoms, string(s[i]))
				} else {
					atoms = append(atoms, string(s[i])+string(s[i+1]))
					i++
				}
			} else {
				panic("Invalid molecule")
			}
		} else {
			atoms = append(atoms, string(s[i]))
		}
	}
	return molecule{atoms: atoms, cost: 0}
}

func Create(data []byte) ([]replacement, molecule) {
	parts := strings.Split(string(data), "\r\n\r\n")
	lines := strings.Split(parts[0], "\r\n")
	rules := make([]replacement, 0, len(lines))
	for _, line := range lines {
		operands := strings.Split(line, " => ")
		rules = append(rules, replacement{from: operands[0], to: toMolecule(operands[1]).atoms})
	}
	goal := toMolecule(parts[1])
	return rules, goal
}

func getChildren(m molecule, rules []replacement) unique.UniqueArray {
	arr := unique.Create(nil)
	for _, rule := range rules {
		children := m.replace(rule)
		for _, child := range children {
			arr.Push(child)
		}
	}
	return arr
}

func getMin(goal molecule) int {
	numberOfAtoms := len(goal.atoms)
	radons := array.CountString(goal.atoms, func(s string) bool {return s == "Rn"})
	argons := array.CountString(goal.atoms, func(s string) bool {return s == "Ar"})
	yttriums := array.CountString(goal.atoms, func(s string) bool {return s == "Y"})
	return numberOfAtoms - radons - argons - 2 * yttriums - 1
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
		rules, start := Create(data)
		children := getChildren(start, rules)

		fmt.Println("The solution is: ", children.Len())
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		_, goal := Create(data)
		min := getMin(goal)

		fmt.Println("The solution is: ", min)
	}
}
