package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	grid := map[image.Point]rune{}
	for i, field := range strings.Fields(string(input)) {
		for t, rune := range field {
			grid[image.Point{t, i}] = rune
		}
	}

	superString := func(p image.Point, length int) []string {
		eeh := []image.Point{
			{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
			{0, -1}, {1, 0}, {0, 1}, {-1, 0},
		}

		words := make([]string, len(eeh))
		for i, d := range eeh {
			for n := range length {
				words[i] += string(grid[p.Add(d.Mul(n))])
			}
		}
		return words
	}

	first, second := 0, 0
	for p := range grid {
		first += strings.Count(strings.Join(superString(p, 4), " "), "XMAS")
		second += strings.Count("AMAMASASAMAMAS", strings.Join(superString(p, 2)[4:], ""))
	}
	fmt.Println(first)
	fmt.Println(second)
}
