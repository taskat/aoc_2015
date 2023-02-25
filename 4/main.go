package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
)

func startWithNZeros(hash string, zeros int) bool {
	for i := 0; i < zeros; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

func getSmallest(key string, zeros int) int {
	for i := 0; ; i++ {
		fullKey := key + strconv.Itoa(i)
		hash := md5.Sum([]byte(fullKey))
		if startWithNZeros(hex.EncodeToString(hash[:]), zeros) {
			return i
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
		key := string(data)
		smallest := getSmallest(key, 5)

		fmt.Println("The solution is: ", smallest)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		key := string(data)
		smallest := getSmallest(key, 6)

		fmt.Println("The solution is: ", smallest)
	}
}
