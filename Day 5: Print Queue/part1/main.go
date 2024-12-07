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
		log.Fatal(err)
	}
	defer file.Close()

	ruleMap := make(map[string]map[string]struct{})
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()

		if strings.Contains(l, "|") {
			// split and add it to the map
			split := strings.Split(l, "|")

			if _, ok := ruleMap[split[0]]; ok {
				// append
				ruleMap[split[0]][split[1]] = struct{}{}
			} else {
				ruleMap[split[0]] = map[string]struct{}{split[1]: {}}
			}
		} else if strings.Contains(l, ",") {
			split := strings.Split(l, ",")

			valid := true
			for i := 0; i < len(split); i++ {
				for j := i - 1; j >= 0; j-- {
					if _, ok := ruleMap[split[i]][split[j]]; ok {
						// it should not be present here
						valid = false
						break
					}
				}

				if !valid {
					break
				}
			}

			if valid {
				n, err := strconv.Atoi(split[len(split)/2])
				if err != nil {
					log.Fatal(err)
				}
				sum += n
			}
		}
	}

	fmt.Println(sum)
}
