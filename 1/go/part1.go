package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var list1 []int
	var list2 []int
	var totalValue int

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(line) > 1 {
			values := strings.Split(string(line), "   ")
			value1, err := strconv.Atoi(values[0])
			if err != nil {
				log.Fatal(err)
			}
			value2, err := strconv.Atoi(values[1])
			if err != nil {
				log.Fatal(err)
			}
			list1 = append(list1, value1)
			list2 = append(list2, value2)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i, value := range list1 {
		totalValue += int(math.Abs(float64(value - list2[i])))
	}

	log.Println(totalValue)
}
