package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	directions := map[string][2]int{}
	directions[">"] = [2]int{0, 1}
	directions["v"] = [2]int{1, 0}
	directions["<"] = [2]int{0, -1}
	directions["^"] = [2]int{-1, 0}

	inputParts := strings.Split(string(inputBytes), "\n\n")
	warehouseRows := strings.Split(inputParts[0], "\n")
	warehouse := parseWarehouse(warehouseRows)
	for _, row := range warehouse {
		fmt.Println(row)
	}
	robot := getRobotStart(warehouse)

	instructions := strings.Split(inputParts[1], "\n")

	for _, row := range instructions {
		for _, col := range row {
			attemptToMove(&robot, directions, string(col), warehouse)
		}

	}

	fmt.Println(robot[0], robot[1])
	sumOfCoordinates := 0
	for i, row := range warehouse {
		for j, col := range row {
			if robot[0] == i && robot[1] == j {
				fmt.Print("@")
			} else {
				fmt.Print(col)
			}
			if string(col) == "[" {
				sumOfCoordinates += 100*i + j
			}
		}
		fmt.Println()
	}
	fmt.Println(sumOfCoordinates)
}

func attemptToMove(robot *[2]int, directions map[string][2]int, char string, warehouse [][]string) {
	dir := directions[string(char)]
	nextPos := [2]int{dir[0] + robot[0], dir[1] + robot[1]}
	nextPosChar := warehouse[nextPos[0]][nextPos[1]]
	if nextPosChar == "." {
		robot[0] = nextPos[0]
		robot[1] = nextPos[1]
		return
	}
	if nextPosChar == "[" || nextPosChar == "]" {
		fmt.Println("työnnä")
		success := attemptToPush(warehouse, dir, nextPos)
		if success {
			robot[0] = nextPos[0]
			robot[1] = nextPos[1]
		}
		fmt.Println(success)
	}
}

func attemptToPush(warehouse [][]string, dir [2]int, pos [2]int) bool {
	if dir[1] == 1 || dir[1] == -1 {
		return pushHorizontally(warehouse, dir, pos)
	}
	if canPushVertically(warehouse, dir, pos) {
		pushVertically(warehouse, dir, pos, "")
		fmt.Println(pos)
		if warehouse[pos[0]][pos[1]] == "[" {
			warehouse[pos[0]][pos[1]] = "."
			warehouse[pos[0]][pos[1]+1] = "."

		}
		if warehouse[pos[0]][pos[1]] == "]" {
			warehouse[pos[0]][pos[1]] = "."
			warehouse[pos[0]][pos[1]-1] = "."
		}
		return true
	}
	return false
}

func pushVertically(warehouse [][]string, dir, pos [2]int, prev string) {
	if warehouse[pos[0]][pos[1]] == "#" {
		return
	}
	if warehouse[pos[0]][pos[1]] == "]" {
		pushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1]}, "]")
		pushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1] - 1}, "[")
		warehouse[pos[0]][pos[1]] = "."
		warehouse[pos[0]][pos[1]-1] = "."
	}
	if warehouse[pos[0]][pos[1]] == "[" {
		pushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1] + 1}, "]")
		pushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1]}, "[")
		warehouse[pos[0]][pos[1]] = "."
		warehouse[pos[0]][pos[1]+1] = "."
	}
	if prev == "]" || prev == "[" {
		warehouse[pos[0]][pos[1]] = prev
	}
}

func canPushVertically(warehouse [][]string, dir, pos [2]int) bool {
	if warehouse[pos[0]][pos[1]] == "]" {
		pushRight := canPushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1]})
		pushLeft := canPushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1] - 1})
		return pushLeft && pushRight
	}
	if warehouse[pos[0]][pos[1]] == "[" {
		pushRight := canPushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1] + 1})
		pushLeft := canPushVertically(warehouse, dir, [2]int{pos[0] + dir[0], pos[1]})
		return pushLeft && pushRight
	}
	if warehouse[pos[0]][pos[1]] == "." {
		return true
	}
	return false
}

func pushHorizontally(warehouse [][]string, dir, pos [2]int) bool {
	k := 1
	for {
		newPos := [2]int{}
		newPos[0] = pos[0] + dir[0]*k
		newPos[1] = pos[1] + dir[1]*k
		if warehouse[newPos[0]][newPos[1]] == "#" {
			return false
		}
		if warehouse[newPos[0]][newPos[1]] == "." {
			break
		}
		k += 1
	}
	for i := k; i > 0; i-- {
		warehouse[pos[0]][pos[1]+i*dir[1]] = warehouse[pos[0]][pos[1]+(i-1)*dir[1]]
		warehouse[pos[0]][pos[1]+(i-1)*dir[1]] = "."
	}
	return true
}

func parseWarehouse(warehouseRows []string) [][]string {
	warehouse := [][]string{}
	for _, row := range warehouseRows {
		newRow := []string{}
		for _, char := range row {
			if char == '#' {
				newRow = append(newRow, "#")
				newRow = append(newRow, "#")
			}
			if char == 'O' {
				newRow = append(newRow, "[")
				newRow = append(newRow, "]")
			}
			if char == '.' {
				newRow = append(newRow, ".")
				newRow = append(newRow, ".")
			}
			if char == '@' {
				newRow = append(newRow, "@")
				newRow = append(newRow, ".")
			}
		}
		warehouse = append(warehouse, newRow)
	}

	return warehouse
}

func getRobotStart(warehouse [][]string) [2]int {
	robot := [2]int{} // row, col

	for i, row := range warehouse {
		for j, col := range row {
			if col == "@" {
				robot[0] = i
				robot[1] = j
				warehouse[i][j] = "."
			}
		}
	}
	return robot
}
