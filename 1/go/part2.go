package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	leftSide := map[int]int{}
	rightSide := map[int]int{}
	var similarityScore int

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 1 {
			values := strings.Split(string(line), "   ")
			value1, err := strconv.Atoi(values[0])
			if err != nil {
				log.Fatal(err)
			}
			value2, err := strconv.Atoi(values[1])
			if err != nil {
				log.Fatal(err)
			}

			leftSide[value1] += 1
			rightSide[value2] += 1
		}
	}
	for id, count := range leftSide {
		similarityScore += id * count * rightSide[id]
	}

	log.Println(similarityScore)
}
