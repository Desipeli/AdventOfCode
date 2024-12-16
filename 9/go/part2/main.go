package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {

	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := string(inputBytes)

	var disk [][]int
	// -1 == "."

	for i, char := range input {
		digit, err := strconv.Atoi(string(char))
		if digit == 0 {
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		index := i / 2

		if i%2 == 0 {
			fileBlock := []int{}
			for range digit {
				fileBlock = append(fileBlock, index)
			}
			disk = append(disk, fileBlock)
		} else {
			emptyBlock := []int{}
			for range digit {
				emptyBlock = append(emptyBlock, -1)
			}
			disk = append(disk, emptyBlock)
		}
	}

	log.Println(disk)

	for i := len(disk) - 1; i >= 0; i-- {
		for j, block := range disk {
			if j == i {
				break
			}
			if disk[i][0] != -1 {
				if block[0] == -1 && len(block) >= len(disk[i]) {
					leftOverSpace := block[len(disk[i]):]
					disk[j] = slices.Clone(disk[i])

					// Korvaa siirretyt -1
					for k := range disk[i] {
						disk[i][k] = -1
					}

					//
					if len(leftOverSpace) > 0 {
						// Jos siirretyn paikan perään jää tyhjiä, lisää ne.
						lo := [][]int{leftOverSpace}
						disk = slices.Concat(disk[:j+1], lo, disk[j+1:])
						i++
					}
				}
			}
		}
	}

	log.Println(disk)

	var checksum int
	index := 0
	for _, value := range disk {
		if value[0] == -1 {
			index += len(value)
			continue
		}

		for _, v := range value {
			// log.Println(i+j, v, checksum)
			checksum += v * index
			index += 1
		}
	}

	fmt.Println(checksum)
}
