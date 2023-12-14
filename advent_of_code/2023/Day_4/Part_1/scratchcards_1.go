package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("cards.txt")

	if err != nil {
		fmt.Println(err)
	}

	sum := float64(0)

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		card_data := strings.Split(str, ":")[1]
		owned_cards_string := strings.TrimSpace(strings.Split(card_data, "|")[0])
		need_cards_string := strings.TrimSpace(strings.Split(card_data, "|")[1])
		owned_cards_string_array := strings.Split(owned_cards_string, " ")
		need_cards_string_array := strings.Split(need_cards_string, " ")

		owned_cards_map := map[string]int{}
		for _, owned_card := range owned_cards_string_array {
			if owned_card == "" {
				continue
			}
			_, exists := owned_cards_map[owned_card]
			if exists {
				owned_cards_map[owned_card] = owned_cards_map[owned_card] + 1
			} else {
				owned_cards_map[owned_card] = 1
			}
		}

		need_cards_map := map[string]int{}
		for _, need_card := range need_cards_string_array {
			if need_card == "" {
				continue
			}
			_, exists := need_cards_map[need_card]
			if exists {
				need_cards_map[need_card] = need_cards_map[need_card] + 1
			} else {
				need_cards_map[need_card] = 1
			}
		}

		total_occurences := float64(0)
		for key := range owned_cards_map {
			if _, found := need_cards_map[key]; found {
				occurences := math.Min(float64(owned_cards_map[key]), float64(need_cards_map[key]))
				total_occurences = total_occurences + occurences
			}
		}
		if total_occurences != 0 {
			sum = sum + math.Pow(2, float64(total_occurences-1))
		}
	}
	readFile.Close()

	fmt.Println(sum)
}
