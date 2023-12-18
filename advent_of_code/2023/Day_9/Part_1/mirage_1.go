package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
)

func mapper(line string) []int {
	var values []int
	value_strings := strings.Split(strings.TrimSpace(line), " ")
	for _, string_value := range value_strings {
		int_value, _ := strconv.Atoi(string_value)
		values = append(values, int_value)
	}
	return values
}

func extractor(values []int) []int {
	var new_values []int
	length := len(values)
	for index := 1; index < length; index++ {
		difference := values[index] - values[index-1]
		new_values = append(new_values, difference)
	}
	return new_values
}

func zero_checker(list []int) bool {
	for _, num := range list {
		if num != 0 {
			return false
		}
	}
	return true
}

func value_finder(line string) int {
	values := mapper(line)
	sum := values[len(values)-1]

	for {
		if zero_checker(values) {
			break
		}
		values = extractor(values)
		sum = sum + values[len(values)-1]
	}

	return sum
}

func main() {

	readFile, err := os.Open("oasis.txt")

	if err != nil {
		fmt.Println(err)
	}

	var sum int = 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		value := value_finder(str)
		sum = sum + value
	}
	readFile.Close()
	fmt.Println(sum)
}
