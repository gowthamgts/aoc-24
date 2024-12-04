package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	matches := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// check top to bottom
			if i+3 < len(matrix) && matrix[i][j] == 'X' && matrix[i+1][j] == 'M' && matrix[i+2][j] == 'A' && matrix[i+3][j] == 'S' {
				matches++
			}

			// check bottom to up
			if i-3 >= 0 && matrix[i][j] == 'X' && matrix[i-1][j] == 'M' && matrix[i-2][j] == 'A' && matrix[i-3][j] == 'S' {
				matches++
			}

			// check left to right
			if j+3 < len(matrix[i]) && matrix[i][j] == 'X' && matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
				matches++
			}

			// check right to left
			if j-3 >= 0 && matrix[i][j] == 'X' && matrix[i][j-1] == 'M' && matrix[i][j-2] == 'A' && matrix[i][j-3] == 'S' {
				matches++
			}

			// right bottom diagonal
			if i+3 < len(matrix) && j+3 < len(matrix[i]) && matrix[i][j] == 'X' && matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
				matches++
			}

			// left bottom diagonal
			if i+3 < len(matrix) && j-3 >= 0 && matrix[i][j] == 'X' && matrix[i+1][j-1] == 'M' && matrix[i+2][j-2] == 'A' && matrix[i+3][j-3] == 'S' {
				matches++
			}

			// left top diagonal
			if i-3 >= 0 && j-3 >= 0 && matrix[i][j] == 'X' && matrix[i-1][j-1] == 'M' && matrix[i-2][j-2] == 'A' && matrix[i-3][j-3] == 'S' {
				matches++
			}

			// top right diagonal
			if i-3 >= 0 && j+3 < len(matrix[i]) && matrix[i][j] == 'X' && matrix[i-1][j+1] == 'M' && matrix[i-2][j+2] == 'A' && matrix[i-3][j+3] == 'S' {
				matches++
			}
		}
	}

	fmt.Println(matches)
}
