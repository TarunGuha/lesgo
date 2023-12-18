package main

import (
	"bufio"
	"fmt"
	"sort"

	// "math"
	"os"
	"strconv"
	"strings"
)

func count_occurrences(list []int, target int) int {
	count := 0
	for _, num := range list {
		if num == target {
			count++
		}
	}
	return count
}

func string_convertor(line string) (string, int) {
	var int_value int
	var string_value string
	var values []string
	complete_string := (strings.Split(line, " "))
	for _, value := range complete_string {
		if value != "" {
			values = append(values, value)
		}
	}
	string_value = values[0]
	int_value, _ = strconv.Atoi(values[1])
	return string_value, int_value
}

func get_card_class(cards string) int {

	var card_occurrences []int
	card_counts := make(map[string]int)

	for _, card := range cards {
		if card_count, present := card_counts[string(card)]; present {
			card_counts[string(card)] = card_count + 1
		} else {
			card_counts[string(card)] = 1
		}
	}

	joker_value, found := card_counts["J"]
	if found {
		for key, value := range card_counts {
			card_counts[key] = value + joker_value
		}
		delete(card_counts, "J")
	}

	for _, count := range card_counts {
		card_occurrences = append(card_occurrences, count)
	}

	if count_occurrences(card_occurrences, 5) == 1 {
		return 7
	}

	if count_occurrences(card_occurrences, 4) == 1 {
		return 6
	}

	if count_occurrences(card_occurrences, 3) == 1 && count_occurrences(card_occurrences, 2) == 1 {
		return 5
	}

	if count_occurrences(card_occurrences, 3) == 1 && count_occurrences(card_occurrences, 1) == 2 {
		return 4
	}

	if count_occurrences(card_occurrences, 2) == 2 {
		return 3
	}

	if count_occurrences(card_occurrences, 2) == 1 {
		return 2
	}

	return 1
}

func get_chosen_card(card_1_card string, card_2_card string, chosen_card string) (bool, bool) {
	var valid bool = false
	var response bool

	if card_1_card == chosen_card {
		response = true
		valid = true
	}
	if card_2_card == chosen_card {
		response = false
		valid = true
	}
	return valid, response
}

func get_higher_order(card_1 string, card_2 string) bool {
	var valid bool
	var response bool
	card_types := [...]string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

	for _, card_type := range card_types {
		valid, response = get_chosen_card(card_1, card_2, string(card_type))
		if valid {
			return response
		}
	}
	return true
}

func card_sorter(card_1 string, card_2 string) bool {
	card_1_class := get_card_class(card_1)
	card_2_class := get_card_class(card_2)

	if card_1_class > card_2_class {
		return true
	}
	if card_1_class < card_2_class {
		return false
	}

	for i, card_1_card_rune := range card_1 {
		card_1_card := string(card_1_card_rune)
		card_2_card := string(card_2[i])

		if card_1_card == card_2_card {
			continue
		}
		return get_higher_order(card_1_card, card_2_card)
	}
	return false
}

func card_sorter_wrapper(slice []string) func(i, j int) bool {
	return func(i, j int) bool {
		return card_sorter(slice[i], slice[j])
	}
}

func main() {

	readFile, err := os.Open("cards.txt")

	if err != nil {
		fmt.Println(err)
	}

	card_values := make(map[string]int)
	var cards []string
	var winnings int = 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		extracted_card, extracted_card_value := string_convertor(str)
		card_values[extracted_card] = extracted_card_value
		cards = append(cards, extracted_card)
	}
	readFile.Close()

	sort_function := card_sorter_wrapper(cards)
	sort.Slice(cards, sort_function)

	value := len(cards)
	
	for i, card := range cards {
		// fmt.Println(card, card_values[card])
		winnings = winnings + ((value - i) * card_values[card])
	}

	fmt.Println(winnings)
}
