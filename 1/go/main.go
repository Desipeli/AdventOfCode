package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("provide part(1 or 2) and filename")
	}
	part := os.Args[1]
	filename := os.Args[2]

	switch part {
	case "1":
		part1(filename)
	case "2":
		part2(filename)
	default:
		log.Fatal("part can only be 1 or 2")
	}
}
