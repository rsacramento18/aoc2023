package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to readh file: %v", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func checkNumber(line string, count int) bool {
	return unicode.IsNumber([]rune(line)[count])
}

func checkWordAsNumber(word string, start int, end int) string {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	slice := word[start:end]
	for i := 0; i < len(numbers); i++ {
		if strings.Contains(slice, numbers[i]) {
			switch numbers[i] {
			case "one":
				return "1"
			case "two":
				return "2"
			case "three":
				return "3"
			case "four":
				return "4"
			case "five":
				return "5"
			case "six":
				return "6"
			case "seven":
				return "7"
			case "eight":
				return "8"
			case "nine":
				return "9"
			default:
				return ""
			}
		}
	}
	return ""
}

func getSum(line string, result string, count int) string {
	if len(line) == count {
		return result
	}

	if len(result) == 0 {
		wordAsNumber := checkWordAsNumber(line, 0, count+1)
		if wordAsNumber != "" {
			result = wordAsNumber
			wordAsNumber = ""
		} else if checkNumber(line, count) {
			result = string(line[count])
		}
	}
	count++
	result = getSum(line, result, count)
	count--

	if len(result) == 1 {
		wordAsNumber := checkWordAsNumber(line, count, len(line))
		if wordAsNumber != "" {
			result += wordAsNumber
		} else if checkNumber(line, count) {
			result += string(line[count])
		}
		return result
	}
	return result
}

func main() {
	lines := ReadFile("input.txt")
	finalCount := 0
	for i := 0; i < len(lines); i++ {
		var result string
		line := lines[i]
		result = getSum(line, "", 0)

		intResult, err := strconv.Atoi(result)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		finalCount += intResult
	}
	fmt.Println(finalCount)
}
