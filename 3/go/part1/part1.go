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

	m, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	d, _ := regexp.Compile(`[0-9]+`)

	muls := m.FindAllString(string(input), -1)

	total := 0

	for _, mul := range muls {
		digits := d.FindAllString(string(mul), -1)
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

	log.Println(total)
}
