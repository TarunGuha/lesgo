package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	// "strconv"
	"strings"
)

func mapper(line string) (string, []string) {
	var key string
	var value []string
	value_complete_string_spilt := (strings.Split(line, "="))
	key = strings.TrimSpace(value_complete_string_spilt[0])

	values := (strings.TrimSpace(value_complete_string_spilt[1]))[1 : len(value_complete_string_spilt[1])-2]
	values_split := strings.Split(values, ",")
	value = append(value, strings.TrimSpace(values_split[0]))
	value = append(value, strings.TrimSpace(values_split[1]))
	return key, value
}

func step_pattern(line string) []int {
	var steps []int
	for _, value := range line {
		if string(value) == "L" {
			steps = append(steps, 0)
		}
		if string(value) == "R" {
			steps = append(steps, 1)
		}
	}
	return steps
}

func main() {

	readFile, err := os.Open("path.txt")

	if err != nil {
		fmt.Println(err)
	}

	var line_counter int = 0
	var steps []int
	mappings := make(map[string][]string)
	var step_counter int = 0
	var current_node = "AAA"

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		line_counter = line_counter + 1

		if line_counter == 1 {
			steps = step_pattern(str)
			continue
		}

		if str == "" {
			continue
		}

		key, value := mapper(str)
		mappings[key] = value

	}
	readFile.Close()

	for ;; {
		for _,next_move := range steps {
			if current_node == "ZZZ" {
				break
			}
			current_node = mappings[current_node][next_move]
			step_counter = step_counter + 1
		}
		if current_node == "ZZZ" {
			break
		}
	}

	fmt.Println(step_counter)
}
