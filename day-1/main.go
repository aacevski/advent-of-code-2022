package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	input := string(data)

	elves := strings.Split(input, "\n\n")

	max := 0

	for _, line := range elves {
		calories := strings.Split(line, "\n")

		total := 0

		for _, calorie := range calories {
			amount, _ := strconv.Atoi(calorie)

			total += amount
		}

		if total > max {
			max = total
		}
	}

	println(max)
}
