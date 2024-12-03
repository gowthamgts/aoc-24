package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	totalSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, s := range regex.FindAllStringSubmatch(line, -1) {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				log.Fatal("failed to convert str to int", s[1], err)
			}

			j, err := strconv.Atoi(s[2])
			if err != nil {
				log.Fatal("failed to convert str to int", s[2], err)
			}
			totalSum += i * j
		}
	}

	fmt.Println(totalSum)
}
