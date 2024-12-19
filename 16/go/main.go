package main

import (
	"day16/priority"
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

	input := strings.Split(string(inputBytes), "\n")

	start := [2]int{}
	for i, row := range input {
		for j, col := range row {
			if col == 'S' {
				start = [2]int{i, j}
			}
		}
	}

	score, _ := bfs(input, start)
	fmt.Println(score)
}

func bfs(maze []string, start [2]int) (int, bool) {
	visited := map[[2]int]int{} // (row, col): smallest score
	visited[start] = 0
	startNode := priority.PriorityQueueNode{Key: 0, Value: [2][2]int{start, {0, 1}}}
	pq := priority.NewPriorityQueue()
	pq.Insert(startNode)
	foundEnd := false
	totalScore := 0

	for !foundEnd {
		qNode, ok := pq.Pop()
		if !ok {
			fmt.Println("No more nodes")
			break
		}
		node, ok := qNode.Value.([2][2]int)
		if !ok {
			log.Fatal("invalid value")
		}
		pos := node[0]
		prevDir := node[1]
		if string(maze[pos[0]][pos[1]]) == "#" {
			continue
		}

		score := qNode.Key

		if string(maze[pos[0]][pos[1]]) == "E" {
			return score, true
		}

		for _, dir := range [4][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		} {
			newScore := score + 1
			if dir != prevDir {
				if dir[0]+prevDir[0] == 0 && dir[1]+prevDir[1] == 0 {
					newScore += 1000
				}
				newScore += 1000
			}
			newNode := priority.PriorityQueueNode{
				Key:   newScore,
				Value: [2][2]int{{pos[0] + dir[0], pos[1] + dir[1]}, dir},
			}

			_, hasVisited := visited[[2]int{pos[0] + dir[0], pos[1] + dir[1]}]
			if hasVisited {
				continue
			}
			visited[[2]int{pos[0] + dir[0], pos[1] + dir[1]}] = newScore

			pq.Insert(newNode)
		}
	}

	return totalScore, foundEnd
}
