package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func sumMap(m map[string]interface{}, filter string) int {
	sum := 0
	for _, value := range m {
		s, ok := value.(string)
		if ok && s == filter {
			return 0
		}
		sum += sumAny(value, filter)
	}
	return sum
}

func sumArray(arr []interface{}, filter string) int {
	sum := 0
	for _, value := range arr {
		sum += sumAny(value, filter)
	}
	return sum
}

func sumAny(any interface{}, filter string) int {
	switch any.(type) {
	case []interface{}:
		return sumArray(any.([]interface{}), filter)
	case float64:
		return int(any.(float64))
	case map[string]interface{}:
		return sumMap(any.(map[string]interface{}), filter)
	default:
		return 0
	}
}

func sumNums(data []byte, filter string) int {
	var j interface{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		panic("invalid json")
	}
	return sumAny(j, filter)
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
		sum := sumNums(data, "")

		fmt.Println("The solution is: ", sum)
	case 2:
		fmt.Println("Solving part 2")
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		sum := sumNums(data, "red")

		fmt.Println("The solution is: ", sum)
	}
}
