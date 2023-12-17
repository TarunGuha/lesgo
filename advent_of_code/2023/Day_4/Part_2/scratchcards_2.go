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

	sum := 0

	fileScanner := bufio.NewScanner(readFile)

	cardCounter := make(map[int]int)

	cardNumber := 0

	for fileScanner.Scan() {

		cardNumber = cardNumber + 1

		if _, found := cardCounter[cardNumber]; found {
			cardCounter[cardNumber] = cardCounter[cardNumber] + 1
		} else {
			cardCounter[cardNumber] = 1
		}

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

			for index := cardNumber + 1; index <= cardNumber+int(total_occurences); index++ {

				if _, found := cardCounter[index]; found {
					cardCounter[index] = cardCounter[index] + cardCounter[cardNumber]
				} else {
					cardCounter[index] = cardCounter[cardNumber]
				}

			}
		}
	}

	for _, value := range cardCounter {
		sum += value
	}

	readFile.Close()

	fmt.Println(sum)
}
