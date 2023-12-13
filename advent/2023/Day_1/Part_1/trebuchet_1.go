package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isInt(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	} else {
		return false
	}
}

func main() {

	readFile, err := os.Open("calibration_values.txt")

	if err != nil {
		fmt.Println(err)
	}

	sum := 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		number := ""
		str := fileScanner.Text()

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

	fmt.Println(sum)
}
