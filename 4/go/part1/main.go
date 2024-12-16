package main

import (
	"log"
	"os"
	"strings"
)

type Direction struct {
	row, col int
}

func main() {
	inputFile := os.Args[1]

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(inputBytes), "\n")

	directions := []Direction{
		{0, 1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	totalXmas := 0

	for rowIndex, row := range input {
		for colIndex, char := range row {
			if char != 'X' {
				continue
			}
			for _, dir := range directions {
				if rowIndex < 3 && dir.row == -1 {
					continue
				}
				if rowIndex > len(input)-4 && dir.row == 1 {
					continue
				}
				if colIndex > len(row)-4 && dir.col == 1 {
					continue
				}
				if colIndex < 3 && dir.col == -1 {
					continue
				}
				totalXmas += findXmas(input, rowIndex, colIndex, "", dir)
			}
		}
	}

	log.Println(totalXmas)
}

func findXmas(input []string, rowIndex, colIndex int, word string, dir Direction) int {
	word += string(input[rowIndex][colIndex])
	if word == "XMAS" {
		return 1
	}
	if !strings.HasPrefix("XMAS", word) {
		return 0
	}
	newRowIndex := rowIndex + dir.row
	newColIndex := colIndex + dir.col

	return findXmas(input, newRowIndex, newColIndex, word, dir)
}
