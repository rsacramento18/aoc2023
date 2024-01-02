package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part_1(false)
	part_1(true)
}

func part_1(part2 bool) {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalf("Error opening the file %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	keyValues := []int{}
	results := 0

	for scanner.Scan() {
		sequence := convertLine(processLine(scanner.Text()))
		newSequence := []int{}
		sequenceZero := false
		for !sequenceZero {
			sequenceZero = true
			for i := 0; i < len(sequence)-1; i++ {
				value := sequence[i+1] - sequence[i]
				newSequence = append(newSequence, value)
				if value != 0 {
					sequenceZero = false
				}
			}
			if part2 {
				keyValues = append(keyValues, sequence[0])
			} else {
				keyValues = append(keyValues, sequence[len(sequence)-1])
			}
			sequence = newSequence
			newSequence = nil
		}
		keyValues = append(keyValues, 0)
		sum := 0
		for i := len(keyValues) - 1; i >= 0; i-- {
			if part2 {
				sum = keyValues[i] - sum
			} else {
				sum = sum + keyValues[i]
			}
		}
		keyValues = nil
		results = results + sum
	}
	fmt.Println(results)
}

func processLine(line string) []string {
	regex := regexp.MustCompile(" ")
	return regex.Split(line, -1)
}

func convertLine(line []string) []int {
	converted := []int{}
	for _, i := range line {
		value, _ := strconv.Atoi(i)
		converted = append(converted, value)
	}
	return converted
}
