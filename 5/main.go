package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	rules := [][2]int{}
	pages := [][]int{}

	pagesStart := false
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			pagesStart = true
			continue
		}

		if !pagesStart {
			parts := strings.Split(line, "|")
			rules = append(rules, [2]int{parseInt(parts[0]), parseInt(parts[1])})
		} else {
			parts := strings.Split(line, ",")
			intParts := []int{}
			for _, part := range parts {
				intParts = append(intParts, parseInt(part))
			}

			pages = append(pages, intParts)
		}
	}

	fmt.Println(solve1(rules, pages))
	fmt.Println(solve2(rules, pages))
}

func solve1(rules [][2]int, pages [][]int) (sum int) {
lineLoop:
	for _, list := range pages {
		for i, page := range list {
			// Find it in rules
			for _, pair := range rules {
				if pair[0] == page {
					for t, pageB := range list {
						// Assume no duplicates
						if pair[1] == pageB && t < i {
							continue lineLoop
						}
					}
				} else if pair[1] == page {
					for t, pageB := range list {
						// Assume no duplicates
						if pair[0] == pageB && t > i {
							continue lineLoop
						}
					}
				} else {
					continue
				}
			}
		}

		sum += list[int(math.Floor(float64(len(list))/2))]
	}

	return sum
}

func solve2(rules [][2]int, pages [][]int) (sum int) {
	for _, list := range pages {
        needCheck := true
        lineChanged := false

        for needCheck {
            anythingFixed := false

            for i, page := range list {
                smthFixed := checkAndFix(rules, list, page, i)
                if smthFixed {
                    anythingFixed = true
                    lineChanged = true
                }
            }

            if !anythingFixed {
                needCheck = false
            }
        }

        if lineChanged {
            sum += list[int(math.Floor(float64(len(list))/2))]
        }
	}

	return sum
}

func checkAndFix(rules [][2]int, list []int, page int, i int) bool {
	for _, pair := range rules {
		if pair[0] == page {
			for t, pageB := range list {
				// Assume no duplicates
				if pair[1] == pageB && t < i {
                    list[i] = pageB
                    list[t] = page
                    return true
				}
			}
		} else if pair[1] == page {
			for t, pageB := range list {
				// Assume no duplicates
				if pair[0] == pageB && t > i {
                    list[i] = pageB
                    list[t] = page
                    return true
				}
			}
		} else {
			continue
		}
	}

    return false
}

func parseInt(s string) int {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(a)
}
