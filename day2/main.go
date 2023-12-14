package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

type Set struct {
	red   int
	blue  int
	green int
}

func splitLine(line, r string) []string {
	re := regexp.MustCompile(r)
	return re.Split(line, -1)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Pleave provide a file path")
	}

	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening the file:%v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()[5:]

		content := splitLine(line, ": ")
		// id, _ := strconv.Atoi(content[0])
		sets := splitLine(content[1], "; ")
		s1 := Set{}
		// ok := true

		for _, set := range sets {
			cubes := splitLine(set, ", ")
			s2 := Set{}

			for _, cube := range cubes {
				content := splitLine(cube, " ")
				x, _ := strconv.Atoi(content[0])

				switch content[1] {
				case "red":
					if s2.red += x; s2.red > s1.red {
						s1.red = s2.red
					}
				case "blue":
					if s2.blue += x; s2.blue > s1.blue {
						s1.blue = s2.blue
					}
				case "green":
					if s2.green += x; s2.green > s1.green {
						s1.green = s2.green
					}
				}
			}
			// if s1.red > red || s1.green > green || s1.blue > blue {
			// 	ok = false
			// }
		}
		// if ok {
		//   result += id
		// }
		result += s1.blue * s1.red * s1.green
	}
	fmt.Println(result)
}
