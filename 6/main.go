package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	data := [][]string{}
	for _, row := range strings.Split(string(input), "\n") {
		if row == "" {
			break
		}

		data = append(data, strings.Split(row, ""))
	}

	fmt.Println(solve1(data))
}

func solve1(data [][]string) (int) {
    visited := map[string]bool{}

gameLoop:
	for true {
		for y, row := range data {
			for x, pixel := range row {
				switch pixel {
				case "^":
					{
						if y > 0 {
							if data[y-1][x] == "." {
								data[y-1][x] = "^"
								data[y][x] = "."
							} else if x < len(data[y])-1 {
								// assumming no corners
								data[y][x+1] = ">"
								data[y][x] = "."
							} else {
								break gameLoop
							}
						} else {
							break gameLoop
						}
					}
				case "<":
					{
						if x > 0 {
							if data[y][x-1] == "." {
								data[y][x-1] = "<"
								data[y][x] = "."
							} else if y > 0 {
								// assumming no corners
								data[y-1][x] = "^"
								data[y][x] = "."
							} else {
								break gameLoop
							}
						} else {
							break gameLoop
						}
					}
				case "v":
					{
						if y < len(data)-1 {
							if data[y+1][x] == "." {
								data[y+1][x] = "v"
								data[y][x] = "."
							} else if x > 0 {
								// assumming no corners
								data[y][x-1] = "<"
								data[y][x] = "."
							} else {
								break gameLoop
							}
						} else {
							break gameLoop
						}
					}
				case ">":
					{
						if x < len(data[y])-1 {
							if data[y][x+1] == "." {
								data[y][x+1] = ">"
								data[y][x] = "."
							} else if y < len(data)-1 {
								// assumming no corners
								data[y+1][x] = "v"
								data[y][x] = "."
							} else {
								break gameLoop
							}
						} else {
							break gameLoop
						}
					}
				default:
					continue
				}

                visited[fmt.Sprintf("%d-%d", y, x)] = true
				continue gameLoop
			}
		}
	}

	return len(visited) + 1
}
