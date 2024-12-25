package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var totalSum = 0

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`(\d+)+`)
	for scanner.Scan() {
		l := scanner.Text()
		// 3267: 81 40 27
		p := r.FindAllString(l, -1)
		pi := make([]int, len(p))

		for idx, n := range p {
			pi[idx], _ = strconv.Atoi(n)
		}

		exp := pi[0]
		pi = pi[1:]

		var result []string
		permutate([]string{"+", "*", "||"}, []string{}, result, exp, pi)
	}

	fmt.Println(totalSum)
}

func permutate(operators []string, current []string, result []string, expected int, p []int) bool {
	if len(current) == len(p)-1 {
		// do the op
		v := p[0]

		for i := 0; i < len(current); i++ {
			switch current[i] {
			case "+":
				v += p[i+1]
			case "*":
				v *= p[i+1]
			case "||":
				v, _ = strconv.Atoi(fmt.Sprintf("%d%d", v, p[i+1]))
			}
			if v > expected {
				return false
			}
		}

		if v == expected {
			totalSum += v
			return true
		}
		return false
	}

	for _, c := range operators {
		if permutate(operators, append(current, c), result, expected, p) {
			return true
		}
	}
	return false
}
