package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve(data [][]int) int {
	sum := 0

	// sort cols
	slices.Sort(data[0])
	slices.Sort(data[1])

	// calc sum
	for i := 0; i < len(data[0]); i++ {
		pair := []int{data[0][i], data[1][i]}
		fmt.Println(pair, math.Abs(float64(pair[0]-pair[1])), int(math.Abs(float64(pair[0]-pair[1]))))

		sum += int(math.Abs(float64(pair[0]-pair[1])))
	}

	return sum
}

func main() {
	// read input
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// parse input
	var data [][]int

	data = append(data, []int{})
	data = append(data, []int{})

	for _, line := range strings.Split(string(input), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		res := strings.Split(line, "   ")

		data[0] = append(data[0], parseInt(res[0]))
		data[1] = append(data[1], parseInt(res[1]))
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
