package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
)

func string_convertor(line string) []int {
	var values []int
	value_complete_string := strings.TrimSpace(strings.Split(line, ":")[1])
	value_strings := strings.Split(value_complete_string, " ")
	for _, value := range value_strings {
		if value != "" {
			value_int, _ := strconv.Atoi(value)
			values = append(values, value_int)
		}
	}
	return values
}

func can_cross(speed int, time int, distance int) bool {
	if distance < speed*time {
		return true
	} else {
		return false
	}
}

func main() {

	readFile, err := os.Open("stats.txt")

	if err != nil {
		fmt.Println(err)
	}

	var ways int = 1
	var times []int
	var distances []int

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		if str[0:1] == "T" {
			times = string_convertor(str)
		}
		if str[0:1] == "D" {
			distances = string_convertor(str)
		}

	}
	readFile.Close()

	for index, time := range times {
		distance := distances[index]
		var min_limit int
		var max_limit int

		for speed := 1; speed < time; speed++ {
			if can_cross(speed, time-speed, distance) {
				min_limit = speed
				break
			}
		}

		for speed := time - 1; speed > 0; speed-- {
			if can_cross(speed, time-speed, distance) {
				max_limit = speed
				break
			}
		}

		ways = ways * ((max_limit - min_limit) + 1)
	}
	fmt.Println(ways)
}
