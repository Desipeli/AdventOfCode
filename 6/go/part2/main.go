package main

import (
	"errors"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
)

type position struct {
	row int
	col int
}

type guard struct {
	position   position
	directions []position
	direction  int
	path       map[position][]position
}

func (g *guard) move(nextPosition position) {
	g.path[g.position] = append(g.path[g.position], nextPosition)
	g.position = nextPosition
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

func (g *guard) duplicateMove(nextPosition position) bool {
	return slices.Contains(g.path[g.position], nextPosition)
}

func (g *guard) StartMoving(input []string) int {
	for {
		nextPosition, err := g.getNextPosition(input)
		if err != nil {
			break
		}
		if g.duplicateMove(nextPosition) {
			return 1
		}
		nextPositionSymbol := string(input[nextPosition.row][nextPosition.col])
		if nextPositionSymbol == "#" {
			g.turn()
			continue
		}
		g.move(nextPosition)
	}
	return 0
}

func (g *guard) turn() {
	g.direction = (g.direction + 1) % len(g.directions)
}

func NewGuard(initialPosition position) *guard {
	return &guard{
		directions: []position{
			{row: -1, col: 0},
			{row: 0, col: 1},
			{row: 1, col: 0},
			{row: 0, col: -1},
		},
		direction: 0,
		position:  initialPosition,
		path:      make(map[position][]position),
	}
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

	originalGuard := NewGuard(initialPosition)
	originalGuard.StartMoving(input)
	var nextPositions = make(map[position]struct{})
	for _, paths := range originalGuard.path {
		for _, path := range paths {
			nextPositions[path] = struct{}{}
		}
	}

	totalLoops := make(chan int)
	var wg sync.WaitGroup

	for pos := range nextPositions {
		wg.Add(1)
		go func(p position) {
			defer wg.Done()
			inputCopy := slices.Clone(input)
			if check(inputCopy, p, initialPosition) {
				totalLoops <- 1
			}
		}(pos)
	}

	go func() {
		wg.Wait()
		close(totalLoops)
	}()

	sum := 0
	for n := range totalLoops {
		sum += n
	}
	log.Println(sum)
}

func check(input []string, pos, initialPosition position) bool {
	row := input[pos.row]
	newRow := row[:pos.col] + "#" + row[pos.col+1:]
	input[pos.row] = newRow

	g := NewGuard(initialPosition)
	res := g.StartMoving(input)
	return res == 1
}
