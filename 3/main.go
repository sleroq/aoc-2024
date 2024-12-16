package main

import (
	"fmt"
	"os"
	"strconv"
)

type Walker struct {
	value string
	pos   int
}

func (w *Walker) Next() string {
	if w.pos >= len(w.value) {
		return ""
	}

	c := w.value[w.pos]
	w.pos++

	return string(c)
}

func (w *Walker) Current() string {
	if w.pos >= len(w.value) {
		return ""
	}

	return string(w.value[w.pos])
}

func (w *Walker) BackPos(offset int) {
	w.pos -= offset
}

func (w *Walker) Try(s string) bool {
	for i, c := range s {
		if w.Current() != string(c) {
			w.BackPos(i)
			return false
		}

		w.Next()
	}

	return true
}

func solve2(w Walker) int {
	sum := 0
	enabled := false

	for w.Current() != "" {
		if w.Try("mul(") {
			first := ""

			for w.Current() != "," && len(first) < 4 {
				first += w.Current()
				w.Next()
			}

			firstInt, err := strconv.Atoi(first)
			if err != nil {
				w.BackPos(len(first))
				continue
			}

			// Skip comma
			w.Next()

			second := ""
			for w.Current() != ")" && len(second) < 4 {
				second += w.Current()
				w.Next()
			}

			secondInt, err := strconv.Atoi(second)
			if err != nil {
				w.BackPos(len(second))
				continue
			}

			if !enabled {
				continue
			}

			sum += firstInt * secondInt
		} else if w.Try("do()") {
			enabled = true
		} else if w.Try("don't()") {
			enabled = false
		} else {
			w.Next()
		}
	}

	return sum
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	walker := Walker{value: string(input)}
	walker2 := Walker{value: string(input)}

	fmt.Println(solve(walker))
	fmt.Println(solve2(walker2))
}

func solve(w Walker) int {
	sum := 0

	for w.Current() != "" {
		if w.Next() == "m" &&
			w.Next() == "u" &&
			w.Next() == "l" &&
			w.Next() == "(" {

			first := ""

			for w.Current() != "," && len(first) < 4 {
				first += w.Current()
				w.Next()
			}
			w.Next()

			firstInt, err := strconv.Atoi(first)
			if err != nil {
				w.BackPos(len(first))
				continue
			}

			second := ""
			for w.Current() != ")" && len(second) < 4 {
				second += w.Current()
				w.Next()
			}

			secondInt, err := strconv.Atoi(second)
			if err != nil {
				w.BackPos(len(second))
				continue
			}

			sum += firstInt * secondInt
		}
	}

	return sum
}
