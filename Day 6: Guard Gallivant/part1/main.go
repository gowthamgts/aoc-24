package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	guardPosX, guardPosY      int
	uniques, currentDirection int
	matrix                    [][]rune
	directions                = [][2]int{
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ra := []rune(scanner.Text())
		matrix = append(matrix, ra)

		up := slices.Index(ra, '^')
		if up >= 0 {
			guardPosX, guardPosY = len(matrix)-1, up
			matrix[len(matrix)-1][up] = 'x'
			uniques++
		}
	}

	for {
		nx, ny := guardPosX+directions[currentDirection][0], guardPosY+directions[currentDirection][1]
		if nx < 0 || nx >= len(matrix) || ny < 0 || ny >= len(matrix[0]) {
			break
		}

		if matrix[nx][ny] == '#' {
			// rotate
			currentDirection = (currentDirection + 1) % 4
		} else {
			guardPosX, guardPosY = nx, ny
			if matrix[nx][ny] == '.' {
				matrix[guardPosX][guardPosY] = 'x'
				uniques++
			}
		}
	}
	fmt.Println(uniques)
}
