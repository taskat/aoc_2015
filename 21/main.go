package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Fighter struct {
	hitpoints int
	damage    int
	armor     int
}

func (f *Fighter) addGear(g Gear) {
	f.damage += g.damage
	f.armor += g.armor
}

func (f *Fighter) removeGear(g Gear) {
	f.damage -= g.damage
	f.armor -= g.armor
}

func (f * Fighter) attackedBy(other Fighter) {
	fullDamage := other.damage - f.armor
	if fullDamage < 1 {
		fullDamage = 1
	}
	f.hitpoints -= fullDamage
}

func (f Fighter) isDead() bool {
	return f.hitpoints <= 0
}

type Gear struct {
	cost   int
	damage int
	armor  int
}

func createWeapons() []Gear {
	return []Gear{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
}

func createArmors() []Gear {
	return []Gear{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}
}

func createRings() []Gear {
	return []Gear{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}
}

func createBoss(data []byte) Fighter {
	lines := strings.Split(string(data), "\r\n")
	boss := Fighter{}
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[1])
		switch i {
		case 0: boss.hitpoints = value
		case 1: boss.damage = value
		case 2: boss.armor = value
		}
	}
	return boss
}

func fight(player, boss Fighter) bool {
	for {
		boss.attackedBy(player)
		if boss.isDead() {
			return true
		}
		player.attackedBy(boss)
		if player.isDead() {
			return false
		}
	}
}

func minwin(result bool, cost, prevCost int) int {
	if result && cost < prevCost {
		return cost
	}
	return prevCost
}

func maxlose(result bool, cost, prevCost int) int {
	if !result && cost > prevCost {
		return cost
	}
	return prevCost
}

func bruteFroce(player, boss Fighter, oracle func(result bool, cost, prev int) int, defaultcost int) int {
	weapons := createWeapons()
	armors := createArmors()
	rings := createRings()
	prevCost := defaultcost
	for w := 0; w < len(weapons); w++ {
		for a := 0; a < len(armors); a++ {
			for r1 := 0; r1 < len(rings) - 1; r1++ {
				for r2 := r1 + 1; r2 < len(rings); r2++ {
					player.addGear(weapons[w])
					player.addGear(armors[a])
					player.addGear(rings[r1])
					player.addGear(rings[r2])
					fullCost := weapons[w].cost + armors[a].cost + rings[r1].cost + rings[r2].cost
					prevCost = oracle(fight(player, boss), fullCost, prevCost)
					player.removeGear(weapons[w])
					player.removeGear(armors[a])
					player.removeGear(rings[r1])
					player.removeGear(rings[r2])
				}
			}
		}
	}
	return prevCost
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
		boss := createBoss(data)
		player := Fighter{100, 0, 0}
		minCost := bruteFroce(player, boss, minwin, 500)

		fmt.Println("The solution is: ", minCost)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		boss := createBoss(data)
		player := Fighter{100, 0, 0}
		maxCost := bruteFroce(player, boss, maxlose, 0)

		fmt.Println("The solution is: ", maxCost)
	}
}
