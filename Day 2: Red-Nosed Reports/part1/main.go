package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		str := strings.Fields(scanner.Text())

		report := make([]int, len(str))
		for idx, s := range str {
			ci, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("failed to parse string to int", s)
			}
			report[idx] = ci
		}

		fmt.Println(report)

		if isSafeReport(report) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isSafeReport(report []int) bool {
	// convert str to int
	var order rune = 0

	for i := 0; i < len(report); i++ {
		if i > 0 {
			currentDiff := report[i] - report[i-1]
			if currentDiff == 0 {
				return false
			}

			if order == 0 {
				if currentDiff < 0 {
					order = 'd'
				} else {
					order = 'i'
				}
			}

			if math.Abs(float64(currentDiff)) > 3 || (order == 'd' && currentDiff > 0) || (order == 'i' && currentDiff < 0) {
				return false
			}
		}
	}

	return true
}
