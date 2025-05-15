package d2

import (
	"bufio"
	"examples/aoc/math"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	f, err := os.Open("/Users/aschur/Desktop/examples/aoc/d2/d2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var list [][]int
	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, " ")
		convertedList := []int{}
		for _, l := range chars {
			num, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}
			convertedList = append(convertedList, num)
		}
		list = append(list, convertedList)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}

func ExerciseOne() int {
	list := parseInput()
	var isSafeCount int
	for _, line := range list {
		isAsc := line[0] < line[1]
		isSafe := true
		for index := 1; index < len(line); index++ {
			if line[index-1] > line[index] && isAsc {
				isSafe = false
			}
			if line[index-1] < line[index] && !isAsc {
				isSafe = false
			}

			if math.Abs(line[index-1]-line[index]) > 3 {
				isSafe = false
			}
		}
		if isSafe {
			isSafeCount++
		}

	}
	return isSafeCount
}
