package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	inputFile := os.Args[1]
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	inputString := string(inputBytes)

	input := strings.Split(inputString, "\n")

	var wg sync.WaitGroup
	c := make(chan int)

	for _, row := range input {
		row_split := strings.Split(row, ":")
		testValue, err := strconv.Atoi(row_split[0])
		if err != nil {
			log.Fatal(err)
		}
		operands := strings.Split(strings.TrimSpace(row_split[1]), " ")

		firstValue, err := strconv.Atoi(operands[0])
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			if testValueSearch(testValue, 1, firstValue, operands) {
				c <- testValue
			}
		}()

	}

	go func() {
		wg.Wait()
		close(c)
	}()

	sum := 0
	for n := range c {
		sum += n
	}
	fmt.Println(sum)
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	cc := strconv.Itoa(a) + strconv.Itoa(b)
	result, err := strconv.Atoi(cc)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

type operation func(int, int) int

func testValueSearch(testValue, currentIndex, accumulated int, operands []string) bool {
	if currentIndex == len(operands) {
		return false
	}
	currentValue, err := strconv.Atoi(operands[currentIndex])
	if err != nil {
		log.Fatal(err)
	}
	for _, op := range []operation{sum, multiply, concat} {
		newValue := op(accumulated, currentValue)
		if newValue == testValue && currentIndex == len(operands)-1 {
			return true
		}
		if testValueSearch(testValue, currentIndex+1, newValue, operands) {
			return true
		}
	}
	return false
}
