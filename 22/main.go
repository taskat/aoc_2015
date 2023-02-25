package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func createBoss(data []byte) Fighter {
	lines := strings.Split(string(data), "\r\n")
	boss := Fighter{effects: make([]Effect, 0)}
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[1])
		switch i {
		case 0:
			boss.hitpoints = value
		case 1:
			boss.damage = value
		}
	}
	return boss
}

type SpellList struct {
	data    []int
	current int
}

func NewSpellList() SpellList {
	return SpellList{data: []int{0}, current: -1}
}

func (sl *SpellList) GetNext() int {
	sl.current++
	if sl.current >= len(sl.data) {
		sl.data = append(sl.data, 0)
		return 0
	}
	return sl.data[sl.current]
}

func (sl *SpellList) SpellFailed() int {
	sl.data[sl.current]++
	if sl.data[sl.current] <= 4 {
		return sl.data[sl.current]
	}
	for i := sl.current - 1; i >= 0; i-- {
		sl.data[i]++
		if sl.data[i] <= 4 {
			sl.data[i]--
			sl.data = sl.data[:i+1]
			break
		}
	}
	return -1
}

func (sl *SpellList) Update() bool {
	sl.current = -1
	for i := len(sl.data) - 1; i >= 0; i-- {
		sl.data[i]++
		if sl.data[i] <= 4 {
			sl.data = sl.data[:i+1]
			return true
		}
		sl.data[i] = 0
	}
	return false
}

type Fighter struct {
	hitpoints int
	armor     int
	damage    int
	mana      int
	effects   []Effect
	spellList *SpellList
	usedMana  int
	limit     int
}

func (f Fighter) Clone() Fighter {
	return Fighter{hitpoints: f.hitpoints, armor: f.armor, damage: f.damage, mana: f.mana}
}

func (f *Fighter) attack(i int) {
	f.hitpoints -= i
}

func (f *Fighter) heal(i int) {
	f.hitpoints += i
}

func (f *Fighter) addEffect(e Effect) {
	f.effects = append(f.effects, e)
	f.armor += e.armor
}

func (f *Fighter) activateEffects() {
	for i := 0; i < len(f.effects); i++ {
		f.effects[i].activate(f)
		if f.effects[i].turns == 0 {
			f.armor -= f.effects[i].armor
			f.effects = append(f.effects[:i], f.effects[i+1:]...)
			i--
		}
	}
}

func (f *Fighter) castSpell(spells []Spell, boss *Fighter) bool {
	for next := f.spellList.GetNext(); ; {
		if !spells[next].Cast(f, boss) {
			next = f.spellList.SpellFailed()
			if next == -1 {
				return false
			}
		} else {
			f.usedMana += spells[next].GetCost()
			return f.usedMana <= f.limit
		}
	}
}

func (f *Fighter) isDead() bool {
	return f.hitpoints <= 0
}

func (f *Fighter) attackedBy(other Fighter) {
	fullDamage := other.damage - f.armor
	if fullDamage < 1 {
		fullDamage = 1
	}
	f.hitpoints -= fullDamage
}

type Effect struct {
	turns  int
	armor  int
	damage int
	mana   int
}

func (e *Effect) activate(f *Fighter) {
	e.turns--
	f.attack(e.damage)
	f.mana += e.mana
}

type Spell interface {
	Castable(player, boss *Fighter) bool
	Cast(player, boss *Fighter) bool
	GetCost() int
}

type spell struct {
	cost int
}

func (s spell) GetCost() int {
	return s.cost
}

type MagicMissile struct{ spell }

func (m MagicMissile) Castable(player, boss *Fighter) bool {
	return m.cost <= player.mana
}

func (m MagicMissile) Cast(player, boss *Fighter) bool {
	if !m.Castable(player, boss) {
		return false
	}
	player.mana -= m.cost
	boss.attack(4)
	return true
}

type Drain struct{ spell }

func (d Drain) Castable(player, boss *Fighter) bool {
	return d.cost <= player.mana
}

func (d Drain) Cast(player, boss *Fighter) bool {
	if !d.Castable(player, boss) {
		return false
	}
	player.mana -= d.cost
	boss.attack(2)
	player.heal(2)
	return true
}

type Shield struct{ spell }

func (s Shield) Castable(player, boss *Fighter) bool {
	if s.cost > player.mana {
		return false
	}
	for _, e := range player.effects {
		if e.armor == 7 {
			return false
		}
	}
	return true
}

func (s Shield) Cast(player, boss *Fighter) bool {
	if !s.Castable(player, boss) {
		return false
	}
	player.mana -= s.cost
	player.addEffect(Effect{turns: 6, armor: 7})
	return true
}

type Poison struct{ spell }

func (p Poison) Castable(player, boss *Fighter) bool {
	if p.cost > player.mana {
		return false
	}
	if len(boss.effects) == 1 {
		return false
	}
	return true
}

func (p Poison) Cast(player, boss *Fighter) bool {
	if !p.Castable(player, boss) {
		return false
	}
	player.mana -= p.cost
	boss.addEffect(Effect{turns: 6, damage: 3})
	return true
}

type Recharge struct{ spell }

func (r Recharge) Castable(player, boss *Fighter) bool {
	if r.cost > player.mana {
		return false
	}
	for _, e := range player.effects {
		if e.mana == 101 {
			return false
		}
	}
	return true
}

func (r Recharge) Cast(player, boss *Fighter) bool {
	if !r.Castable(player, boss) {
		return false
	}
	player.mana -= r.cost
	player.addEffect(Effect{turns: 5, mana: 101})
	return true
}

func createSpells() []Spell {
	return []Spell{
		MagicMissile{spell: spell{cost: 53}},
		Drain{spell: spell{cost: 73}},
		Shield{spell: spell{cost: 113}},
		Poison{spell: spell{cost: 173}},
		Recharge{spell: spell{cost: 229}},
	}
}

func fight(player, boss *Fighter, spells []Spell, isHard bool) bool {
	for {
		if isHard {
			player.attack(1)
			if player.isDead() {
				return false
			}
		}
		player.activateEffects()
		boss.activateEffects()
		if boss.isDead() {
			return true
		}
		if !player.castSpell(spells, boss) {
			return false
		}
		if boss.isDead() {
			return true
		}
		player.activateEffects()
		boss.activateEffects()
		if boss.isDead() {
			return true
		}
		player.attackedBy(*boss)
		if player.isDead() {
			return false
		}
	}
}

func start(player, boss Fighter, isHard bool) int {
	spells := createSpells()
	spellList := NewSpellList()
	minCost := 10000000000000
	for {
		clonedPlayer := player.Clone()
		clonedPlayer.limit = minCost
		clonedPlayer.spellList = &spellList
		clonedBoss := boss.Clone()
		if fight(&clonedPlayer, &clonedBoss, spells, isHard) {
			cost := clonedPlayer.usedMana
			if cost < minCost {
				minCost = cost
			}
		}
		if !spellList.Update() {
			return minCost
		}
	}
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
		player := Fighter{hitpoints: 50, mana: 500, effects: make([]Effect, 0)}
		minCost := start(player, boss, false)

		fmt.Println("The solution is: ", minCost)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		boss := createBoss(data)
		player := Fighter{hitpoints: 50, mana: 500, effects: make([]Effect, 0)}
		minCost := start(player, boss, true)

		fmt.Println("The solution is: ", minCost)
	}
}
