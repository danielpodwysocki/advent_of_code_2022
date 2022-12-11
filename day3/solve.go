package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func get_priority(item string) int {
	priority := 0
	alphabet_lower := "abcdefghijklmnopqrstuvwxyz"
	alphabet_upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabet := ""
	if unicode.IsUpper(rune(item[0])) {
		alphabet = alphabet_upper
		priority = 26
	} else {
		alphabet = alphabet_lower
	}

	for i, letter := range alphabet {
		if string(letter) == item {
			priority += i + 1
			break
		}
	}

	if priority == 0 || alphabet == alphabet_upper && priority == 26 {
		panic("Invalid item")
	}

	return priority
}

func find_badge_candidates(rucksack1 string, rucksack2 string) string {
	badge_candidates := ""
	for i := 0; i < len(rucksack1); i++ {
		if strings.Contains(badge_candidates, string(rucksack1[i])) {
			continue
		}
		if strings.Contains(rucksack2, string(rucksack1[i])) {
			badge_candidates += string(rucksack1[i])
		}
	}
	return badge_candidates
}

func find_mistake(comp1 string, comp2 string) string {
	var mistake string = ""
	for i := 0; i < len(comp1); i++ {
		for j := 0; j < len(comp2); j++ {
			if string(comp1[i]) == string(comp2[j]) {
				mistake = string(comp1[i])
			}
		}
	}
	if mistake == "" {
		panic("No mistake found in: " + comp1 + " and " + comp2)
	}

	return mistake
}

func main() {
	input_file := "input.txt"
	file, _ := os.Open(input_file)
	scanner := bufio.NewScanner(file)
	mistakes_priority_sum := 0
	badge_priority_sum := 0
	rucksack_number := 0
	badge_candidates := ""

	for scanner.Scan() {
		rucksack_number += 1
		rucksack := scanner.Text()
		comp1 := string(rucksack[0 : len(rucksack)/2])
		comp2 := string(rucksack[len(rucksack)/2:])
		mistake := find_mistake(comp1, comp2)
		mistakes_priority_sum += get_priority(mistake)

		if badge_candidates == "" {
			badge_candidates = rucksack
		} else {
			badge_candidates = find_badge_candidates(badge_candidates, rucksack)
		}
		if rucksack_number == 3 {
			var badge string = string(badge_candidates[0])
			badge_priority_sum += get_priority(badge)
			badge_candidates = ""
			rucksack_number = 0
		}
	}
	fmt.Println("Sum of prios of the mistakes the elf made:", mistakes_priority_sum)
	fmt.Println("Sum of prios of the the badges:", badge_priority_sum)
}
