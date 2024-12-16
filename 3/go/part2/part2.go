package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide part and input file")
	}
	inputFile := os.Args[1]

	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	c, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	d, _ := regexp.Compile(`[0-9]+`)

	commands := c.FindAllString(string(input), -1)
	enabled := true

	total := 0

	for _, cmd := range commands {
		log.Println("command ", cmd)
		switch cmd {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if !enabled {
				continue
			}
			digits := d.FindAllString(string(cmd), -1)
			first, err := strconv.Atoi(digits[0])
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.Atoi(digits[1])
			if err != nil {
				log.Fatal(err)
			}
			total += first * second
		}
	}
	log.Println(total)
}
