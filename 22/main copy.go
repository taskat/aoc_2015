package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// type Stack struct {
// 	data []Spell
// }

// func (s *Stack) Push(spell Spell) {
// 	s.data = append(s.data)
// }

// func (s *Stack) Pop() Spell {
// 	spell := s.data[len(s.data)]
// 	s.data = s.data[:len(s.data)-1]
// 	return spell
// }

// func (s *Stack) Top() Spell {
// 	return s.data[len(s.data)]
// }

// type Fighter struct {
// 	hitpoints int
// 	damage    int
// 	armor     int
// 	mana      int
// }

// func (f *Fighter) effected(s Spell) {
// 	f.hitpoints += s.heal
// 	f.hitpoints -= s.damage
// 	f.armor += s.armor
// 	f.mana += s.mana
// }

// func (f *Fighter) effectedInverse(s Spell) {
// 	f.hitpoints -= s.heal
// 	f.hitpoints += s.damage
// 	f.armor -= s.armor
// 	f.mana -= s.mana
// }

// func (f *Fighter) attackedBy(other Fighter) {
// 	fullDamage := other.damage - f.armor
// 	if fullDamage < 1 {
// 		fullDamage = 1
// 	}
// 	f.hitpoints -= fullDamage
// }

// var calls = 0
// func (f *Fighter) chooseSpell(spells []Spell, active map[Spell]*Effect) (Spell, bool, bool) {
// 	defer func() {calls++}()
// 	switch calls {
// 	case 0: return spells[3]
// 	}
// }

// func (f Fighter) isDead() bool {
// 	return f.hitpoints <= 0
// }

// type Spell struct {
// 	cost   int
// 	damage int
// 	heal   int
// 	armor  int
// 	mana   int
// 	turns int
// }

// type Effect struct {
// 	turns  int
// 	target *Fighter
// }

// func createSpells() []Spell {
// 	return []Spell{
// 		{cost: 53, damage: 4},
// 		{cost: 73, damage: 2, heal: 2},
// 		{cost: 113, armor: 7, turns: 6},
// 		{cost: 173, damage: 3, turns: 6},
// 		{cost: 229, mana: 101, turns: 5},
// 	}
// }

// func createBoss(data []byte) Fighter {
// 	lines := strings.Split(string(data), "\r\n")
// 	boss := Fighter{}
// 	for i, line := range lines {
// 		parts := strings.Split(line, ": ")
// 		value, _ := strconv.Atoi(parts[1])
// 		switch i {
// 		case 0:
// 			boss.hitpoints = value
// 		case 1:
// 			boss.damage = value
// 		}
// 	}
// 	return boss
// }

// type Game struct {
// 	player  Fighter
// 	boss    Fighter
// 	spells  []Spell
// 	effects map[Spell]*Effect
// 	stack   Stack
// }

// func (g *Game) doEffects() {
// 	for s, e := range g.effects {
// 		e.target.effected(s)
// 		e.turns--
// 		if e.turns == 0 {
// 			delete(g.effects, s)
// 		}
// 	}
// }

// func (g *Game) removeArmor() {
// 	for s, e := range g.effects {
// 		if s.armor != 0 {
// 			e.target.effectedInverse(s)
// 		}
// 	}
// }

// func (g *Game) fullRound() bool {
// 	g.doEffects()
// 	spell, self, lost := g.player.chooseSpell(g.spells, g.effects)
// 	if lost {
// 		return true
// 	}
// 	target := g.boss
// 	if self {
// 		target = g.player
// 	}
// 	g.effects[spell] = &Effect{turns: spell.turns, target: &target}
// 	if g.boss.isDead() {
// 		return true
// 	}
// 	g.removeArmor()
// 	g.doEffects()
// 	g.player.attackedBy(g.boss)
// 	g.removeArmor()
// 	if g.player.isDead() {
// 		return true
// 	}
// 	return false
// }

// func (g *Game) run() bool {
// 	for !g.fullRound() {
// 		//empty
// 	}
// 	return !g.player.isDead() && g.boss.isDead()
// }

// func main() {
// 	part := 0
// 	validAnswer := false
// 	for !validAnswer {
// 		fmt.Println("Which part? (1 or 2)")
// 		fmt.Scanf("%d\n", &part)
// 		if part < 1 || part > 2 {
// 			fmt.Println("Invalid answer!")
// 			continue
// 		}
// 		validAnswer = true
// 	}
// 	switch part {
// 	case 1:
// 		fmt.Println("Solving part 1")
// 		data, err := ioutil.ReadFile("data.txt")
// 		if err != nil {
// 			panic(err)
// 		}
// 		boss := createBoss(data)
// 		player := Fighter{50, 0, 0, 500}

// 		fmt.Println("The solution is: ", boss)
// 	case 2:
// 		fmt.Println("Solving part 2")
// 		data, err := ioutil.ReadFile("data.txt")
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("The solution is: ", data)
// 	}
// }
