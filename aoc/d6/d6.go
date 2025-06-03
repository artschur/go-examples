package d6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int16
}

func parse_matrix() [][]string {
	file, err := os.Open("/Users/aschur/Desktop/examples/aoc/d6/d6.input")
	if err != nil {
		log.Println("error opening file")
		return [][]string{}
	}
	scanner := bufio.NewScanner(file)

	matrix := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		matrix_line := []string{}
		for _, char := range line {
			matrix_line = append(matrix_line, string(char))
		}
		matrix = append(matrix, matrix_line)
	}
	return matrix
}
func get_guard_pos(matrix [][]string) coord {
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == "^" {
				return coord{
					x: int16(col),
					y: int16(row),
				}
			}
		}
	}
	return coord{-1, -1}
}

func get_new_direction(dir coord) coord {
	if dir.x == 1 && dir.y == 0 {
		return coord{0, 1}
	}

	if dir.x == 0 && dir.y == 1 {
		return coord{-1, 0}
	}

	if dir.x == -1 && dir.y == 0 {
		return coord{0, -1}
	}
	return coord{1, 0}
}

func walk_matrix(matrix *[][]string, dir, guard_pos coord, walk_count *int16) {
	going_to := coord{guard_pos.x + dir.x, guard_pos.y + dir.y}
	fmt.Println(going_to)

	if going_to.x >= int16(len(*matrix)) || going_to.y >= int16(len((*matrix)[0])) || going_to.y < 0 || going_to.x < 0 {
		return
	}
	if (*matrix)[guard_pos.y][guard_pos.x] != "X" {
		(*matrix)[guard_pos.y][guard_pos.x] = "X"
		*walk_count += 1
	}
	if (*matrix)[going_to.y][going_to.x] == "#" {
		new_dir := get_new_direction(dir)
		walk_matrix(matrix, new_dir, guard_pos, walk_count)
	} else {
		walk_matrix(matrix, dir, going_to, walk_count)
	}
}

func Solution() {
	matrix := parse_matrix()
	pos := get_guard_pos(matrix)
	if pos.x == -1 || pos.y == -1 {
		log.Println("No guard position found")
		return
	}

	var walk_count int16 = 1
	walk_matrix(&matrix, coord{0, -1}, pos, &walk_count)
}
