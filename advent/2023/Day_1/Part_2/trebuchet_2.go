package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToWord(num int) string {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return words[num]
}

func convertToWordLength(num int) string {
	lengths := []string{"4", "3", "3", "5", "4", "4", "3", "5", "5", "4"}
	return lengths[num]
}

func hmatcher() map[int][]string {
	myMap := make(map[int][]string)
	for i := 0; i < 10; i++ {
		myMap[i] = []string{convertToWordLength(i), convertToWord(i)}
	}
	return myMap
}

func isInt(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	} else {
		return false
	}
}

func fmatcher(s string) int {
	smap := hmatcher()
	fmt.Println(smap)
	for i := 0; i < len(s); i++ {
		if isInt(s[i]) {
			value, _ := strconv.Atoi(string(s[i]))
			return value
		}
		for j := 0; j < 10; j++ {
			length, _ := strconv.Atoi(smap[j][0])
			if i+length < len(s) {
				flag := 1
				for k := 0; k < length; k++ {
					if s[i:][k] != smap[j][1][k] {
						flag = -1
						break
					}
				}
				if flag == 1 {
					return j
				}
			}
		}
	}
	return -1
}

func main() {

	readFile, err := os.Open("calibration_values.txt")

	if err != nil {
		fmt.Println(err)
	}

	sum := 0
	val := 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		number := ""
		str := fileScanner.Text()
		val = fmatcher(str)

		for i := 0; i < len(str); i++ {
			if isInt(str[i]) {
				number = number + string(str[i])
				break
			}
		}

		for i := (len(str) - 1); i >= 0; i-- {
			if isInt(str[i]) {
				number = number + string(str[i])
				break
			}
		}

		value, _ := strconv.Atoi(number)
		sum = sum + value
	}
	readFile.Close()

	// fmt.Println(sum)

	fmt.Println(val)
}
