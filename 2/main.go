package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(data [][]int) int {
	var safeCount int

outer:
	for _, report := range data {
		last := report[0]
		dec := last > report[1]

		if last == report[1] {
			continue
		}

		for _, num := range report[1:] {
			if num == last {
				continue outer
			}

			if dec {
				if last < num {
					continue outer
				}

				if last-num > 3 {
					continue outer
				}
			} else {
				if last > num {
					continue outer
				}

				if num-last > 3 {
					continue outer
				}
			}

			last = num
		}

		safeCount++
	}

	return safeCount
}

func isSafe(report []int) bool {
	diffs := []int{}

	for i := 0; i < len(report)-1; i++ {
		diffs = append(diffs, report[i+1]-report[i])
	}

	if diffs[0] > 0 {
		for _, diff := range diffs {
			if diff > 3 || diff < 1 {
				return false
			}
		}
	} else {
		for _, diff := range diffs {
			if diff < -3 || diff > -1 {
				return false
			}
		}
	}

	return true
}

func removeElement(s []int, i int) []int {
	if i >= len(s) || i < 0 {
		return s
	}
	newS := make([]int, 0, len(s)-1)
	newS = append(newS, s[:i]...)
	newS = append(newS, s[i+1:]...)
	return newS
}

func solve2(data [][]int) int {
	var safeCount int

	for _, report := range data {
		for i := range report {
			filtered := removeElement(report, i)

			if isSafe(filtered) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func main() {
	// read input
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// parse input
	var data [][]int

	for i, line := range strings.Split(string(input), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		data = append(data, []int{})
		for _, num := range strings.Split(line, " ") {
			data[i] = append(data[i], parseInt(num))
		}
	}

	// solve
	fmt.Println(solve(data))
	fmt.Println(solve2(data))
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
