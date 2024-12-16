package main

import (
	"log"
	"strconv"
)

func IsSafe(levels []string) bool {
	direction := 0
	for i, level := range levels[1:] {
		current, err := strconv.Atoi(level)
		if err != nil {
			log.Fatal(err)
		}
		previous, err := strconv.Atoi(levels[i])
		if err != nil {
			log.Fatal(err)
		}
		diff := current - previous
		if diff > 0 {
			if direction < 0 {
				return false
			}
			direction = diff
		} else if diff < 0 {
			if direction > 0 {
				return false
			}
			direction = diff
		}
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

func IsSafe2(levels []string) (bool, int) {
	direction := 0
	for i, level := range levels[1:] {
		current, err := strconv.Atoi(level)
		if err != nil {
			log.Fatal(err)
		}
		previous, err := strconv.Atoi(levels[i])
		if err != nil {
			log.Fatal(err)
		}
		diff := current - previous
		if diff > 0 {
			if direction < 0 {
				return false, i
			}
			direction = diff
		} else if diff < 0 {
			if direction > 0 {
				return false, i
			}
			direction = diff
		}
		if diff < -3 || diff > 3 || diff == 0 {
			return false, i
		}
	}
	return true, 0
}
