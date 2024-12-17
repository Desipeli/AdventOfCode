package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	Row, Col   int
	Vrow, Vcol int
}

func (r *Robot) Move(seconds, width, height int) {
	r.Row = (r.Row + r.Vrow*seconds) % height
	r.Col = (r.Col + r.Vcol*seconds) % width

	if r.Row < 0 {
		r.Row += height
	}
	if r.Col < 0 {
		r.Col += width
	}
}

func NewRobot(row, col, vRow, vCol string) (robot *Robot, err error) {
	r, err := strconv.Atoi(row)
	if err != nil {
		return robot, err
	}
	c, err := strconv.Atoi(col)
	if err != nil {
		return robot, err
	}
	vr, err := strconv.Atoi(vRow)
	if err != nil {
		return robot, err
	}
	vc, err := strconv.Atoi(vCol)
	if err != nil {
		return robot, err
	}

	return &Robot{
		r, c, vr, vc,
	}, nil
}

func main() {
	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(inputBytes), "\n")

	r, _ := regexp.Compile("-?([0-9]+)")

	quadrants := [4]int{0, 0, 0, 0}
	roomWidth := 101
	roomHeigh := 103

	for _, row := range input {
		numbers := r.FindAllString(row, -1) // pos x, pos y, vel x, vel y
		robot, err := NewRobot(numbers[1], numbers[0], numbers[3], numbers[2])
		if err != nil {
			log.Fatal(err)
		}

		robot.Move(100, roomWidth, roomHeigh)
		if robot.Row < roomHeigh/2 {
			if robot.Col < roomWidth/2 {
				quadrants[0] += 1
			}
			if robot.Col > roomWidth/2 {
				quadrants[1] += 1
			}
		}
		if robot.Row > roomHeigh/2 {
			if robot.Col < roomWidth/2 {
				quadrants[2] += 1
			}
			if robot.Col > roomWidth/2 {
				quadrants[3] += 1
			}
		}
		// fmt.Println(robot.Row, robot.Col)
	}

	fmt.Println(quadrants)
	safetyFactor := quadrants[0]
	for _, v := range quadrants[1:] {
		safetyFactor *= v
	}

	fmt.Println(safetyFactor)
}
