package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(inputBytes), "\n")
	costP1 := 0
	costP2 := 0

	for i := 0; i < len(input); i += 4 {
		buttonA := []float64{}
		buttonB := []float64{}
		prize := []float64{}
		var values []float64
		values = findIntsReturnFloats(input[i])
		buttonA = append(buttonA, values...)
		values = findIntsReturnFloats(input[i+1])
		buttonB = append(buttonB, values...)
		values = findIntsReturnFloats(input[i+2])
		prize = append(prize, values...)

		// p1 := solveMatrix(buttonA, buttonB, prize, 0)
		// if p1[0] < 0 || p1[1] < 0 {
		// 	continue
		// }
		// if p1[0] != float64(int(p1[0])) || p1[1] != float64(int(p1[1])) {
		// 	continue
		// }
		// costP1 += int(p1[0])*3 + int(p1[1])

		p2 := solveMatrix(buttonA, buttonB, prize, 10000000000000)
		if p2[0] < 0 || p2[1] < 0 {
			continue
		}
		if p2[0] != float64(int(p2[0])) || p2[1] != float64(int(p2[1])) {
			continue
		}
		costP2 += int(p2[0])*3 + int(p2[1])

		// fmt.Println(solvedMatrix)
	}

	fmt.Println(costP1)
	fmt.Println(costP2)
}

func solveMatrix(buttonA, buttonB, prize []float64, addition float64) [2]float64 {
	//part2
	prize[0] = prize[0] + addition
	prize[1] = prize[1] + addition
	// Matriisi napeista
	matrix := [2][2]float64{{buttonA[0], buttonB[0]}, {buttonA[1], buttonB[1]}}
	// fmt.Println(matrix)
	// determinantti a*d - b*c
	det := matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]

	// fmt.Println(det)
	// Käänteismatriisi ilman determinanttia.
	// Lasketaan determinantti mukaan vasta lopuksi
	var invMatrix [2][2]float64
	invMatrix[0][0] = matrix[1][1]
	invMatrix[0][1] = -matrix[0][1]
	invMatrix[1][0] = -matrix[1][0]
	invMatrix[1][1] = matrix[0][0]

	resultMatrix := [2]float64{
		(invMatrix[0][0]*prize[0] + invMatrix[0][1]*prize[1]) / det,
		(invMatrix[1][0]*prize[0] + invMatrix[1][1]*prize[1]) / det,
	}

	return resultMatrix
}

func findIntsReturnFloats(s string) (floats []float64) {
	r, _ := regexp.Compile("[0-9]+")
	values := r.FindAllString(s, -1)
	for _, v := range values {
		inted, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		floats = append(floats, float64(inted))
	}

	return floats
}
