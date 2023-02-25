package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func read(data []byte) int {
	i, _ := strconv.Atoi(string(data))
	return i
}

type House struct {
	number  int
	gifts   int
	primefactors []int
}

func (h House) getValue() float64 {
	return float64(h.gifts) / float64(h.number)
}

func (h House) copy() House {
	newh := House{number: h.number, gifts: h.gifts, primefactors: make([]int, len(h.primefactors))}
	copy(newh.primefactors, h.primefactors)
	return newh
}

func (h House) calculateGifts() int {
	factors := make(map[int]int)
	for subsetBits := 0; subsetBits < (1 << len(h.primefactors)); subsetBits++ {
		curr := 1

		for object := uint(0); object < uint(len(h.primefactors)); object++ {
			if (subsetBits>>object)&1 == 1 {
				curr *= h.primefactors[object]
			}
		}
		factors[curr] = 1
	}

	sum := 0
	for factor := range factors {
		sum += factor
	}
	return sum * 10
}

func (h House) calculateGifts2() int {
	factors := make(map[int]int)
	for subsetBits := 0; subsetBits < (1 << len(h.primefactors)); subsetBits++ {
		curr := 1

		for object := uint(0); object < uint(len(h.primefactors)); object++ {
			if (subsetBits>>object)&1 == 1 {
				curr *= h.primefactors[object]
			}
		}
		factors[curr] = 1
	}

	sum := 0
	for factor := range factors {
		if factor * 50 >= h.number {
			sum += factor
		}
	}
	return sum * 11
}

func (h *House) addFactor(f int, calculate func(House) int) {
	h.primefactors = append(h.primefactors, f)
	h.number *= f
	h.gifts = calculate(*h)
}

func max(arr []House) int {
	maxidx := 0
	max := arr[0].getValue()
	for i, h := range arr {
		if h.getValue() > max {
			max = h.getValue()
			maxidx = i
		}
	}
	return maxidx
}

func min(arr []House) int {
	minidx := 0
	min := arr[0].number
	for i, h := range arr {
		if h.number < min {
			min = h.number
			minidx = i
		}
	}
	return minidx
}

func isPrime(a int) bool {
	for i := 3; i <= a / 3; i += 2 {
		if a % i == 0 {
			return false
		}
	}
	return true
}

func getPrimes() []int {
	primes := make([]int, 0)
	primes = append(primes, 2)
	for i := 3; i < 100_000; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func algorithm(goal int, calculate func(House) int) int {
	h := House{number: 1, gifts: 10, primefactors: []int{1}}
	plus := getPrimes()
	opps := make([]House, 1)
	opps[0] = h
	solutions := make([]House, 0)
	next := 0
	for 0 < len(opps) {
		for _, p := range plus {
			newh := opps[next].copy()
			newh.addFactor(p, calculate)
			if newh.gifts >= goal {
				solutions = append(solutions, newh)
			} else {
				opps = append(opps, newh)
			}
		}
		opps = append(opps[:next], opps[next+1:]...)
		next = max(opps)
		if len(solutions) > 100_000 {
			idx := min(solutions)
			return solutions[idx].number
		}
	}
	return -1
}

func getPrimeFactors(a int) []int {
	factors := make([]int, 0)
	for i := 2; i <= a; i++ {
		if a % i == 0 {
			factors = append(factors, i)
			a /= i
			i--
		}
	}
	return factors
}

func brute(from, to, goal int) int {
	for i := from; i < to; i++ {
		h := House{number: i, primefactors: getPrimeFactors(i)}
		h.gifts = h.calculateGifts2()
		if h.gifts > goal {
			return i
		}
	}
	return -1
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
		goal := read(data)
		house := algorithm(goal, House.calculateGifts)

		fmt.Println("The solution is: ", house)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		goal := read(data)
		house := brute(776160, 803880, goal)

		fmt.Println("The solution is: ", house)
	}
}
