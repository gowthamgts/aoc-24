package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	sum := 0
	for i := 0; i < len(leftArr); i++ {
		sum += int(math.Abs(float64(leftArr[i]) - float64(rightArr[i])))
	}

	fmt.Println(sum)
}
