package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

func main() {

	inputBytes, err := os.ReadFile(string(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(inputBytes), "\n")
	antennaPositions := map[string][]position{}

	for i, row := range input {
		for j, col := range row {
			if col != '.' {
				char := string(col)
				_, ok := antennaPositions[char]
				if ok {
					antennaPositions[char] = append(antennaPositions[char], position{i, j})
				} else {
					antennaPositions[char] = []position{{i, j}}
				}
			}
		}
	}

	width := len(input[0])
	height := len(input)
	antinodes := map[position]struct{}{}

	for _, antennas := range antennaPositions {
		if len(antennas) == 1 {
			continue
		}
		for i, a := range antennas {
			antinodes[a] = struct{}{}
			for _, oa := range antennas[i+1:] {

				rowDiff := oa.row - a.row
				colDiff := oa.col - a.col

				newAntinode := position{oa.row, oa.col}
				for {
					newAntinode = position{
						row: newAntinode.row + rowDiff,
						col: newAntinode.col + colDiff,
					}
					if newAntinode.row >= height || newAntinode.col < 0 || newAntinode.col >= width {
						break
					}
					antinodes[newAntinode] = struct{}{}
				}

				newAntinode = position{a.row, a.col}
				for {
					newAntinode = position{
						row: newAntinode.row - rowDiff,
						col: newAntinode.col - colDiff,
					}
					if newAntinode.row < 0 || newAntinode.col < 0 || newAntinode.col >= width {
						break
					}
					antinodes[newAntinode] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antinodes))

}
