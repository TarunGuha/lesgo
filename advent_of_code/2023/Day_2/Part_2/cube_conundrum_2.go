package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("games.txt")

	if err != nil {
		fmt.Println(err)
	}

	sum := 0
	count := 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		count = count + 1
		game_data := strings.Split(str, ":")[1]
		all_rounds := strings.Split(game_data, ";")

		red_count := 0
		green_count := 0
		blue_count := 0

		for _, round := range all_rounds {
			round_data := strings.Split(round, ",")

			for _, ball_details := range round_data {

				cleaned_ball_details := strings.TrimSpace(ball_details)
				ball_split_details := strings.Split(cleaned_ball_details, " ")

				ball_count, _ := strconv.Atoi(ball_split_details[0])
				ball_colour := ball_split_details[1]

				switch ball_colour {
				case "red":
					if ball_count > red_count {
						red_count = ball_count
					}
				case "green":
					if ball_count > green_count {
						green_count = ball_count
					}
				case "blue":
					if ball_count > blue_count {
						blue_count = ball_count
					}
				}
			}
		}
		power := red_count * green_count * blue_count
		sum = sum + power
	}
	readFile.Close()
	fmt.Println(sum)
}
