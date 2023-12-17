package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
)

func string_convertor(line string) int {
	var value_string string
	var final_int_value int
	value_complete_string := strings.TrimSpace(strings.Split(line, ":")[1])
	value_strings := strings.Split(value_complete_string, " ")
	for _, value := range value_strings {
		if value != "" {
			value_string = value_string + value
		}
	}
	final_int_value, _ = strconv.Atoi(value_string)
	return final_int_value
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

	var ways int
	var time int
	var distance int
	var min_limit int
	var max_limit int

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		if str[0:1] == "T" {
			time = string_convertor(str)
		}
		if str[0:1] == "D" {
			distance = string_convertor(str)
		}

	}
	readFile.Close()

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

	ways = ((max_limit - min_limit) + 1)

	fmt.Println(ways)
}
