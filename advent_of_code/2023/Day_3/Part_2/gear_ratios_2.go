package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func number_extractor(line string, current_index int, salt string) (string, int) {
	number := string(line[current_index])
	start_index := current_index
	end_index := current_index

	for i := current_index - 1; i >= 0; i-- {
		if isInt(line[i]) {
			number = string(line[i]) + number
			start_index = i
		} else {
			break
		}
	}

	for i := current_index + 1; i < len(line); i++ {
		if isInt(line[i]) {
			number = number + string(line[i])
			end_index = i
		} else {
			break
		}
	}

	key := salt + "-" + strconv.Itoa(start_index) + "-" + strconv.Itoa(end_index)
	int_number, _ := strconv.Atoi(number)

	return key, int_number
}

func checkValid(line_index int, star_index int, lines []string) (int, bool) {
	values := make(map[string]int)

	start_index := star_index
	if start_index-1 >= 0 {
		start_index = start_index - 1
		if isInt(lines[line_index][start_index]) {
			key, value := number_extractor(lines[line_index], start_index, "1")
			values[key] = value
		}
	}

	end_index := star_index
	if end_index+1 < len(lines[0]) {
		end_index = end_index + 1
		if isInt(lines[line_index][end_index]) {
			key, value := number_extractor(lines[line_index], end_index, "1")
			values[key] = value
		}
	}

	if line_index-1 >= 0 {
		for i := start_index; i <= end_index; i++ {
			if isInt(lines[line_index-1][i]) {
				key, value := number_extractor(lines[line_index-1], i, "0")
				values[key] = value
			}
		}
	}

	if line_index+1 < len(lines) {
		for i := start_index; i <= end_index; i++ {
			if isInt(lines[line_index+1][i]) {
				key, value := number_extractor(lines[line_index+1], i, "2")
				values[key] = value
			}
		}
	}

	count := 0
	power := 1

	for _, value := range values {
		power = power * value
		count = count + 1
	}

	valid := false

	if count > 1 {
		valid = true
	}

	return power, valid
}

func isInt(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	} else {
		return false
	}
}

func main() {

	readFile, err := os.Open("engine_schematic.txt")

	if err != nil {
		fmt.Println(err)
	}

	sum := 0
	var lines []string

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		lines = append(lines, str)
	}
	readFile.Close()

	for line_index, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '*' {
				value, can_add := checkValid(line_index, i, lines)
				if can_add {
					sum = sum + value
				}
			}
		}
	}
	fmt.Println(sum)
}
