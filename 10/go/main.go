package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row, col int
}

func main() {
	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(inputBytes), "\n")

	// log.Println(input)

	var part1TotalScore int
	var part2TotalScore int

	for i, row := range input {
		for j, height := range row {
			if height == '0' {
				visitedP1 := map[position]struct{}{}
				p := position{row: i, col: j}
				part1Score := searchTrails(input, p, visitedP1)
				part2Score := searchDistinctTrails(input, p)
				// log.Println(p, score)
				part1TotalScore += part1Score
				part2TotalScore += part2Score
			}
		}
	}

	fmt.Println("part 1:", part1TotalScore)
	fmt.Println("part 2:", part2TotalScore)
}

func searchTrails(input []string, p position, visited map[position]struct{}) int {
	visited[p] = struct{}{}
	h, err := strconv.Atoi(string(input[p.row][p.col]))
	if err != nil {
		log.Fatal(err)
		// return 0
	}

	if h == 9 {
		return 1
	}

	var score int

	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		newPosition := position{row: p.row + dir[0], col: p.col + dir[1]}
		if newPosition.row < 0 || newPosition.row >= len(input) || newPosition.col < 0 || newPosition.col >= len(input[0]) {
			continue
		}
		if input[newPosition.row][newPosition.col]-input[p.row][p.col] != 1 {
			continue
		}
		_, ok := visited[newPosition]
		if ok {
			continue
		}
		score += searchTrails(input, newPosition, visited)
	}

	return score
}

func searchDistinctTrails(input []string, p position) int {
	h, err := strconv.Atoi(string(input[p.row][p.col]))
	if err != nil {
		// log.Fatal(err)
		return 0
	}

	if h == 9 {
		return 1
	}

	var score int

	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		newPosition := position{row: p.row + dir[0], col: p.col + dir[1]}
		if newPosition.row < 0 || newPosition.row >= len(input) || newPosition.col < 0 || newPosition.col >= len(input[0]) {
			continue
		}
		if input[newPosition.row][newPosition.col]-input[p.row][p.col] != 1 {
			continue
		}
		score += searchDistinctTrails(input, newPosition)

	}

	return score
}
