package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	var leftArr []int
	var rightArr []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		split := strings.Split(currentLine, "   ")

		t, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("error occurred while converting string to int", err)
		}
		leftArr = append(leftArr, t)

		t, err = strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("error occurred while converting string to int", err)
		}
		rightArr = append(rightArr, t)
	}

	rightMap := map[int]int{}
	for i := 0; i < len(rightArr); i++ {
		elem := rightArr[i]
		if _, ok := rightMap[elem]; ok {
			rightMap[elem]++
		} else {
			rightMap[elem] = 1
		}
	}

	score := 0
	for i := 0; i < len(leftArr); i++ {
		elem := leftArr[i]
		if _, ok := rightMap[elem]; ok {
			score += elem * rightMap[elem]
		}
	}

	fmt.Println(score)
}
