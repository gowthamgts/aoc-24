package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	matrix     [][]rune
	directions = [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		loops                int
		guardPosX, guardPosY int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ra := []rune(scanner.Text())
		matrix = append(matrix, ra)

		up := slices.Index(ra, '^')
		if up >= 0 {
			guardPosX, guardPosY = len(matrix)-1, up
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '#' {
				continue
			}

			if i == guardPosX && j == guardPosY {
				continue
			}

			if isLoop(i, j, guardPosX, guardPosY, 0) {
				loops++
			}
		}
	}

	fmt.Println(loops)
}

func isLoop(obsX, obsY, guardPosX, guardPosY, currentDirection int) bool {
	localVisitedMap := make(map[string]struct{})
	localVisitedMap[fmt.Sprintf("%d,%d,%d", guardPosX, guardPosY, currentDirection)] = struct{}{}

	for {
		nx, ny := guardPosX+directions[currentDirection][0], guardPosY+directions[currentDirection][1]
		if nx < 0 || nx >= len(matrix) || ny < 0 || ny >= len(matrix[0]) {
			break
		}

		if matrix[nx][ny] == '#' || (nx == obsX && ny == obsY) {
			// rotate
			currentDirection = (currentDirection + 1) % 4
		} else {
			if _, ok := localVisitedMap[fmt.Sprintf("%d,%d,%d", nx, ny, currentDirection)]; ok {
				return true
			}
			guardPosX, guardPosY = nx, ny
			localVisitedMap[fmt.Sprintf("%d,%d,%d", guardPosX, guardPosY, currentDirection)] = struct{}{}
		}
	}
	return false
}
