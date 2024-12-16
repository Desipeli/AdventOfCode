package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

type guard struct {
	position   position
	directions []position
	direction  int
	path       map[position]struct{}
}

func (g *guard) move(nextPosition position) {
	g.position = nextPosition
	g.path[nextPosition] = struct{}{}
}

func (g *guard) getNextPosition(input []string) (position, error) {
	// Return new position and nil, if next position is inside the input
	// If out of bounds, return error
	nextPosition := position{
		row: g.position.row + g.directions[g.direction].row,
		col: g.position.col + g.directions[g.direction].col,
	}
	if nextPosition.row < 0 || nextPosition.row >= len(input) ||
		nextPosition.col < 0 || nextPosition.col >= len(input[0]) {
		return nextPosition, errors.New("out of bounds")
	}
	return nextPosition, nil
}

func (g *guard) StartMoving(input []string) {
	for {
		nextPosition, err := g.getNextPosition(input)
		if err != nil {
			break
		}

		nextPositionSymbol := string(input[nextPosition.row][nextPosition.col])
		if nextPositionSymbol == "#" {
			g.turn()
			continue
		}
		g.move(nextPosition)
	}
}

func (g *guard) turn() {
	g.direction = (g.direction + 1) % len(g.directions)
}

func NewGuard(initialPosition position) guard {
	newGuard := guard{
		directions: []position{
			{row: -1, col: 0},
			{row: 0, col: 1},
			{row: 1, col: 0},
			{row: 0, col: -1},
		},
		direction: 0,
		position:  initialPosition,
		path:      make(map[position]struct{}),
	}
	newGuard.path[initialPosition] = struct{}{}
	return newGuard
}

func main() {
	inputFile := os.Args[1]
	if inputFile == "" {
		log.Fatal("provide input file")
	}

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	inputString := string(inputBytes)
	input := strings.Split(inputString, "\n")

	initialPosition := position{}
	for i, row := range input {
		for j, char := range row {
			if char == '^' {
				initialPosition.row = i
				initialPosition.col = j
			}
		}
	}

	g := NewGuard(initialPosition)

	g.StartMoving(input)

	for i, row := range input {
		fmt.Println()
		for j, col := range row {
			_, ok := g.path[position{i, j}]
			if ok {
				fmt.Print("X")
			} else {
				fmt.Print(string(col))
			}
		}
	}

	log.Println(len(g.path))
}
