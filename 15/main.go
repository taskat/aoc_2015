package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Cookie map[Ingredient]int

func (c Cookie) Score() int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	for ingredient, amount := range c {
		capacity += ingredient.capacity * amount
		durability += ingredient.durability * amount
		flavor += ingredient.flavor * amount
		texture += ingredient.texture * amount
	}
	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0
	}
	return capacity * durability * flavor * texture
}

func (c Cookie) GetCalories() int {
	calories := 0
	for ingredient, amount := range c {
		calories += ingredient.calories * amount
	}
	return calories
}

func CreateIngredients(data []byte) []Ingredient {
	lines := strings.Split(string(data), "\r\n")
	ingredients := make([]Ingredient, 0, len(lines))
	for _, line := range lines {
		line = strings.ReplaceAll(line, ",", "")
		words := strings.Split(line, " ")
		ing := Ingredient{}
		ing.name = words[0][:len(words[0]) - 1]
		ing.capacity, _ = strconv.Atoi(words[2])
		ing.durability, _ = strconv.Atoi(words[4])
		ing.flavor, _ = strconv.Atoi(words[6])
		ing.texture, _ = strconv.Atoi(words[8])
		ing.calories, _ = strconv.Atoi(words[10])
		ingredients = append(ingredients, ing)
	}
	return ingredients
}

func getRatio() func() []int {
	var last []int = nil
	return func() []int {
		if last == nil {
			last = []int{100, 0, 0, 0}
			return last
		}
		for i := 2; i >= 0; i-- {
			if last[i] > 0 {
				last[i]--
				last[i + 1]++
				return last
			}
			if last[i + 1] > 0 {
				last[i] = last[i + 1]
				last[i + 1] = 0
			}
		}
		return nil
	}
}

func CreateCookie(ingredients []Ingredient, ratio []int) Cookie {
	c := make(Cookie)
	for i, ingredient := range ingredients {
		c[ingredient] = ratio[i]
	}
	return c
}

func getBestCookie(ingredients []Ingredient, filter func(c Cookie) bool) Cookie {
	nextRatio := getRatio()
	maxScore := 0
	bestCookie := Cookie{}
	ratio := nextRatio()
	for ratio != nil {
		cookie := CreateCookie(ingredients, ratio)
		score := cookie.Score()
		if score > maxScore && filter(cookie) {
			maxScore = score
			bestCookie = cookie
		}
		ratio = nextRatio()
	}
	return bestCookie
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
		ingredients := CreateIngredients(data)
		best := getBestCookie(ingredients, func(Cookie) bool {return true})

		fmt.Println("The solution is: ", best.Score())
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		ingredients := CreateIngredients(data)
		best := getBestCookie(ingredients, func(c Cookie) bool {return c.GetCalories() == 500})

		fmt.Println("The solution is: ", best.Score())
	}
}
