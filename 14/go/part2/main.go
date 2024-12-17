package main

import (
	"bufio"
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

	const roomWidth = 101
	const roomHeight = 103

	robots := []Robot{}

	for _, row := range input {
		numbers := r.FindAllString(row, -1) // pos x, pos y, vel x, vel y
		robot, err := NewRobot(numbers[1], numbers[0], numbers[3], numbers[2])
		if err != nil {
			log.Fatal(err)
		}
		robots = append(robots, *robot)
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for i := range 10000 {
		room := [roomHeight][roomWidth]string{}
		for i := 0; i < roomHeight; i++ {
			for j := 0; j < roomWidth; j++ {
				room[i][j] = "."
			}
		}
		for _, robot := range robots {
			robot.Move(i, roomWidth, roomHeight)
			room[robot.Row][robot.Col] = "X"
		}

		_, err = writer.WriteString(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		for _, row := range room {
			for _, col := range row {
				_, err = writer.WriteString(col)
				if err != nil {
					log.Fatal(err)
				}
			}
			_, err = writer.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = writer.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}

	}
}
