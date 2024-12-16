package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("invalid input file arg")
	}

	input := string(inputBytes)

	inputs := strings.Split(input, "\n\n")
	inputRules, updates := inputs[0], inputs[1]

	// right: left
	rules := map[string][]string{}

	middleSum := 0
	incorrectSum := 0

	for _, rule := range strings.Split(inputRules, "\n") {
		r := strings.Split(rule, "|")
		left := r[0]
		right := r[1]
		rules[right] = append(rules[right], left)
	}

	for _, update := range strings.Split(updates, "\n") {
		pages := strings.Split(update, ",")
		correct := true
		for i, page := range pages {
			prevs := rules[page]
			for _, jPage := range pages[i:] {
				if slices.Contains(prevs, jPage) {
					correct = false
					break
				}
			}
		}
		if correct {
			middleInt, err := strconv.Atoi(pages[len(pages)/2])
			if err != nil {
				log.Fatal(err)
			}
			middleSum += middleInt
		} else {
			incorrectSum += getIncorrectMiddle(pages, rules)
		}
	}

	log.Println(middleSum)
	log.Println(incorrectSum)
}

func getIncorrectMiddle(pages []string, rules map[string][]string) int {
	orderedUpdate := []string{pages[0]}
	for _, page := range pages[1:] {
		splitted := false
		for j, orderedPage := range orderedUpdate {
			prevs := rules[orderedPage]
			if slices.Contains(prevs, page) {
				leftSide := orderedUpdate[:j]
				middle := []string{page}
				rightSide := orderedUpdate[j:]
				orderedUpdate = slices.Concat(leftSide, middle, rightSide)
				splitted = true
				break
			}
		}
		if !splitted {
			orderedUpdate = append(orderedUpdate, page)
		}
	}
	middleInt, err := strconv.Atoi(orderedUpdate[len(orderedUpdate)/2])
	if err != nil {
		log.Fatal(err)
	}

	return middleInt
}
