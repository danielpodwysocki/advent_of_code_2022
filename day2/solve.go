package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_round_result(my_shape string, opponent_shape string) string {
	result := "draw"
	if my_shape == "rock" && opponent_shape == "scissors" {
		result = "win"
	}
	if my_shape == "rock" && opponent_shape == "paper" {
		result = "loose"
	}
	if my_shape == "paper" && opponent_shape == "rock" {
		result = "win"
	}
	if my_shape == "paper" && opponent_shape == "scissors" {
		result = "loose"
	}
	if my_shape == "scissors" && opponent_shape == "paper" {
		result = "win"
	}
	if my_shape == "scissors" && opponent_shape == "rock" {
		result = "loose"
	}
	return result
}

func get_round_score(my_shape string, opponent_shape string) int {
	// Write your code here
	point_shape_map := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
	point_result_map := map[string]int{
		"win":   6,
		"draw":  3,
		"loose": 0,
	}

	fmt.Println(point_result_map[get_round_result(my_shape, opponent_shape)], point_shape_map[my_shape])
	return point_result_map[get_round_result(my_shape, opponent_shape)] + point_shape_map[my_shape]
}

func main() {
	input_file := "input.txt"
	file, _ := os.Open(input_file)
	scanner := bufio.NewScanner(file)
	codename_shape_map := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	winning_shapes := map[string]string{
		"rock":     "paper",
		"paper":    "scissors",
		"scissors": "rock",
	}
	loosing_shapes := map[string]string{}
	for key, value := range winning_shapes {
		loosing_shapes[value] = key
	}

	result_mappings := map[string]string{
		"X": "loose",
		"Y": "draw",
		"Z": "win",
	}
	// first part of the AOC question - we assume X Y Z are shapes
	total_score_1 := 0
	// second part of the question - this time X Y Z are codenames for whethere
	// we win draw or loose a round of rock paper scissors
	total_score_2 := 0
	var my_shape_2 string
	for scanner.Scan() {
		var round string = scanner.Text()
		my_strategy := string(round[2])
		opponent_shape_string := string(round[0])
		my_shape := codename_shape_map[my_strategy]
		opponent_shape := codename_shape_map[opponent_shape_string]

		my_required_result := result_mappings[my_strategy]
		if my_required_result == "win" {
			my_shape_2 = winning_shapes[opponent_shape]
		} else if my_required_result == "loose" {
			my_shape_2 = loosing_shapes[opponent_shape]
		} else {
			my_shape_2 = opponent_shape
		}

		total_score_1 += get_round_score(my_shape, opponent_shape)
		total_score_2 += get_round_score(my_shape_2, opponent_shape)

	}
	fmt.Println("Strategy 1 score:", total_score_1)
	fmt.Println("Strategy 2 score:", total_score_2)

}
