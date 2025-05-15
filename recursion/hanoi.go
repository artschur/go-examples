package main

import (
	"fmt"
)

func hanoi(n int, origem, aux, dest string) {
	if n > 0 {
		hanoi(n-1, origem, dest, aux)
		fmt.Printf("%s -> %s \n", origem, dest)
		hanoi(n-1, aux, origem, dest)
	}
}
func main() {
	hanoi(5, "a", "b", "c")
}
