package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func lineSplit(line, r string) []string {
	re := regexp.MustCompile(r)
	return re.Split(line, -1)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Error reading filepath")
	}

	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error reading file:%v", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	result := 0
	cardId := 0
	copies := []int{1}
	for fileScanner.Scan() {
		line := fileScanner.Text()[7:]
		// fmt.Println(line)
		splits := lineSplit(line, `\|`)

		splits[0] = strings.TrimSpace(splits[0])
		splits[1] = strings.TrimSpace(splits[1])
		splits[0] = strings.ReplaceAll(splits[0], "  ", " ")
		splits[1] = strings.ReplaceAll(splits[1], "  ", " ")

		winNumbers := lineSplit(splits[0], " ")
		numbers := lineSplit(splits[1], " ")

		for i := 0; i < copies[cardId]; i++ {
			match := 0
			result++
			for _, winNum := range winNumbers {
				winNumInt, _ := strconv.Atoi(winNum)
				for _, num := range numbers {
					numInt, _ := strconv.Atoi(num)
					if numInt == winNumInt {
						match++
					}
				}
			}

			// fmt.Printf("%v\n", copies)
			// fmt.Println("match---", match)
			// fmt.Println("cardId---", cardId)
			if match > 0 {
				for h := cardId + 1; h < match+cardId+1; h++ {
					if h < len(copies) {
						copies[h] += 1
					} else {
						copies = append(copies, 2)
					}
				}
			} else {
				copies = append(copies, 1)
			}
			// fmt.Printf("after----%v\n", copies)
		}
		cardId++
	}
	// fmt.Printf("%v", copies)
	fmt.Println(result)
}
