package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const input string = "input.txt"
	fmt.Println("Hello, World!")
	var current_total_calories int = 0
	var top_elves []int = []int{0, 0, 0}

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			if current_total_calories > top_elves[len(top_elves)-1] {
				// we loop through to keep it sorted
				last_calories := top_elves[0]
				// when true, move out everyone one place down the list and kick out the last one
				shift := false
				for i, calories := range top_elves {
					if shift {
						tmp_last_calories := calories
						top_elves[i] = last_calories
						last_calories = tmp_last_calories
						continue
					}
					if current_total_calories > top_elves[i] {
						last_calories = calories
						top_elves[i] = current_total_calories
						shift = true
					}
				}
			}
			current_total_calories = 0
			continue
		}
		// convert line contents to int
		// add to current_total_calories
		line_calories, _ := strconv.Atoi(line)
		current_total_calories += line_calories

	}
	top_elves_sum := 0
	for _, top_elf := range top_elves {
		top_elves_sum += top_elf
	}
	fmt.Println("Elf with the most total calories: ", top_elves[0])
	fmt.Println("Sum of top 3 elves: ", top_elves_sum)

}
