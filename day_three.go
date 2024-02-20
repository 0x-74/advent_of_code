package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// LineData represents the data extracted from each line
type LineData struct {
	Numbersindexes [][]int
	Numbers        []string
	Starsindexes   [][]int
}

func main() {
	total := 0
	dig, _ := regexp.Compile(`\d+`)
	star, _ := regexp.Compile(`[^.\d]+`)
	f, err := os.Open("day_three.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []LineData

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numidx := dig.FindAllStringSubmatchIndex(line, -1)
		num := dig.FindAllString(line, -1)
		starsidx := star.FindAllStringSubmatchIndex(line, -1)

		// Create a LineData instance for this line
		lineData := LineData{
			Numbersindexes: numidx,
			Numbers:        num,
			Starsindexes:   starsidx,
		}

		lines = append(lines, lineData)
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i].Numbersindexes); j++ {
			IsAdded := false
			if i < len(lines)-1 {
				// fmt.Print("next_line", lines[i+1].Starsindexes)
				for k := 0; k < len(lines[i+1].Starsindexes); k++ {
					// fmt.Println("+1", lines[i+1].Starsindexes[k])
					if lines[i].Numbersindexes[j][0] >= lines[i+1].Starsindexes[k][0]-1 && lines[i].Numbersindexes[j][0] <= lines[i+1].Starsindexes[k][0]+1 {
						IsAdded = true
					}
					if lines[i].Numbersindexes[j][1] >= lines[i+1].Starsindexes[k][1]-1 && lines[i].Numbersindexes[j][1] <= lines[i+1].Starsindexes[k][1]+1 {
						IsAdded = true
					}
				}
			}
			if i > 0 {
				// fmt.Print("previous_line", lines[i-1].Starsindexes)
				for k := 0; k < len(lines[i-1].Starsindexes); k++ {
					// fmt.Println("-1", lines[i-1].Starsindexes[k])
					if lines[i].Numbersindexes[j][0] >= lines[i-1].Starsindexes[k][0]-1 && lines[i].Numbersindexes[j][0] <= lines[i-1].Starsindexes[k][0]+1 {
						IsAdded = true
					}
					if lines[i].Numbersindexes[j][1] >= lines[i-1].Starsindexes[k][1]-1 && lines[i].Numbersindexes[j][1] <= lines[i-1].Starsindexes[k][1]+1 {
						IsAdded = true
					}
				}
			}
			// fmt.Print("current_line", lines[i].Starsindexes)
			for k := 0; k < len(lines[i].Starsindexes); k++ {
				// fmt.Println("0", lines[i].Starsindexes[k])
				if lines[i].Numbersindexes[j][0] >= lines[i].Starsindexes[k][0]-1 && lines[i].Numbersindexes[j][0] <= lines[i].Starsindexes[k][0]+1 {
					IsAdded = true
				}
				if lines[i].Numbersindexes[j][1] >= lines[i].Starsindexes[k][1]-1 && lines[i].Numbersindexes[j][1] <= lines[i].Starsindexes[k][1]+1 {
					IsAdded = true
				}
			}
			if IsAdded {
				num, err := strconv.Atoi(lines[i].Numbers[j])
				if err != nil {
					fmt.Println("failed to convert to integer: ", err)
				}
				total += num
			}

		}
	}
	fmt.Println(total)
}
