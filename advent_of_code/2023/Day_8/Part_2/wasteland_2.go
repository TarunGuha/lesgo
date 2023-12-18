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

func check_completion(current_nodes []string) bool {
	for _, node := range current_nodes {
		if node[2:3] != "Z" {
			return false
		}
	}
	return true
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
	var current_nodes []string
	var all_done bool = false

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

	for key, _ := range mappings {
		if key[2:3] == "A" {
			current_nodes = append(current_nodes, key)
		}
	}

	for {
		for _, next_move := range steps {
			all_done = check_completion(current_nodes)
			if all_done {
				break
			}
			for node_index, node_value := range current_nodes {
				current_nodes[node_index] = mappings[node_value][next_move]
			}
			step_counter = step_counter + 1
			if (step_counter%1000000) == 0 {
				fmt.Println(step_counter)
			}
		}
		if all_done {
			break
		}
	}

	fmt.Println(step_counter)
}
