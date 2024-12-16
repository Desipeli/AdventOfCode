package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(inputBytes), " ")

	fmt.Println(input)

	memo := map[string]map[int]int{}

	stones := 0
	for _, v := range input {
		stones += blink(0, 75, v, memo)
	}

	fmt.Println(stones)
}

func blink(i, max int, value string, memo map[string]map[int]int) (stones int) {
	if i == max {
		// fmt.Println(value)
		return 1
	}

	mappedValues, ok := memo[value]
	if ok {
		s, ok := mappedValues[i]
		if ok {
			return s
		}
	}

	if value == "0" {
		stones = blink(i+1, max, "1", memo)
	} else if len(value)%2 == 0 {
		left := value[:len(value)/2]
		right := value[len(value)/2:]
		ri, err := strconv.Atoi(right)
		// ylimääräiset nollat pois
		if err != nil {
			log.Fatal(err)
		}
		right = strconv.Itoa(ri)
		lv := blink(i+1, max, left, memo)
		rv := blink(i+1, max, right, memo)
		stones = lv + rv

	} else {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		stones = blink(i+1, max, strconv.Itoa(intValue*2024), memo)
	}

	_, ok = memo[value]
	if !ok {
		memo[value] = map[int]int{}
	}
	memo[value][i] = stones
	return stones
}
