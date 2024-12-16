package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("provide part and input file")
	}
	part := os.Args[1]
	inputFile := os.Args[2]

	file, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	switch part {
	case "1":
		part1(string(file))
	case "2":
		part2(string(file))
	default:
		log.Fatal("invalid arguments")
	}
}

func part1(input string) {
	safeReports := 0
	reports := strings.Split(input, "\n")
	for _, report := range reports {
		levels := strings.Split(report, " ")
		if IsSafe(levels) {
			safeReports += 1
		}
	}
	log.Println(safeReports)
}

func part2(input string) {
	safeReports := 0

	reports := strings.Split(input, "\n")

	for _, report := range reports {
		levels := strings.Split(report, " ")
		safe, index := IsSafe2(levels)
		if safe {
			safeReports += 1
		} else {
			for i := -1; i < 2; i++ {
				if index == 0 && i == -1 {
					continue
				}
				newReport := append(append([]string{}, levels[:i+index]...), levels[index+i+1:]...)
				if IsSafe(newReport) {
					safeReports += 1
					break
				}
			}
		}
	}
	log.Println(safeReports)
}
