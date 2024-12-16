package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := string(inputBytes)

	var disk []int
	// -1 == "."

	for i, char := range input {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		index := i / 2

		if i%2 == 0 {
			for range digit {
				disk = append(disk, index)
			}
		} else {
			for range digit {
				disk = append(disk, -1)
			}
		}
	}

	// log.Println(disk)

	j := len(disk) - 1
	for i, ID := range disk {
		if i >= j {
			break
		}
		if ID == -1 {
			for {
				if i >= j {
					break
				}
				last := disk[j]
				if last != -1 {
					disk[i] = last
					disk[j] = -1
					j -= 1
					break
				}
				j -= 1
			}
		}
	}

	// log.Println(disk)

	var checksum int
	for i, value := range disk {
		if value == -1 {
			break
		}
		checksum += i * value
	}

	fmt.Println(checksum)
}
