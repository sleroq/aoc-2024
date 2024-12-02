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

                if last - num > 3 {
                    continue outer
                }
            } else {
                if last > num {
                    continue outer
                }

                if num - last > 3 {
                    continue outer
                }
            }

            last = num
        }

        safeCount++
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
            // slices.Reverse(data[i])
        }
	}

	// solve
	fmt.Println(solve(data))
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
