package d4

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput() (rules map[int][]int, updates [][]int) {
	f, err := os.Open("/Users/aschur/Desktop/examples/aoc/d4/d4.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var rulesMap = make(map[int][]int)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")

		mustAppearBefore, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		mustAppearAfter, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		if _, ok := rulesMap[mustAppearBefore]; ok {
			rulesMap[mustAppearBefore] = append(rulesMap[mustAppearBefore], mustAppearAfter)
		} else {
			rulesMap[mustAppearBefore] = []int{mustAppearAfter}
		}
	}

	for s.Scan() {
		update := s.Text()

		var batchUpdate []int
		for page := range strings.SplitSeq(update, ",") {
			pageNum, err := strconv.Atoi(page)
			if err != nil {
				panic(err)
			}
			batchUpdate = append(batchUpdate, pageNum)
		}
		updates = append(updates, batchUpdate)
	}

	return rulesMap, updates
}

func ExerciseOne() int {
	rules, updates := parseInput()

	var validUpdates [][]int
	for _, currentUpdate := range updates {

		isUpdateValid := true
		for pageIdx := range currentUpdate {
			cantAppearBefore := rules[currentUpdate[pageIdx]]
			for _, numberCantAppear := range cantAppearBefore {
				if slices.Contains(currentUpdate[:pageIdx], numberCantAppear) {
					// I go to the current page, parse the invalid numbers, and check for each of them if they are inside the left side of the slice(before the current page). if they are, they are violating the rules (which said that those numbers that MUST appear after.)
					isUpdateValid = false
					break
				}
			}
			if !isUpdateValid {
				continue
			}
		}

		if isUpdateValid {
			validUpdates = append(validUpdates, currentUpdate)
		}
	}
	fmt.Printf("validUpdates: %v\n", validUpdates)
	var middlePageSum int
	for _, update := range validUpdates {
		middlePageSum += update[(0+len(update))/2]
	}

	return middlePageSum
}
