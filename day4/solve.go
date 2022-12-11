package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sections_overlap_fully(first_begin, first_end, second_begin, second_end int) bool {
	return first_begin <= second_begin && first_end >= second_end ||
	second_begin <= first_begin && second_end >= first_end
}

func sections_overlap_partially(first_begin, first_end, second_begin, second_end int) bool {
	return !(first_begin > second_end || first_end < second_begin)
}

func main() {
	input_file := "input.txt"
	file, _ := os.Open(input_file)
	scanner := bufio.NewScanner(file)
	fully_overlapping_sections := 0
	partially_overlapping_sections := 0
	for scanner.Scan() {
		sections := scanner.Text()
		section_1 := strings.Split(sections, ",")[0]
		section_2 := strings.Split(sections, ",")[1]
		section_1_begin, _ := strconv.Atoi(strings.Split(section_1, "-")[0])
		section_1_end, _ := strconv.Atoi(strings.Split(section_1, "-")[1])
		section_2_begin, _ := strconv.Atoi(strings.Split(section_2, "-")[0])
		section_2_end, _ := strconv.Atoi(strings.Split(section_2, "-")[1])
		if sections_overlap_fully(section_1_begin, section_1_end, section_2_begin, section_2_end) {
			fully_overlapping_sections ++
		}
		if sections_overlap_partially(section_1_begin, section_1_end, section_2_begin, section_2_end) {
			partially_overlapping_sections ++
		}

	}
	fmt.Println("We have that many fully overlapping sections within elf pairs: ", fully_overlapping_sections)
	fmt.Println("We have that many partially overlapping sections within elf pairs: ", partially_overlapping_sections)
}
