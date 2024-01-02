package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Map struct {
	left  string
	right string
}

func splitline(line string) Map {
	left := line[7:10]
	right := line[12:15]
	return Map{left, right}
}

func checkSide(current, instruction string, puzzleInput map[string]Map) string {
	if instruction == "R" {
		return puzzleInput[current].right
	} else {
		return puzzleInput[current].left
	}
}

func main() {
	// part_1()
	part_2()
}

func part_1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening the file:%v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	puzzleInput := make(map[string]Map)
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		linePuzzle := splitline(line)
		puzzleInput[line[:3]] = linePuzzle
	}

	// fmt.Printf("%v\n", puzzleInput)

	start := "AAA"
	end := "ZZZ"

	for !(start == end) {
		for _, inst := range instructions {
			start = checkSide(start, string(inst), puzzleInput)
			// fmt.Println(start)
			res++
			if start == end {
				break
			}
		}
	}
	fmt.Println(res, "result")
}

func part_2() {
	file, err := os.Open("input3.txt")
	if err != nil {
		log.Fatalf("Error opening the file:%v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	instructions := scanner.Text()

	puzzleInput := make(map[string]Map)

	start := "A"
	end := "Z"

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		input := line[:3]
		linePuzzle := splitline(line)
		puzzleInput[input] = linePuzzle
	}

	results := []int{}
	for input := range puzzleInput {
		if !strings.HasSuffix(input, start) {
			continue
		}
		steps := 0
		for !strings.HasSuffix(input, end) {
			next_inst := string(instructions[steps%len(instructions)])
			input = checkSide(input, next_inst, puzzleInput)
			steps++
		}
		results = append(results, steps)
	}
	val := results[0]
	for i := 1; i < len(results); i++ {
		val = lcm(val, results[i])
	}

	fmt.Println("result", val)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
