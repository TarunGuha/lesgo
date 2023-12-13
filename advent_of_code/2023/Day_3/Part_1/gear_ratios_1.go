package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkValid(number string, line_index int, start_index int, end_index int, lines []string) bool {

	if start_index-1 >= 0 {
		start_index = start_index - 1
		if lines[line_index][start_index] != '.' {
			return true
		}
	}

	if lines[line_index][end_index] != '.' && !isInt(lines[line_index][end_index]) {
		return true
	}

	if line_index-1 >= 0 {
		for i := start_index; i <= end_index; i++ {
			if lines[line_index-1][i] != '.' {
				return true
			}
		}
	}

	if line_index+1 < len(lines) {
		for i := start_index; i <= end_index; i++ {
			if lines[line_index+1][i] != '.' {
				return true
			}
		}
	}

	return false
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
		start_index := -1
		end_index := -1
		number_started := false
		number := ""

		for i := 0; i < len(line); i++ {
			if isInt(line[i]) {
				if !number_started {
					number_started = true
					start_index = i
				}
				number = number + string(line[i])

			} else {
				if number_started {
					end_index = i
					if checkValid(number, line_index, start_index, end_index, lines) {
						integer_value, _ := strconv.Atoi(number)
						sum = sum + integer_value
					} else {
					}
					number_started = false
					number = ""
					start_index = -1
					end_index = -1
				}
			}
		}
		if number_started {
			end_index = len(line) - 1
			if checkValid(number, line_index, start_index, end_index, lines) {
				integer_value, _ := strconv.Atoi(number)
				sum = sum + integer_value
			}
			number_started = false
			number = ""
			start_index = -1
			end_index = -1
		}
	}
	fmt.Println(sum)
}
