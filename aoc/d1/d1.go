package d1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(path string) ([]int, []int) {
	f, err := os.Open(path)
	if err != nil {
		panic("Failed opening file: " + err.Error())
	}
	defer f.Close()

	var (
		list1 []int
		list2 []int
	)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fields := strings.Split(line, "   ")
		if len(fields) < 2 {
			continue // skip lines that don't have two numbers
		}
		num1, err := strconv.Atoi(fields[0])
		if err != nil {
			panic("Not a number: " + fields[0])
		}
		num2, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("Not a number: " + fields[1])
		}

		list1, list2 = append(list1, num1), append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	return list1, list2
}

func CalculateTotalDistance(l, r []int) int {
	if len(l) != len(r) {
		panic("Lists are not the same length")
	}
	var totalDistance int
	for index := range l {
		totalDistance += abs(l[index] - r[index])
	}
	return totalDistance
}

func ExerciseTwo(l, r []int) int {
	countMap := make(map[int]int)

	for _, number := range l {
		for _, rightNumber := range r {
			if number == rightNumber {
				countMap[number] += 1
			}
		}
	}

	var sum int
	for _, number := range l {
		sum += countMap[number] * number
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
