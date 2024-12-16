package main

import (
	"fmt"
	"log"
	"os"
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

	memo := map[position]struct{}{}
	memo2 := map[position]struct{}{}

	var p1Price int
	var p2Price int

	for i, row := range input {
		for j, col := range row {
			_, alreadyChecked := memo[position{row: i, col: j}]
			if alreadyChecked {
				continue
			}
			area, perimeter := getAreaAndPerimeter(input, position{row: i, col: j}, string(col), memo)
			// log.Println("pos", i, j, "plant", string(col), "area", area, "p", perimeter, "cost", area*perimeter)
			p1Price += area * perimeter

			area, sides := getAreaAndCorners(input, position{row: i, col: j}, string(col), memo2)
			// fmt.Println("area:", area, "sides", sides)
			p2Price += area * sides
			// break
		}
		// break
	}

	fmt.Println(p1Price)
	fmt.Println(p2Price)

	// a, s := getAreaAndCorners(input, position{row: 0, col: 1}, "B", memo2)
	// fmt.Println(a, s)

}

func getAreaAndCorners(input []string, pos position, plant string, memo map[position]struct{}) (area, corners int) {
	if pos.row < 0 || pos.row >= len(input) || pos.col < 0 || pos.col >= len(input[0]) {
		return 0, 0
	}
	_, alreadyChecked := memo[pos]
	if alreadyChecked {
		return
	}
	if string(input[pos.row][pos.col]) != plant {
		return 0, 0
	}

	memo[pos] = struct{}{}

	kulmat := []position{
		{row: 0, col: 1}, {row: 1, col: 0},
		{row: 1, col: 0}, {row: 0, col: -1},
		{row: 0, col: -1}, {row: -1, col: 0},
		{row: -1, col: 0}, {row: 0, col: 1},
	}
	for i := 1; i < len(kulmat); i += 2 {
		// fmt.Println(i)
		first := position{pos.row + kulmat[i-1].row, pos.col + kulmat[i-1].col}
		second := position{pos.row + kulmat[i].row, pos.col + kulmat[i].col}
		firstOut := first.row < 0 || first.row >= len(input) || first.col < 0 || first.col >= len(input[0])
		secondOut := second.row < 0 || second.row >= len(input) || second.col < 0 || second.col >= len(input[0])
		// fmt.Println("fs", first, second, firstOut, secondOut, first.col, len(input[0]))
		if firstOut && secondOut {
			corners++
			continue
		}
		if firstOut && !secondOut {
			if string(input[second.row][second.col]) != plant {
				corners++
			}
			continue
		}
		if !firstOut && secondOut {
			if string(input[first.row][first.col]) != plant {
				corners++
			}
			continue
		}
		if string(input[first.row][first.col]) != plant && string(input[second.row][second.col]) != plant {
			corners++
		}
	}

	nurkat := []position{
		{row: 0, col: 1}, {row: 1, col: 1}, {row: 1, col: 0},
		{row: 1, col: 0}, {row: 1, col: -1}, {row: 0, col: -1},
		{row: 0, col: -1}, {row: -1, col: -1}, {row: -1, col: 0},
		{row: -1, col: 0}, {row: -1, col: 1}, {row: 0, col: 1},
	}
	for i := 2; i < len(nurkat); i += 3 {
		first := position{pos.row + nurkat[i-2].row, pos.col + nurkat[i-2].col}
		second := position{pos.row + nurkat[i-1].row, pos.col + nurkat[i-1].col}
		third := position{pos.row + nurkat[i].row, pos.col + nurkat[i].col}

		firstOut := first.row < 0 || first.row >= len(input) || first.col < 0 || first.col >= len(input[0])
		secondOut := second.row < 0 || second.row >= len(input) || second.col < 0 || second.col >= len(input[0])
		thirdOut := third.row < 0 || third.row >= len(input) || third.col < 0 || third.col >= len(input[0])
		if firstOut || secondOut || thirdOut {
			continue
		}

		if string(input[first.row][first.col]) == plant && string(input[second.row][second.col]) != plant && string(input[third.row][third.col]) == plant {
			// fmt.Println(pos, first, second, third)
			corners++
		}
	}

	for _, newPos := range []position{
		{row: pos.row, col: pos.col + 1},
		{row: pos.row + 1, col: pos.col},
		{row: pos.row, col: pos.col - 1},
		{row: pos.row - 1, col: pos.col},
	} {
		a, c := getAreaAndCorners(input, newPos, plant, memo)
		area += a
		corners += c
	}

	return area + 1, corners
}

func getAreaAndPerimeter(input []string, pos position, plant string, memo map[position]struct{}) (area, perimeter int) {

	for _, newPos := range []position{
		{row: pos.row, col: pos.col + 1},
		{row: pos.row + 1, col: pos.col},
		{row: pos.row, col: pos.col - 1},
		{row: pos.row - 1, col: pos.col},
	} {

		rowsOutOfRange := newPos.row < 0 || newPos.row >= len(input)
		colsOutOfRange := newPos.col < 0 || newPos.col >= len(input[0])
		if colsOutOfRange || rowsOutOfRange {
			perimeter += 1
			continue
		}

		differentPlant := string(input[newPos.row][newPos.col]) != plant
		if differentPlant {
			perimeter += 1
			continue
		}

		_, alreadyChecked := memo[position{row: newPos.row, col: newPos.col}]
		memo[pos] = struct{}{}

		if alreadyChecked {
			continue
		}

		a, p := getAreaAndPerimeter(input, newPos, plant, memo)
		area += a
		perimeter += p
	}

	return area + 1, perimeter
}
