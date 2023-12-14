package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	// "fmt"
	"log"
	"os"
)

type Race struct {
	time int
	dist int
}

func splitLine(line, t string) []string {
	r := regexp.MustCompile(t)
	return r.Split(line, -1)
}

func removeFromLine(line, t string) string {
	r := regexp.MustCompile(t)
	return r.ReplaceAllString(line, "")
}

func populateInput(timeArr, destArr []string) []Race {
	race := []Race{}
	for i := 0; i < len(timeArr); i++ {
		time, _ := strconv.Atoi(timeArr[i])
		dest, _ := strconv.Atoi(destArr[i])
		race = append(race, Race{time, dest})
	}
	return race
}

func populateInputPt2(time, dest string) Race {
	timeNum, _ := strconv.Atoi(time)
	destNum, _ := strconv.Atoi(dest)
	return Race{timeNum, destNum}
}

func calculateWinsWays(inputPuzzle []Race) int {
	result := 0
	for _, race := range inputPuzzle {
		fmt.Printf("%v\n", race)
		winWays := 0
		for i := 0; i <= race.time; i++ {
			timeRemaining := race.time - i
			result := timeRemaining * i
			if result > race.dist {
				winWays++
			}
		}
		if result == 0 {
			result = winWays
		} else {
			result *= winWays
		}
	}
	return result
}

func calculateWinsWaysPart2(race Race) int {
	winWays := 0
	for i := 0; i <= race.time; i++ {
		timeRemaining := race.time - i
		result := timeRemaining * i
		if result > race.dist {
			winWays++
		}
	}
	return winWays
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalf("Error while reading file: %v\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time := splitLine(strings.TrimSpace(scanner.Text()[5:]), ` \s+`)
	timePart2 := removeFromLine(strings.TrimSpace(scanner.Text()[5:]), ` \s+`)
	scanner.Scan()
	dest := splitLine(strings.TrimSpace(scanner.Text()[9:]), ` \s+`)
	destPart2 := removeFromLine(strings.TrimSpace(scanner.Text()[9:]), ` \s+`)
	inputPuzzle := populateInput(time, dest)
	fmt.Printf("%v\n", inputPuzzle)
	fmt.Println("Result-->", calculateWinsWays(inputPuzzle))
	fmt.Println("----------------------------- PART 2---------------------")
	inputPuzzlePt2 := populateInputPt2(timePart2, destPart2)
	fmt.Printf("%v\n", inputPuzzlePt2)
	fmt.Println("Result-->", calculateWinsWaysPart2(inputPuzzlePt2))
}
