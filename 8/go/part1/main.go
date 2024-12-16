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

	// log.Println(antennas)

	width := len(input[0])
	height := len(input)
	antinodes := map[position]struct{}{}

	for _, antennas := range antennaPositions {
		for i, a := range antennas {
			for _, oa := range antennas[i+1:] {
				upperAntinode := position{
					row: oa.row - a.row + oa.row,
					col: oa.col - (a.col - oa.col),
				}
				lowerAntinode := position{
					row: a.row - (oa.row - a.row),
					col: a.col - (oa.col - a.col),
				}

				if upperAntinode.row < height {
					if upperAntinode.col < width && upperAntinode.col >= 0 {
						antinodes[upperAntinode] = struct{}{}
					}
				}
				if lowerAntinode.row >= 0 {
					if lowerAntinode.col < width && lowerAntinode.col >= 0 {
						antinodes[lowerAntinode] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))

}
